package homework13

import (
	"sync"
)

type RWMutex struct {
	readerCount    int
	mutex          sync.Mutex
	writerCond     *sync.Cond
	writerActive   bool
	writersWaiting int
}

func NewRWMutex() *RWMutex {
	rwmutex := RWMutex{}
	rwmutex.writerCond = sync.NewCond(&rwmutex.mutex)
	return &rwmutex
}

func (m *RWMutex) Lock() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.underLock()
}

func (m *RWMutex) Unlock() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.writerActive = false
	m.writerCond.Broadcast()
}

func (m *RWMutex) RLock() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.underRLock()
}

func (m *RWMutex) RUnlock() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.readerCount--
	if m.readerCount == 0 {
		m.writerCond.Broadcast()
	}
}

func (m *RWMutex) TryLock() bool {
	ok := m.mutex.TryLock()
	if !ok {
		return false
	}
	defer m.mutex.Unlock()

	m.underLock()

	return true
}

func (m *RWMutex) TryRLock() bool {
	ok := m.mutex.TryLock()
	if !ok {
		return false
	}
	defer m.mutex.Unlock()

	m.underRLock()

	return true
}

func (m *RWMutex) underLock() {
	m.writersWaiting++
	for m.writerActive || m.readerCount != 0 {
		m.writerCond.Wait()
	}
	m.writerActive = true
	m.writersWaiting--
}

func (m *RWMutex) underRLock() {
	for m.writerActive || m.writersWaiting > 0 {
		m.writerCond.Wait()
	}
	m.readerCount++
}
