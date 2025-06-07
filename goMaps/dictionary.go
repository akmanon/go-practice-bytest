package gomaps

type Dictionary map[string]string

const (
	ErrNotFound        = DictionaryErr("Word not found during search")
	ErrWordExist       = DictionaryErr("word already exists in dictionary")
	ErrWordDoesntExist = DictionaryErr("word does not exists in dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	value, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = newDefinition
	case ErrNotFound:
		return ErrWordDoesntExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case ErrNotFound:
		return ErrWordDoesntExist
	default:
		return err
	}
	return nil

}
