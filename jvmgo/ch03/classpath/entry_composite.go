package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(path string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(path, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (dir CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range dir {
		data, form, err := entry.readClass(className)
		if err == nil {
			return data, form, nil
		}
	}
	return nil, nil, errors.New("class not found")
}

func (dir CompositeEntry) String() string {
	strs := make([]string, len(dir))
	for i, entry := range dir {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
