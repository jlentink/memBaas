package cache

import "time"

type Entry struct {
	Key            string
	Data           string
	TTL            int64
	AllowOverwrite bool
	Created        time.Time
}

func NewEntry(key, value string, ttl int64) Entry {
	return Entry{
		Key:     key,
		Data:    value,
		TTL:     ttl,
		Created: time.Now(),
	}
}
