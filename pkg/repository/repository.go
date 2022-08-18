package repository

type Balance interface {
}

type Repository struct {
	Balance
}

func NewRepository() *Repository {
	return &Repository{}
}
