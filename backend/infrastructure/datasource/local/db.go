package ldb

type localRepository struct{}

func NewLocalRepository() *localRepository {
	return &localRepository{}
}
