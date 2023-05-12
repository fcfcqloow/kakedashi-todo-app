package ldb

import (
	"encoding/base64"
	"os"
	"path/filepath"

	"github.com/fcfcqloow/first-todo-list/backend/domain/generated"

	"github.com/fcfcqloow/go-advance/log"
)

var (
	UserHome, userHomeErr = os.UserHomeDir()
	AppDir                = filepath.Join(UserHome, ".todoapp")
	StaticDir             = filepath.Join(AppDir, "static")
	DataDir               = filepath.Join(AppDir, "data")
	LogDir                = filepath.Join(AppDir, "log")
	BackupDir             = filepath.Join(DataDir, "backup")
	RestoreDir            = filepath.Join(DataDir, "restore")
	TodoDbFile            = filepath.Join(DataDir, "todo.db")
	TopicDbFile           = filepath.Join(DataDir, "topic.db")
	SettingDbFile         = filepath.Join(DataDir, "settings.db")
	MemoFile              = filepath.Join(DataDir, "memo.db")
)

func Install() error {
	if userHomeErr != nil {
		return userHomeErr
	}

	safeMkdir(AppDir, os.ModePerm)
	safeMkdir(StaticDir, os.ModePerm)
	safeMkdir(DataDir, os.ModePerm)
	safeMkdir(BackupDir, os.ModePerm)
	safeMkdir(LogDir, os.ModePerm)
	safeMkdir(RestoreDir, os.ModePerm)
	safeCreateFile(TodoDbFile, ptrString("{}"))
	safeCreateFile(TopicDbFile, ptrString("[]"))
	safeCreateFile(SettingDbFile, ptrString("{}"))

	for distPath, fileValue := range generated.FrontendDist {
		dir, file := filepath.Split(distPath)
		staticFile, err := base64.StdEncoding.DecodeString(fileValue)
		if err != nil {
			panic(err)
		}

		safeMkdir(filepath.Join(StaticDir, dir), os.ModePerm)
		safeCreateFile(filepath.Join(StaticDir, dir, file), ptrString(string(staticFile)))
	}

	return nil
}

func safeMkdir(dir string, perm os.FileMode) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, perm); err != nil {
			log.Error(dir, err)
		} else {
			log.Info("Create " + dir)
		}
	}
}

func safeCreateFile(filePath string, value *string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if file, err := os.Create(filePath); err != nil {
			log.Error(filePath, err)
		} else {
			defer file.Close()
			log.Info("Create", filePath)
			if value == nil {
				return
			}
			if _, err := file.WriteString(*value); err != nil {
				log.Error(err)
			}
		}
	}
}

func ptrString(value string) *string {
	return &value
}
