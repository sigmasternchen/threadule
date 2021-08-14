package logic

import (
	"fmt"
	"threadule/backend/internal/app"
	"threadule/backend/internal/data/models"
)

func Setup(ctx *app.Context) (app.Logic, error) {
	logic := &Logic{
		ctx: ctx,
	}

	logic.startScheduler()

	err := logic.firstTimeSetup()
	if err != nil {
		return nil, err
	}

	return logic, nil
}

func (l *Logic) firstTimeSetup() error {
	c, err := l.ctx.Data.CountUsers()
	if err != nil {
		l.ctx.Log.Errorf("error during first time setup check: %v", err)
		return err
	}

	if c != 0 {
		return nil
	}

	// no users -> probably first time setup
	l.ctx.Log.Info("executing first time setup")

	adminGroup := models.GetDefaultAdminGroup()
	err = l.ctx.Data.AddGroup(adminGroup)
	if err != nil {
		l.ctx.Log.Errorf("couldn't create admin group: %v", err)
		return err
	}

	adminUser := models.GetDefaultAdminUser()
	adminUser.Groups = []*models.Group{adminGroup}
	password := l.defaultPassword()
	adminUser.Password, err = l.hashPassword(password)

	if err != nil {
		// if this fails we can't recover anyway
		l.ctx.Log.Fatal(err)
	}

	err = l.ctx.Data.CreateUser(adminUser)
	if err != nil {
		l.ctx.Log.Errorf("couldn't create admin user: %v", err)
		return err
	}

	fmt.Println("initial credentials:")
	fmt.Printf("Username: %s\n", adminUser.Username)
	fmt.Printf("Password: %s\n", password)

	return nil
}
