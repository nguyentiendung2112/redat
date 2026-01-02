package core

type Store struct {
	data map[string]string
}

func (store *Store) Init() {
	store.data = make(map[string]string)
}

func (store *Store) Set(key, value string) {
	store.data[key] = value
}

func (store *Store) Get(key string) (string, bool) {
	val, exists := store.data[key]
	return val, exists
}

func (store *Store) Delete(key string) {
	delete(store.data, key)
}

func (store *Store) Keys() []string {
	keys := make([]string, 0, len(store.data))
	for k := range store.data {
		keys = append(keys, k)
	}
	return keys
}
