package second

import (
	"database/sql"

	"github.com/pkg/errors"
)

//应该往上抛，需要抛一个自定义错误，交给业务曾处理
var (
	ErrRecordNotFound = errors.New("record not found")
)

type User struct {
	ID   int64
	Name string
}

func (u *User) list() ([]User, error) {

	var rows []User

	err := sql.QueryRow("SELECT * FROM user WHERE id IN ?", []int64{1, 2, 3}).Scan(&rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return rows, errors.Warp(ErrRecordNotFound, "not found...")
		}
	}
	return rows, nil
}
