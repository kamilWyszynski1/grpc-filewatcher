package lru

type LRU interface {
	// Add inserts item into LRU, returns item if any is dropped.
	Add(interface{})
	// Get returns last inserted item.
	Get() interface{}
	// GetN returns n last inserted items.
	GetN(n int) []interface{}
	// GetAll returns all inserted items.
	GetAll() []interface{}
}

type OnEvict func(v interface{})

type lru struct {
	items     []interface{}
	size      int
	cursor    int
	onEvicted OnEvict
}

func NewLRU(size int, evict OnEvict) LRU {
	return &lru{
		size:      size,
		items:     make([]interface{}, size),
		onEvicted: evict,
	}
}

func (l *lru) Add(i interface{}) {
	if l.size == l.cursor {
		// drop last item, shift left and insert
		dropped := l.shift()
		if l.onEvicted != nil {
			l.onEvicted(dropped)
		}
		l.insertAt(i, l.cursor-1)
	} else {
		l.items[l.cursor] = i
		l.cursor = l.cursor + 1
	}
}

func (l *lru) shift() interface{} {
	dropped := l.items[0]
	for i := 0; i < l.size-1; i++ {
		l.items[i] = l.items[i+1]
	}
	return dropped
}

func (l *lru) insertAt(item interface{}, at int) {
	l.items[at] = item
}

func (l lru) Get() interface{} {
	if l.cursor == 0 {
		return nil
	}
	return l.items[l.cursor-1]
}

func (l lru) GetN(n int) []interface{} {
	if n > l.size {
		return nil
	}
	ret := make([]interface{}, 0, n)
	for i := l.size - 1; i >= l.size-n-1; i++ {
		ret = append(ret, l.items[i])
	}
	return ret
}

func (l lru) GetAll() []interface{} {
	return l.items
}
