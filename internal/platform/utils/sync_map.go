package utils

import (
	"sync"
)

type SyncMap struct {
	data sync.Map
}

type entry struct {
	Value interface{}
}

func NewSyncMap() *SyncMap {
	return &SyncMap{}
}

func (t *SyncMap) Store(key string, val interface{}) {
	t.data.Store(key, entry{
		Value: val,
	})
}

func (t *SyncMap) Delete(key string) {
	t.data.Delete(key)
}

func (t *SyncMap) Load(key string) (val interface{}) {
	data, ok := t.data.Load(key)
	if !ok {
		return nil
	}

	entry, ok := data.(entry)
	if !ok {
		return nil
	}

	return entry.Value
}
