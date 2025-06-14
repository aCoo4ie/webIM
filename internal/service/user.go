package service

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
	"webIM/internal/models"
	"webIM/internal/pkg"
	"webIM/internal/utils"
)

type UserService struct {
}

var (
	ErrInputError      = errors.New("输入有误")
	ErrPwdNotMatch     = errors.New("密码输入不一致")
	ErrMobileExists    = errors.New("手机号已注册")
	ErrMobileNotExists = errors.New("手机号未注册")
	ErrWrongPwd        = errors.New("密码错误")
)

func (s *UserService) Register(r *models.UserRegister) (*models.User, error) {
	// 输入是否为空
	fmt.Println(r)
	if r.Mobile == "" || r.PlainPwd == "" || r.RePwd == "" {
		return nil, ErrInputError
	}
	// 两次密码输入是否一致
	if r.PlainPwd != r.RePwd {
		return nil, ErrPwdNotMatch
	}
	u := models.User{}
	if pkg.MysqlEngine == nil {
		return nil, errors.New("数据库未初始化")
	}
	_, err := pkg.MysqlEngine.Where("mobile=?", r.Mobile).Get(&u)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	if u.ID > 0 {
		return nil, ErrMobileExists
	}
	// 拼接用户信息
	u.Avatar = r.Avatar
	u.Mobile = r.Mobile
	u.Nickname = r.Nickname
	u.Sex = r.Sex
	u.CreatedAt = time.Now()
	// 加密明文密码
	salt := fmt.Sprintf("%d", rand.Int31n(10000))
	u.Password = utils.HashPassword(r.PlainPwd, salt)

	// 插入数据库中
	_, err = pkg.MysqlEngine.InsertOne(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil

}

func (s *UserService) Login(l models.UserLogin) (*models.User, error) {
	if pkg.MysqlEngine == nil {
		return nil, errors.New("数据库未初始化")
	}
	// 输入是否为空
	if l.Mobile == "" || l.PlainPwd == "" {
		return nil, ErrInputError
	}
	// 用户是否存在
	u := models.User{}
	_, err := pkg.MysqlEngine.Where("mobile=?", l.Mobile).Get(&u)
	if err != nil {
		return nil, ErrMobileNotExists
	}
	// 密码是否正确
	pwd := utils.HashPassword(l.PlainPwd, u.Salt)
	if pwd != u.Password {
		return nil, ErrWrongPwd
	}
	return &u, nil
}
