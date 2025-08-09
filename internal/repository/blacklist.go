package repository

import "sync"

var blacklist = struct {
	sync.RWMutex
	blocked map[string]struct{}
}{blocked: make(map[string]struct{})}

func AddToBlacklist(token string) {
	blacklist.Lock()
	defer blacklist.Unlock()
	blacklist.blocked[token] = struct{}{}
}

func RemoveFromBlacklist(token string) {
	blacklist.Lock()
	defer blacklist.Unlock()
	delete(blacklist.blocked, token)
}

func IsBlacklisted(token string) bool {
	blacklist.RLock()
	defer blacklist.RUnlock()
	_, ok := blacklist.blocked[token]
	return ok
}
