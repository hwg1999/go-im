package service

import (
	"errors"
	"fmt"
	"im/model"
	"im/util"
	"math/rand"
	"time"
)

type UserService struct {
}

func (s *UserService) Register(mobile, plainpwd, nickname, avatar, sex string) (model.User, error) {
	user := model.User{}
	_, err := DbEngine.Where("mobile=?", mobile).Get(&user)
	if err != nil {
		return user, err
	}

	if user.Id > 0 {
		return user, errors.New("该手机号已注册")
	}

	user.Mobile = mobile
	user.Avatar = avatar
	user.Nickname = nickname
	user.Sex = sex
	user.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	user.Passwd = util.MakePasswd(plainpwd, user.Salt)
	user.Createat = time.Now()
	user.Token = fmt.Sprintf("%08d", rand.Int31())

	_, err = DbEngine.InsertOne(&user)

	return user, err
}

func (s *UserService) Login(mobile, plainpwd string) (user model.User, err error) {
	return user, nil
}
