package storage

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
)

func WriteFile(fileName, filePath string, content []byte) error {
	err := os.MkdirAll(filePath, 0777)
	if err != nil {
		return errors.WithStack(err)
	}

	err = ioutil.WriteFile(path.Join(filePath, fileName), []byte(content), 0644)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
