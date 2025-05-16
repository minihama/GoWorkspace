package cmap

import "sync"

type ConcurrentMap struct {
	sync.RWMutex
	items map[string]interface{}
}

func New() *ConcurrentMap {
	cm := new(ConcurrentMap)
	cm.items = make(map[string]interface{})
	return cm
}

func (cm *ConcurrentMap) Set(key string, value interface{}) {
	cm.Lock()
	cm.items[key] = value
	cm.Unlock()
}

func (cm *ConcurrentMap) Get(key string) (interface{}, bool) {
	cm.RLock()
	defer cm.RUnlock()
	value, ok := cm.items[key]
	return value, ok
}
