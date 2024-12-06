package repository

import (
	"github.com/tidwall/buntdb"
)

const (
	id = "key"
)

type (
	KeyRepository interface {
		Create(key string) error
		Get() (*string, error)
		Close()
	}

	keyRepository struct {
		db *buntdb.DB
	}
)

func NewKeyRepository(dbPath string) (KeyRepository, error) {
	db, err := buntdb.Open(dbPath)
	if err != nil {
		return nil, err
	}
	return &keyRepository{db: db}, nil
}

func (r *keyRepository) Create(key string) error {
	return r.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(id, key, nil)
		return err
	})
}

func (r *keyRepository) Get() (*string, error) {
	var result *string

	err := r.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(id)
		if err != nil {
			return err
		}
		result = &val
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *keyRepository) Close() {
	if err := r.db.Close(); err != nil {
		panic(err)
	}
}
