package cronmanager

import (
	"time"

	"github.com/go-co-op/gocron"
)

type Manager struct {
	Scheduler *gocron.Scheduler
}

func NewScheduler() *Manager {
	sh := gocron.NewScheduler(time.UTC)
	sh.TagsUnique()
	return &Manager{
		Scheduler: sh,
	}
}

func (m *Manager) StartAll() {
	m.Scheduler.StartAsync()
}

func (m *Manager) StopAll() {
	m.Scheduler.Stop()
}

func (m *Manager) Add(name string, interval time.Duration, jobFunc interface{}, params ...interface{}) {
	m.Scheduler.Every(interval.String()).Tag(name).Do(jobFunc, params...)
}

func (m *Manager) Remove(name string) {
	m.Scheduler.RemoveByTag(name)
}

func (m *Manager) RemoveAll() {
	m.Scheduler.Clear()
}

func (m *Manager) Wait() {
	m.Scheduler.StartBlocking()
}
