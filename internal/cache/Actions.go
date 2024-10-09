package cache

import (
	log "github.com/jlentink/yaglogger"
	"sync"
	"time"
)

var (
	data  = make(map[string]Entry)
	mutex sync.Mutex
)

// Set sets a key value pair in the cache
func Set(key, value string, ttl int64, overwrite bool) Status {
	mutex.Lock()
	defer mutex.Unlock()
	if key == "" {
		log.Debug("Key is empty. Returning INVALID_KEY")
		return InvalidKey
	}

	if _, ok := data[key]; ok {
		if !overwrite {
			log.Debug("Key already exists. Returning NOT_UNIQUE")
			return NotUnique
		}
		if data[key].AllowOverwrite == false {
			log.Debug("Key already exists and does not allow overwrite. Returning Override Denied")
			return OverwriteDenied
		}
	}

	log.Debug("Setting key: %s", key)
	data[key] = NewEntry(key, value, ttl)
	return Ok
}

// Get retrieves a value from the cache
func Get(key string) (value string, status Status) {
	if _, ok := data[key]; ok {
		log.Debug("Key found and returning value: %s", key)
		return data[key].Data, Ok
	}
	log.Debug("Key not found: %s", key)
	return "", NotFound
}

// Delete removes a key from the cache
func Delete(key string) (status Status) {
	mutex.Lock()
	defer mutex.Unlock()
	log.Debug("Deleting key: %s", key)
	if _, ok := data[key]; ok {
		delete(data, key)
		return Ok
	}
	return NotFound
}

// Cleanup removes all keys from the cache that have expired
func Cleanup() {
	for _, entry := range data {
		if time.Since(entry.Created) > time.Duration(entry.TTL)*time.Second && entry.TTL != 0 {
			log.Debug("Deleting key: %s - TTL expired", entry.Key)
			status := Delete(entry.Key)
			if status != Ok {
				log.Error("Failed to delete key: %s - %s", entry.Key, status)
			}
		}
	}
}
