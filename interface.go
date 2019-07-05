package lru

// Interface represents the interface for a Cache.
type Interface interface {
	Add(key interface{}, value int, size int) bool
	Get(key interface{}) (value int, ok bool)
//	Contains(key interface{}) (ok bool)
//	Peek(key interface{}) (value int, ok bool)
//	Remove(key interface{}) bool
//	Keys() []interface{}
	Size() int
	Purge()
//	Stop()
}
