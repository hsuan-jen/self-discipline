package user_service

import (
	"self-discipline/model"
)

var _ Service = &service{}

type Service interface {
	i()
	Login(u *model.Users) (err error, userInter *model.Users)
	GetRedisJWT(userName string) (err error, redisJWT string)
	SetRedisJWT(jwt string, userName string) (err error)
	SaveUserRedis(u *model.Users) error
	Register(u *model.Users) (err error, userInter *model.Users)
}

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) i() {}
