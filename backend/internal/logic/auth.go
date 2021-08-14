package logic

func (l *Logic) scheduleTriggerAuth() {
	err := l.ctx.Data.CleanupSessions()
	if err != nil {
		l.ctx.Log.Errorf("couldn't clean up sessions: %v", err)
	}
}
