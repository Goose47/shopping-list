package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/url"
	"shopping-list/internal/domain/dto"
	"shopping-list/internal/lib/responses"
	"shopping-list/internal/repositories"
	"shopping-list/internal/repositories/users"
	"sort"
	"strings"
)

func TelegramAuthMiddleware(log *slog.Logger, db *gorm.DB, botSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		initData := c.GetHeader("Telegram-Init")

		if initData == "" {
			log.Error("missing TelegramInit header")
			responses.UnauthorizedError(c)
			c.Abort()
			return
		}

		data, err := url.ParseQuery(initData)
		if err != nil {
			log.Error("failed to parse init data", slog.Any("error", err))
			responses.UnauthorizedError(c)
			c.Abort()
			return
		}

		err = validateInitData(data, botSecret)
		if err != nil {
			log.Error("invalid init data", slog.Any("error", err))
			responses.UnauthorizedError(c)
			c.Abort()
			return
		}

		userString := data.Get("user")
		var userDTO dto.TelegramInitDataUser
		err = json.Unmarshal([]byte(userString), &userDTO)
		if err != nil {
			log.Error("failed to unmarshal init data user", slog.Any("error", err))
			responses.UnauthorizedError(c)
			c.Abort()
			return
		}

		user, err := users.First(db, map[string]any{
			"telegram_id": userDTO.ID,
		})
		if errors.Is(err, repositories.ErrNotFound) {
			user, err = users.Store(
				db,
				userDTO.ID,
				userDTO.FirstName,
				userDTO.LastName,
				userDTO.Username,
				userDTO.LanguageCode,
				userDTO.AllowsWriteToPm,
				userDTO.PhotoUrl,
			)
			if err != nil {
				log.Error("failed to create user", slog.Any("error", err))
				responses.InternalServerError(c)
				c.Abort()
				return
			}
		}

		// Store the user in the Gin context
		c.Set("user", user)
		c.Next()
	}
}

func validateInitData(
	data url.Values,
	botSecret string,
) error {
	var (
		hash  string
		pairs = make([]string, 0, len(data))
	)

	for k, v := range data {
		if k == "hash" {
			hash = v[0]
			continue
		}
		pairs = append(pairs, k+"="+v[0])
	}

	if hash == "" {
		return fmt.Errorf("missing hash")
	}

	sort.Strings(pairs)
	checkString := strings.Join(pairs, "\n")

	skHmac := hmac.New(sha256.New, []byte("WebAppData"))
	skHmac.Write([]byte(botSecret))

	impHmac := hmac.New(sha256.New, skHmac.Sum(nil))
	impHmac.Write([]byte(checkString))

	calculatedHash := hex.EncodeToString(impHmac.Sum(nil))
	if calculatedHash != hash {
		return fmt.Errorf("hash mismatch")
	}

	return nil
}
