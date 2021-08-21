package logic

func (l *Logic) scheduledCleanup() {
	l.ctx.Log.Info("doing cleanup...")

	l.cleanupSessions()

	err := l.ctx.Data.Cleanup()
	if err != nil {
		l.ctx.Log.Errorf("couldn't clean up database: %v", err)
	}
}
