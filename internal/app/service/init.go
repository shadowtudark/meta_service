package service

import (
	"meta/internal/app/model"
	"meta/internal/app/service/user"
	"meta/internal/app/service/user/impl"
)

var (
	UserRepository user.Repository
)
// Init instantiate the service
func Init()  {
	UserRepository = impl.NewMysqlImpl(model.MysqlHandler)
}