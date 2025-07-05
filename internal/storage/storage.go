package storage

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
)

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func Fm() {
	err := errors.New("aaa")
	fmt.Println(err)
	err = pgx.ErrNoRows
}
