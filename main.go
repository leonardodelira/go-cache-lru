package main

import "fmt"

type Item struct {
	key   string
	value any
	next  *Item
	prev  *Item
}

type Cache struct {
	data     map[string]*Item
	capacity int
	head     *Item
	tail     *Item
}

func main() {
	lru := New(4)
	lru.Set("a", 1)
	lru.Set("a", 2)
	lru.Set("a", 3)
	lru.Show()
}

func New(capacity int) *Cache {

	return &Cache{
		data:     make(map[string]*Item, capacity),
		capacity: capacity,
	}
}

func (l *Cache) Get(key string) any {
	item, ok := l.data[key]
	if !ok {
		return -1
	}

	if item == l.head {
		return item.value
	}

	if item.prev != nil {
		item.prev.next = item.next
	}

	if item.next != nil {
		item.next.prev = item.prev
	}

	if item == l.tail {
		l.tail = item.prev
		l.tail.next = nil
	}

	l.head.prev = item
	item.next = l.head
	l.head = item
	l.head.prev = nil

	return item.value
}

func (l *Cache) Set(key string, value any) {
	//se adicionar a mesma key, remover o valor e adicionar o novo
	if item, ok := l.data[key]; ok {
		if item.prev != nil {
			item.prev.next = item.next
		}

		if item.next != nil {
			item.next.prev = item.prev
		}

		if item == l.tail && len(l.data) > 1 {
			l.tail = item.prev
			l.tail.next = nil
		}

		if item == l.head {
			l.head = item.next
		}

		delete(l.data, key)
	}

	newItem := &Item{key: key, value: value}

	if l.head == nil {
		l.head = newItem
		l.tail = newItem
		l.data[key] = newItem
		return
	}

	if len(l.data) == l.capacity {
		lastItem := l.tail
		l.tail = lastItem.prev
		l.tail.next = nil
		delete(l.data, lastItem.key)
	}

	l.head.prev = newItem
	newItem.next = l.head
	l.head = newItem
	l.head.prev = nil
	l.data[key] = newItem
}

func (l *Cache) Show() {
	curr := l.head
	for curr.next != nil {
		fmt.Println(curr.value, "next: ", curr.next.value)
		curr = curr.next
	}
	fmt.Println(curr.value, "next: nil")
}
