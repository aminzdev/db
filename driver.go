package db

type Driver interface {
	Connect() (*DB, error)
}
