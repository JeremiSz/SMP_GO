package server

import "sync"

var AUTHORS = []string{}
var TEXTS = []string{}
var AUTHORS_LOCK sync.RWMutex
var TEXT_LOCK sync.RWMutex

func addMessage(author string, text string) {
	AUTHORS_LOCK.Lock()
	TEXT_LOCK.Lock()
	defer AUTHORS_LOCK.Unlock()
	defer TEXT_LOCK.Unlock()
	AUTHORS = append(AUTHORS, author)
	TEXTS = append(TEXTS, text)
}

func getAuthros() []string {
	AUTHORS_LOCK.RLock()
	defer AUTHORS_LOCK.RUnlock()
	return AUTHORS
}

func getTexts() []string {
	TEXT_LOCK.RLock()
	defer TEXT_LOCK.RUnlock()
	return TEXTS
}
