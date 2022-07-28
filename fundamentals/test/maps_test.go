package fundamentalstest

import (
	"log"
	"testing"
)

type Dictionary map[string]string

const (
	errNotFound         = DictionaryErr("could not find word you were looking for")
	errWordExist        = DictionaryErr("can not add word because it already exists")
	errWordDoesNotExist = DictionaryErr("can not update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = definition
	case nil:
		return errWordExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefiniton string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errWordDoesNotExist
	case nil:
		d[word] = newDefiniton
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errNotFound
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}

func (d Dictionary) ResetDictionary() {
	for _, value := range d {
		log.Println("value", value)
		delete(d, value)
	}
}

func TestFindMaps(t *testing.T) {
	dictionary := Dictionary{"text": "hallo"}
	testCases := []struct {
		desc       string
		word       string
		expected   string
		errMessage string
	}{
		{
			desc:     "success find words",
			word:     "text",
			expected: "hallo",
		},
		{
			desc:       "failed find words",
			word:       "hallo",
			errMessage: errNotFound.Error(),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := dictionary.Search(tC.word)

			if err != nil {
				if err.Error() != tC.errMessage {
					t.Error("expected error:", tC.errMessage, "got:", err)
				}
			}

			if result != "" {
				if result != tC.expected {
					t.Error("expected:", tC.expected, "got:", result)
				}
			}
		})
	}
}

func TestAddMaps(t *testing.T) {
	dictionary := Dictionary{}
	defer dictionary.ResetDictionary()

	testCases := []struct {
		desc       string
		word       string
		definition string
		errMessage string
	}{
		{
			desc:       "success add word and definiton 1",
			word:       "hallo",
			definition: "text",
		},
		{
			desc:       "success add word and definiton 2",
			word:       "aku",
			definition: "kamu",
		},
		{
			desc:       "failed add definiton existing",
			word:       "hallo",
			definition: "text",
			errMessage: string(errWordExist),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := dictionary.Add(tC.word, tC.definition)

			if err != nil {
				if err.Error() != tC.errMessage {
					t.Error("expected error:", tC.errMessage, "got:", err)
				}
			}
		})
	}
}

func TestUpdateMaps(t *testing.T) {
	dictionary := Dictionary{"hallo": "text"}
	defer dictionary.ResetDictionary()

	testCases := []struct {
		desc          string
		word          string
		newDefinition string
		errMessage    string
	}{
		{
			desc:          "failed update definiton not found",
			word:          "world",
			newDefinition: "text",
			errMessage:    string(errWordDoesNotExist),
		},
		{
			desc:          "success update definition",
			word:          "hallo",
			newDefinition: "new text",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := dictionary.Update(tC.word, tC.newDefinition)

			if err != nil {
				if err.Error() != tC.errMessage {
					t.Error("expected error:", tC.errMessage, "got:", err)
				}
			}
		})
	}
}

func TestDeleteMaps(t *testing.T) {
	dictionary := Dictionary{"hallo": "text"}
	defer dictionary.ResetDictionary()

	testCases := []struct {
		desc       string
		word       string
		errMessage string
	}{
		{
			desc:       "failed delete definiton not found",
			word:       "world",
			errMessage: string(errNotFound),
		},
		{
			desc: "success delete definition",
			word: "hallo",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := dictionary.Delete(tC.word)

			if err != nil {
				if err.Error() != tC.errMessage {
					t.Error("expected error:", tC.errMessage, "got:", err)
				}
			}
		})
	}
}
