package Getter

import (
	"fmt"
	"ginSkill/src/data/mappers"
	"ginSkill/src/models/UserModel"
	"ginSkill/src/result"
)

var UserGetter IUserGetter

func init() {
	UserGetter = NewUserGetterImpl()
}

type IUserGetter interface {
	GetUserListByPage(page, size int) []*UserModel.UserModelImpl
	GetUserByID(id int) *result.ErrorResult
}

type UserGetterImpl struct {
	userMapper *mappers.UserMapper
}

func NewUserGetterImpl() *UserGetterImpl {
	return &UserGetterImpl{&mappers.UserMapper{}}
}

func (this *UserGetterImpl) GetUserListByPage(page, size int) []*UserModel.UserModelImpl {
	us := make([]*UserModel.UserModelImpl, 0)
	//dbs.Orm.Limit(size).Offset(page*size - size).Find(&us)
	this.userMapper.GetUserListByPage(page, size).Query().Find(&us)
	return us
}

func (this *UserGetterImpl) GetUserByID(id int) *result.ErrorResult {
	u := UserModel.New()
	if this.userMapper.GetUserByID(id).Query().Find(&u).RecordNotFound() {
		return &result.ErrorResult{
			Err:  fmt.Errorf("not found user, id: %d", id),
			Data: nil,
		}
	} else {
		return &result.ErrorResult{
			Err:  nil,
			Data: u,
		}
	}
}
