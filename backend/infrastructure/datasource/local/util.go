package ldb

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/fcfcqloow/first-todo-list/backend/aop"
	"github.com/fcfcqloow/go-advance/log"
)

func load[T any](filePath string, cachekey string) (*T, error) {
	if len(cachekey) > 0 {
		if value, ok := aop.GetCache[T](cachekey); ok {
			log.Debug("get cache [key =", cachekey, "]")
			return value, nil
		}
	}

	log.Debug("read file", filePath)
	byts, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var result T
	if err := json.Unmarshal(byts, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func save[T any](filePath string, value T, cachekey string) error {
	lock.Lock()
	defer lock.Unlock()

	byts, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if len(cachekey) > 0 {
		log.Debug("set cache [key =", cachekey, "]")
		aop.SetCache(cachekey, value)
	}

	log.Debug("write file", filePath)
	return os.WriteFile(filePath, byts, os.ModePerm)
}

func listFiles(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, f := range files {
		if !f.IsDir() {
			filenames = append(filenames, f.Name())
		}
	}
	return filenames, nil
}
