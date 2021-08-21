package logic

import "time"

func (l *Logic) startScheduler() {
	twitterTicker := time.NewTicker(time.Minute)
	cleanupTicker := time.NewTicker(time.Hour * 24)

	go func() {
		for {
			_ = <-twitterTicker.C
			l.scheduleTriggerTwitter()
		}
	}()

	go func() {
		for {
			_ = <-cleanupTicker.C
			l.scheduledCleanup()
		}
	}()
}
