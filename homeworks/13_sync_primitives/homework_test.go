package homework13

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRWMutexWithWriter(t *testing.T) {
	mutex := NewRWMutex()
	mutex.Lock() // writer

	var mutualExlusionWithWriter atomic.Bool
	mutualExlusionWithWriter.Store(true)
	var mutualExlusionWithReader atomic.Bool
	mutualExlusionWithReader.Store(true)

	go func() {
		mutex.Lock() // another writer
		mutualExlusionWithWriter.Store(false)
	}()

	go func() {
		mutex.RLock() // another reader
		mutualExlusionWithReader.Store(false)
	}()

	time.Sleep(time.Second)
	assert.True(t, mutualExlusionWithWriter.Load())
	assert.True(t, mutualExlusionWithReader.Load())
}

func TestRWMutexWithReaders(t *testing.T) {
	mutex := NewRWMutex()
	mutex.RLock() // reader

	var mutualExlusionWithWriter atomic.Bool
	mutualExlusionWithWriter.Store(true)

	go func() {
		mutex.Lock() // another writer
		mutualExlusionWithWriter.Store(false)
	}()

	time.Sleep(time.Second)
	assert.True(t, mutualExlusionWithWriter.Load())
}

func TestRWMutexMultipleReaders(t *testing.T) {
	mutex := NewRWMutex()
	mutex.RLock() // reader

	var readersCount atomic.Int32
	readersCount.Add(1)

	go func() {
		mutex.RLock() // another reader
		readersCount.Add(1)
	}()

	go func() {
		mutex.RLock() // another reader
		readersCount.Add(1)
	}()

	time.Sleep(time.Second)
	assert.Equal(t, int32(3), readersCount.Load())
}

func TestRWMutexWithWriterPriority(t *testing.T) {
	mutex := NewRWMutex()
	mutex.RLock() // reader

	var mutualExlusionWithWriter atomic.Bool
	mutualExlusionWithWriter.Store(true)
	var readersCount atomic.Int32
	readersCount.Add(1)

	go func() {
		mutex.Lock() // another writer is waiting for reader
		mutualExlusionWithWriter.Store(false)
	}()

	time.Sleep(time.Second)

	go func() {
		mutex.RLock() // another reader is waiting for a higher priority writer
		readersCount.Add(1)
	}()

	go func() {
		mutex.RLock() // another reader is waiting for a higher priority writer
		readersCount.Add(1)
	}()

	time.Sleep(time.Second)

	assert.True(t, mutualExlusionWithWriter.Load())
	assert.Equal(t, int32(1), readersCount.Load())
}
