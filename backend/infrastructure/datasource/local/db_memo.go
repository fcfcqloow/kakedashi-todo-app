package ldb

var memoCacheKey = "memo-cache-key"

func (t *localRepository) GetMemo() (string, error) {
	settings, err := load[string](MemoFile, memoCacheKey)
	if err != nil {
		return "", err
	}

	return *settings, nil
}

func (t *localRepository) SyncMemo(value string) error {
	if err := save(MemoFile, value, memoCacheKey); err != nil {
		return err
	}

	return nil
}
