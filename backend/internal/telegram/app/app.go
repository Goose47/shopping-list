package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

type App struct {
	log    *slog.Logger
	secret string
	url    string
}

func New(
	log *slog.Logger,
	secret string,
	url string,
) *App {
	return &App{
		log:    log,
		secret: secret,
		url:    url,
	}
}

func (app *App) Run() error {
	const op = "telegram.App.Run"
	log := app.log.With(slog.String("op", op))

	bot, err := tgbotapi.NewBotAPI(app.secret)
	if err != nil {
		log.Error("failed to create bot instance", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	log.Info("start telegram poll")
	for update := range updates {
		log.Info("got telegram message")
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		switch update.Message.Command() {
		case "start":
			sendMiniAppButton(log, bot, update.Message.Chat.ID, app.url)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда")
			_, err = bot.Send(msg)
			log.Error("failed to send message", slog.Any("error", err))
		}
	}

	return nil
}

func sendMiniAppButton(
	log *slog.Logger,
	bot *tgbotapi.BotAPI,
	chatID int64,
	url string,
) {
	msg := tgbotapi.NewMessage(chatID, "Открой мини-приложение:")
	button := tgbotapi.NewInlineKeyboardButtonURL("🛒 Список покупок", url)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button))
	msg.ReplyMarkup = keyboard

	_, err := bot.Send(msg)
	if err != nil {
		log.Error("failed to send button message", slog.Any("error", err))
	}
}
