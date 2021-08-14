package data

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"threadule/backend/internal/app"
	"threadule/backend/internal/data/models"
)

func connect(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func migrate(ctx *app.Context, db *gorm.DB) error {
	var errs []error

	errs = append(errs, db.AutoMigrate(&models.Group{}))
	errs = append(errs, db.AutoMigrate(&models.User{}))
	errs = append(errs, db.AutoMigrate(&models.Account{}))
	errs = append(errs, db.AutoMigrate(&models.Tweet{}))
	errs = append(errs, db.AutoMigrate(&models.Thread{}))

	var last string
	errorBuilder := strings.Builder{}
	for _, err := range errs {
		if err != nil {
			if last != err.Error() {
				errorBuilder.WriteString(err.Error())
				errorBuilder.WriteString("\n")
			}
			last = err.Error()
		}
	}
	errorString := strings.TrimSpace(errorBuilder.String())
	ctx.Log.Errorf("migration error: %v", errorString)
	if errorBuilder.Len() == 0 {
		return nil
	} else {
		return errors.New(errorString)
	}
}
