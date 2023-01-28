package object_store

import (
	"io"
	"os"
	"path"
)

type ObjectStore struct {
	location string
}

func NewObjectStore(location string) (*ObjectStore, error) {
	return &ObjectStore{
		location: location,
	}, os.MkdirAll(location, os.ModePerm)
}

func (s *ObjectStore) Put(data io.Reader, name string) (string, error) {
	path := path.Join(s.location, name)
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, data); err != nil {
		return "", err
	}
	return "local@" + path, nil
}

func (s *ObjectStore) Get(name string) (io.ReadCloser, error) {
	path := path.Join(s.location, name)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *ObjectStore) Delete(name string) error {
	path := path.Join(s.location, name)
	return os.Remove(path)
}
