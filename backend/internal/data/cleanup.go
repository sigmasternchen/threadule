package data

import "threadule/backend/internal/data/models"

func (d *Data) Cleanup() error {
	if err := d.db.Unscoped().Where("deleted_at IS NOT NULL").Delete(&models.Tweet{}).Error; err != nil {
		return err
	}

	return nil
}
