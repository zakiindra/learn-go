package maps

type Dictionary map[string]string

const (
    ErrNotFound = DictionaryErr("Could not find the word you were looking for")
    ErrValueExists = DictionaryErr("Cannot add value, value already exists")
    ErrKeyDoesNotExist = DictionaryErr("Cannot update non existing key")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
    value, ok := d[key]
    if !ok {
        return "", ErrNotFound
    }
    return value, nil
}

func (d Dictionary) Add(key, value string) error {
    _, err := d.Search(key)

    switch err {
    case ErrNotFound:
        d[key] = value
    case nil:
        return ErrValueExists
    default:
        return err
    }

    return nil
}

func (d Dictionary) Update(key, value string) error {
    _, err := d.Search(key)

    switch err {
    case ErrNotFound:
       return ErrKeyDoesNotExist
    case nil:
        d[key] = value
    default:
        return err
    }

    return nil
}

func (d Dictionary) Delete(key string) {
    delete(d, key)
}