package file_manager

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileManager(t *testing.T) {

	fm := New(".")
	fileName := "new.txt"
	data := []byte("hello")
	err := fm.Put(fileName, data)

	require.NoError(t, err)

	readData, err := fm.Get(fileName)
	require.NoError(t, err)
	require.NotEmpty(t, readData)
	require.Equal(t, data, readData)

	err = fm.Delete(fileName)
	require.NoError(t, err)

	_, err = fm.IfFileExists(fileName)
	require.Error(t, err)

}
