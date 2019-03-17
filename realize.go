package cache

import "container/list"

type LRUNots struct {
	cache Cache
	list  *list.List
	size  int
}

type kv struct {
	k string
	v interface{}
}

func NewLRUNots(size int) *LRUNots {
	if size < 1 {
		panic("invaild cache size")
	}

	return &LRUNots{
		cache: NewMemoryNoTs(),
		size:  size,
		list:  list.New(),
	}
}

func (r *LRUNots) Get(key string) (interface{}, error) {
	res, err := r.cache.Get(key)
	if err != nil {
		return nil, err
	}

	element := res.(*list.Element)
	r.list.MoveToFront(element)

	return element.Value.(*kv).v, nil
}

func (r *LRUNots) Set(key string, value interface{}) error {
	res, err := r.cache.Get(key)
	if err != nil && err != NotFoundError {
		return err
	}

	var element *list.Element
	if err == NotFoundError {
		element = &list.Element{Value: &kv{k: key, v: value}}
		r.list.PushFront(element)
	} else {
		element := res.(*list.Element)
		element.Value.(*kv).v = value
		r.list.MoveToFront(element)
	}

	if r.list.Len() > r.size {
		r.moveElement(r.list.Back())
	}

	return nil
}

func (r *LRUNots) Delete(key string) error {
	res, err := r.cache.Get(key)
	if err != nil && err != NotFoundError {
		return err
	}

	if err == NotFoundError {
		return nil
	}

	element := res.(*list.Element)
	return r.moveElement(element)
}

func (r *LRUNots) moveElement(element *list.Element) error {
	r.list.Remove(element)
	return r.cache.Delete(element.Value.(*kv).k)
}
