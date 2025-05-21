package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"net/url"
	"shopping-list/internal/domain/models"
	"sort"
	"strings"
)

func TelegramAuthMiddleware(log *slog.Logger, db *gorm.DB, botSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		initData := c.GetHeader("Telegram-Init")
		log.Info("INITDATA: ", initData)

		if initData == "" {
			log.Error("missing TelegramInit header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing TelegramInit header"})
			c.Abort()
			return
		}

		data, err := url.ParseQuery(initData)
		if err != nil {
			log.Error("failed to parse init data", slog.Any("error", err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("failed to parse init data: %s", err.Error())})
			c.Abort()
			return
		}

		err = validateInitData(data, botSecret)
		if err != nil {
			log.Error("invalid init data", slog.Any("error", err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("invalid init data: %s", err.Error())})
			c.Abort()
			return
		}

		var user models.User
		if err := db.First(&user, 1).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
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
	fmt.Println(checkString)
	if calculatedHash != hash {
		return fmt.Errorf("hash mismatch")
	}

	return nil
}
