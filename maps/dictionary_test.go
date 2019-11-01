package maps

import "testing"

func TestSearch(t *testing.T) {
    dictionary := Dictionary{"test": "this is just a test"}

    t.Run("known key", func(t *testing.T) {
        actual, _ := dictionary.Search("test")
        expected := "this is just a test"

        assertStrings(t, actual, expected)
    })

    t.Run("unknown key", func( *testing.T) {
        _, actual := dictionary.Search("unknown")

        assertError(t, actual, ErrNotFound)
    })
}

func TestAdd(t *testing.T) {
    t.Run("new value", func(t *testing.T) {
        dictionary := Dictionary{}
        key := "test"
        value := "this is just a test"
        err := dictionary.Add(key, value)

        assertError(t, err, nil)
        assertValue(t, dictionary, key, value)
    })
    
    t.Run("existing value", func(t *testing.T) {
        key := "test"
        value := "this is just a test"
        dictionary := Dictionary{key: value}
        err := dictionary.Add(key, "new value")

        assertError(t, err, ErrValueExists)
        assertValue(t, dictionary, key, value)
    })
}

func TestUpdate(t *testing.T) {
    t.Run("existing value", func (t *testing.T) {
        key := "test"
        value := "this is just a test"
        dictionary := Dictionary{key: value}
        newValue := "new test value"

        err := dictionary.Update(key, newValue)

        assertError(t, err, nil)
        assertValue(t, dictionary, key, newValue)
    })

    t.Run("new value", func (t *testing.T) {
        key := "test"
        value := "this is just a test"
        dictionary := Dictionary{}

        err := dictionary.Update(key, value)
        assertError(t, err, ErrKeyDoesNotExist)
    })
}

func TestDelete(t *testing.T) {
    key := "test"
    dictionary := Dictionary{key : "this is just a test"}

    dictionary.Delete(key)

    _, err := dictionary.Search(key)
    if err != ErrNotFound {
        t.Errorf("Expected %s to be deleted", key)
    }

}

func assertValue(t *testing.T, dictionary Dictionary, key, value string) {
    t.Helper()

    actual, _ := dictionary.Search(key)
    assertStrings(t, actual, value)
}

func assertError(t *testing.T, actual, expected error) {
    if actual != expected {
        t.Errorf("Expected %q but got actual %q", expected, actual)
    }
}

func assertStrings(t *testing.T, actual, expected string) {
    t.Helper()

    if actual != expected {
        t.Errorf("Expected %q but got actual %q", expected, actual)
    }
}