package logic

import (
	"threadule/backend/internal/data/models"
)

func (l *Logic) checkPrivilege(originalUser *models.User, currentUser *models.User) error {
	if originalUser.ID == currentUser.ID {
		// own user can always be modified
		return nil
	}

	ok := false
	needsAdmin := false

	for _, group := range originalUser.Groups {
		if group.AdminGroup {
			needsAdmin = true
			break
		}
	}

	for _, group := range currentUser.Groups {
		if group.AdminGroup {
			// reset needs admin flag
			needsAdmin = false
		}
		if group.ManageUsers {
			ok = true
		}
	}

	if !ok || needsAdmin {
		return ErrInsufficientPrivilege
	} else {
		return nil
	}
}

func (l *Logic) UpdateUser(userToUpdate *models.User, currentUser *models.User) error {
	originalUser, err := l.ctx.Data.GetUser(userToUpdate.ID)
	if err != nil {
		return ErrNotFound
	}

	err = l.checkPrivilege(originalUser, currentUser)
	if err != nil {
		return err
	}

	if originalUser.Username != userToUpdate.Username {
		_, err = l.ctx.Data.GetUserByUsername(userToUpdate.Username)
		if err == nil {
			// if no error there exists already a user with that name
			return ErrConflict
		}
	}

	if userToUpdate.Password != "" {
		userToUpdate.Password, err = l.hashPassword(userToUpdate.Password)
		if err != nil {
			return err
		}
	}

	userToUpdate.Groups = nil
	userToUpdate.CreatedAt = originalUser.CreatedAt

	err = l.ctx.Data.UpdateUser(userToUpdate)

	userToUpdate.Groups = originalUser.Groups
	userToUpdate.Password = ""
	return err
}
