package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"warehouse/common/models"
)

type SysUser struct {
	UserId   int      `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	Username string   `json:"username" gorm:"size:64;comment:用户名"`
	Password string   `json:"-" gorm:"size:128;comment:密码"`
	NickName string   `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone    string   `json:"phone" gorm:"size:11;comment:手机号"`
	Salt     string   `json:"-" gorm:"size:255;comment:加盐"`
	Status   string   `json:"status" gorm:"size:4;comment:状态"`
	RoleId   int      `json:"roleId" gorm:"size:20;comment:角色ID"`
	DeptId   int      `json:"deptId" gorm:"size:20;comment:部门"`
	RoleIds  []int    `json:"roleIds" gorm:"-"`
	DeptIds  []int    `json:"deptIds" gorm:"-"`
	Dept     *SysDept `json:"dept"`
	ControlBy
	ModelTime
}

func (SysUser) TableName() string {
	return "sys_user"
}

func (e *SysUser) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysUser) GetId() interface{} {
	return e.UserId
}

// 加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *SysUser) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *SysUser) BeforeUpdate(_ *gorm.DB) error {
	var err error
	if e.Password != "" {
		err = e.Encrypt()
	}
	return err
}

//func (e *SysUser) AfterFind(_ *gorm.DB) error {
//	e.DeptIds = []int{e.DeptId}
//	e.RoleIds = []int{e.RoleId}
//	return nil
//}
