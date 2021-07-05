package repository

type Authorization interface {
}

type Cdb interface {
}

type Repository struct {
	Authorization
	Cdb
}

func NewRepository() *Repository {
	return &Repository{}
}
