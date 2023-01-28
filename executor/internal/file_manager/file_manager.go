package file_manager

import (
	"io/fs"
	"os"
	"path"
)

type FileManager struct {
	rootDir string
}

func New(rootDir string) *FileManager {
	return &FileManager{
		rootDir: rootDir,
	}
}

func (fm *FileManager) Get(filePath string) ([]byte, error) {
	finalFilePath := path.Join(fm.rootDir, filePath)
	return os.ReadFile(finalFilePath)
}

func (fm *FileManager) Put(filePath string, data []byte) error {
	finalFilePath := path.Join(fm.rootDir, filePath)
	file, err := os.Create(finalFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (fm *FileManager) IfFileExists(filePath string) (fs.FileInfo, error) {
	finalFilePath := path.Join(fm.rootDir, filePath)
	return os.Stat(finalFilePath)
}

func (fm *FileManager) Delete(filePath string) error {
	finalFilePath := path.Join(fm.rootDir, filePath)
	return os.Remove(finalFilePath)
}
