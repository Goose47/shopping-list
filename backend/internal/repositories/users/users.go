package users

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"shopping-list/internal/domain/models"
	"shopping-list/internal/repositories"
)

func First(
	db *gorm.DB,
	filters map[string]any,
) (*models.User, error) {
	query := db
	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	var user models.User
	err := query.First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &user, repositories.ErrNotFound
	}

	return &user, err
}

func Store(
	db *gorm.DB,
	telegramID uint,
	firstName string,
	lastName string,
	username string,
	languageCode string,
	allowsWriteToPm bool,
	photoUrl string,
) (*models.User, error) {
	user := models.User{
		TelegramID:      telegramID,
		FirstName:       firstName,
		LastName:        lastName,
		Username:        username,
		LanguageCode:    languageCode,
		AllowsWriteToPm: allowsWriteToPm,
		PhotoUrl:        photoUrl,
	}
	err := db.Create(&user).Error

	return &user, err
}
