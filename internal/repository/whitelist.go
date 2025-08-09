package repository

import "sync"

var whitelist = struct {
	sync.RWMutex
	allowed map[string]struct{}
}{allowed: make(map[string]struct{})}

func AddToWhitelist(token string) {
	whitelist.Lock()
	defer whitelist.Unlock()
	whitelist.allowed[token] = struct{}{}
}

func RemoveFromWhitelist(token string) {
	whitelist.Lock()
	defer whitelist.Unlock()
	delete(whitelist.allowed, token)
}

func IsWhitelisted(token string) bool {
	whitelist.RLock()
	defer whitelist.RUnlock()
	_, ok := whitelist.allowed[token]
	return ok
}
