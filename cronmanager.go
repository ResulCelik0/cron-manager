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

func (m *Manager) Status(name string) (status bool, err error) {
	j, err := m.Scheduler.FindJobsByTag(name)
	for _, v := range j {
		status = v.IsRunning()
	}
	return

}

func (m *Manager) StartAll() {
	m.Scheduler.StartAsync()
}

func (m *Manager) StopAll() {
	m.Scheduler.Stop()
}

func (m *Manager) Check(name string) (err error) {
	_, err = m.Scheduler.FindJobsByTag(name)
	return
}

func (m *Manager) Add(name string, interval time.Duration, jobFunc interface{}, params ...interface{}) (err error) {
	_, err = m.Scheduler.Every(interval.String()).Tag(name).Do(jobFunc, params...)
	return
}

// Şimdilik sadece interval değerini değiştirebiliyoruz
func (m *Manager) Update(name string, interval time.Duration) (err error) {
	j, err := m.Scheduler.FindJobsByTag(name)
	m.Scheduler.Job(j[0]).Every(interval.String()).Update()
	return
}

func (m *Manager) Remove(name string) (err error) {
	err = m.Scheduler.RemoveByTag(name)
	return
}

func (m *Manager) RemoveAll() {
	m.Scheduler.Clear()
}

func (m *Manager) Wait() {
	m.Scheduler.StartBlocking()
}

func (m *Manager) Break() {
	m.Scheduler.StopBlockingChan()
}

func (m *Manager) ListJob() (jobs []string) {
	jobs = m.Scheduler.GetAllTags() //Taglar Uniq olduğundan  çalışan monitör sayılarını verir
	return
}
