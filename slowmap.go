package slowmap

import (
	"sync"
)

type NewMap struct {
	mutex      sync.RWMutex
	maxCount   int
	threshold  int
	count      int
	dictionary map[any]any
}

func NewSlowMap() *NewMap {
	return &NewMap{
		threshold:  1000,
		maxCount:   10000,
		dictionary: make(map[any]any),
	}
}

func (m *NewMap) SetRadius(max, threshold int) {
	m.maxCount = max
	m.threshold = threshold
}

func (m *NewMap) Del(key any) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.count++
	delete(m.dictionary, key)
	if m.count >= m.maxCount && len(m.dictionary) < m.threshold {
		m.count = 0
		mm := make(map[any]any)
		for k, v := range m.dictionary {
			mm[k] = v
		}
		m.dictionary = mm
	}
}

func (m *NewMap) Set(key, value any) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.dictionary[key] = value
}

func (m *NewMap) Get(key any) (any, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	val, ok := m.dictionary[key]
	return val, ok
}
