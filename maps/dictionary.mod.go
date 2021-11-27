package maps

type Dictionary map[string]string

const (
	ErrNotFound           = DictionaryError("could not find the word you are looking for")
	ErrKeyAlreadyAdded    = DictionaryError("key already added")
	ErrorKeyDoesNotExists = DictionaryError("key does not exist")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	result, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}

func (d Dictionary) Add(key, value string) error {
	_, exists := d[key]
	if exists {
		return ErrKeyAlreadyAdded
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, exist := d[key]
	if !exist {
		return ErrorKeyDoesNotExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
