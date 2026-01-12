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

	m.writersWaiting++
	for m.writerActive || m.readerCount != 0 {
		m.writerCond.Wait()
	}
	m.writerActive = true
	m.writersWaiting--
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

	for m.writerActive || m.writersWaiting > 0 {
		m.writerCond.Wait()
	}
	m.readerCount++
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
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.writerActive || m.readerCount > 0 {
		return false
	}

	m.writerActive = true
	return true
}

func (m *RWMutex) TryRLock() bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.writerActive || m.writersWaiting > 0 {
		return false
	}

	m.readerCount++
	return true
}
