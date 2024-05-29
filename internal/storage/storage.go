package storage

import "github.com/Bostanova/route256_tgbot/internal/file"

type Storage struct{}

// NewStorage создает объект типа Storage и возвращает его
func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}
	return newFile, nil
}
