package repository

import "github.com/jphacks/F_2002_1/go/domain/entity"

// User はユーザの永続化を担当するリポジトリです。
type User interface {
	FindByID(id int) (*entity.User, error)
	FindIDByUID(uid int) (*entity.Id, error)
	FindAll() (*entity.Users, error)
	Store(user *entity.User)  (*entity.User, error)
	UpdateByID(id int, user *entity.User) (*entity.User, error)
	DeleteByID(id int) (*entity.User, error)
}
