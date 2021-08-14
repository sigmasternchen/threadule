package logic

import "time"

func (l *Logic) startScheduler() {
	ticker := time.NewTicker(time.Minute)

	go func() {
		for {
			_ = <-ticker.C
			l.scheduleTriggerTwitter()
			l.scheduleTriggerAuth()
		}
	}()
}
