package core

type Store struct {
	data map[string]string
}

func (store *Store) Set(key, value string) {
	store.data[key] = value
}

func (store *Store) Get(key, value string) (string, bool) {
	val, exists := store.data[key]
	return val, exists
}
