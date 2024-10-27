package main

import (
	"fmt"
	"sync"
)

type (
	DictKey    string
	DictVal    string
	Dictionary struct {
		elements map[DictKey]DictVal
		lock     sync.Mutex
	}
)

func (d *Dictionary) Put(key DictKey, val DictVal) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.elements == nil {
		d.elements = make(map[DictKey]DictVal)
	}
	d.elements[key] = val
}

func (d *Dictionary) Remove(key DictKey) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	_, exist := d.elements[key]
	if exist {
		delete(d.elements, key)
	}
	return exist
}

func (d *Dictionary) Contain(key DictKey) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	_, exist := d.elements[key]
	return exist
}

func (d *Dictionary) Find(key DictKey) DictVal {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.elements[key]
}

func (d *Dictionary) Reset() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.elements = make(map[DictKey]DictVal)
}

func (d *Dictionary) NumberOfElements() int {
	d.lock.Lock()
	defer d.lock.Unlock()
	return len(d.elements)
}

func (d *Dictionary) GetKeys() []DictKey {
	d.lock.Lock()
	defer d.lock.Unlock()
	var dictKeys []DictKey
	dictKeys = []DictKey{}
	var key DictKey
	for key = range d.elements {
		dictKeys = append(dictKeys, key)
	}
	return dictKeys
}

func (d *Dictionary) GetVals() []DictVal {
	d.lock.Lock()
	defer d.lock.Unlock()
	var dictKeys []DictVal
	dictKeys = []DictVal{}

	for _, val := range d.elements {
		dictKeys = append(dictKeys, val)
	}
	return dictKeys
}

func main() {
	var dict *Dictionary = &Dictionary{}
	dict.Put("1", "1")
	dict.Put("2", "2")
	dict.Put("3", "3")
	dict.Put("4", "4")
	fmt.Println(dict)
}
