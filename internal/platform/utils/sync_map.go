package utils

import (
	"sync"
	"time"
)

type SyncMap struct {
	TTL  time.Duration
	data sync.Map
}

type expireEntry struct {
	ExpiresAt time.Time
	Value     interface{}
}

func NewSyncMap(ttl time.Duration) *SyncMap {
	return &SyncMap{TTL: ttl}
}

func (t *SyncMap) Store(key string, val interface{}) {
	t.data.Store(key, expireEntry{
		ExpiresAt: time.Now().Add(t.TTL),
		Value:     val,
	})
}

func (t *SyncMap) Delete(key string) {
	t.data.Delete(key)
}

func (t *SyncMap) Load(key string) (val interface{}) {
	entry, ok := t.data.Load(key)
	if !ok {
		return nil
	}

	expEntry, ok := entry.(expireEntry)
	if !ok {
		return nil
	}

	if time.Now().After(expEntry.ExpiresAt) {
		return nil
	}

	return expEntry.Value
}

func (t *SyncMap) CleaningBackground(cleaning time.Duration) {
	go func() {
		for now := range time.Tick(cleaning) {
			t.data.Range(func(k, v interface{}) bool {
				if expEntry, ok := v.(expireEntry); ok && expEntry.ExpiresAt.After(now) {
					t.data.Delete(k)
				}
				return true
			})
		}
	}()
}
