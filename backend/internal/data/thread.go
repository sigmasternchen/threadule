package data

import "threadule/backend/internal/data/models"

func (d *Data) UpdateThread(thread *models.Thread) error {
	return d.db.Save(thread).Error
}

func (d *Data) GetTweetsForThread(thread *models.Thread) ([]models.Tweet, error) {
	var tweets []models.Tweet
	err := d.db.
		Where("thread_id = ?", thread.ID).
		Order("ordinal ASC").
		Find(&tweets).Error
	return tweets, err
}
