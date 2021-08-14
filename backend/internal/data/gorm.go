package data

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"threadule/backend/internal/data/models"
)

func connect(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func migrate(db *gorm.DB) error {
	var errs []error

	errs = append(errs, db.AutoMigrate(&models.Group{}))
	errs = append(errs, db.AutoMigrate(&models.User{}))
	errs = append(errs, db.AutoMigrate(&models.Account{}))
	errs = append(errs, db.AutoMigrate(&models.Tweet{}))
	errs = append(errs, db.AutoMigrate(&models.Thread{}))

	errorBuilder := strings.Builder{}
	for _, err := range errs {
		if err != nil {
			errorBuilder.WriteString(err.Error())
			errorBuilder.WriteString("\n")
		}
	}
	if errorBuilder.Len() == 0 {
		return nil
	} else {
		return errors.New(strings.TrimSpace(errorBuilder.String()))
	}
}
