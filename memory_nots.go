package cache

type MemoryNoTs struct {
	items map[string]interface{}
}

func NewMemoryNoTs() *MemoryNoTs {

	return &MemoryNoTs{
		items: map[string]interface{}{},
	}
}

func (r *MemoryNoTs) Get(key string) (interface{}, error) {
	va, ok := r.items[key]
	if !ok {
		return nil, NotFoundError
	}
	return va, nil
}

func (r *MemoryNoTs) Set(key string, value interface{}) error {
	r.items[key] = value
	return nil
}

func (r *MemoryNoTs) Delete(key string) error {
	delete(r.items, key)
	return nil
}
