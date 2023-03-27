package models

import (
	"gorm.io/gorm"
	"time"
)

type SysDept struct {
	DeptId   int    `json:"deptId" gorm:"primaryKey;autoIncrement;"` //部门编码
	ParentId int    `json:"parentId" gorm:""`                        //上级部门
	DeptPath string `json:"deptPath" gorm:"size:255;"`               //
	DeptName string `json:"deptName"  gorm:"size:128;"`              //部门名称
	Sort     int    `json:"sort" gorm:"size:4;"`                     //排序
	Leader   string `json:"leader" gorm:"size:128;"`                 //负责人
	Phone    string `json:"phone" gorm:"size:11;"`                   //手机
	Status   int    `json:"status" gorm:"size:4;"`                   //状态
	ControlBy
	ModelTime
}

func (SysDept) TableName() string {
	return "sys_dept"
}

type SysMenu struct {
	MenuId     int      `json:"menuId" gorm:"primaryKey;autoIncrement"`
	MenuName   string   `json:"menuName" gorm:"size:128;"`
	Title      string   `json:"title" gorm:"size:128;"`
	Icon       string   `json:"icon" gorm:"size:128;"`
	Path       string   `json:"path" gorm:"size:128;"`
	Paths      string   `json:"paths" gorm:"size:128;"`
	MenuType   string   `json:"menuType" gorm:"size:1;"`
	Action     string   `json:"action" gorm:"size:16;"`
	Permission string   `json:"permission" gorm:"size:255;"`
	ParentId   int      `json:"parentId" gorm:"size:11;"`
	NoCache    bool     `json:"noCache" gorm:"size:8;"`
	Breadcrumb string   `json:"breadcrumb" gorm:"size:255;"`
	Component  string   `json:"component" gorm:"size:255;"`
	Sort       int      `json:"sort" gorm:"size:4;"`
	Visible    string   `json:"visible" gorm:"size:1;"`
	IsFrame    string   `json:"isFrame" gorm:"size:1;DEFAULT:0;"`
	SysApi     []SysApi `json:"sysApi" gorm:"many2many:sys_menu_api_rule"`
	ControlBy
	ModelTime
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

// sys_role_dept
type SysRoleDept struct {
	RoleId int `gorm:"size:11;primaryKey"`
	DeptId int `gorm:"size:11;primaryKey"`
}

func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}

type SysUser struct {
	UserId   int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	Username string `json:"username" gorm:"type:varchar(64);comment:用户名"`
	Password string `json:"-" gorm:"type:varchar(128);comment:密码"`
	NickName string `json:"nickName" gorm:"type:varchar(128);comment:昵称"`
	Phone    string `json:"phone" gorm:"type:varchar(11);comment:手机号"`
	Salt     string `json:"-" gorm:"type:varchar(255);comment:加盐"`
	RoleId   int    `json:"roleId" gorm:"type:bigint;comment:角色ID"`
	DeptId   int    `json:"deptId" gorm:"type:bigint;comment:部门"`
	Status   string `json:"status" gorm:"type:varchar(4);comment:状态"`
	ControlBy
	ModelTime
}

func (SysUser) TableName() string {
	return "sys_user"
}

type SysRole struct {
	RoleId    int       `json:"roleId" gorm:"primaryKey;autoIncrement"` // 角色编码
	RoleName  string    `json:"roleName" gorm:"size:128;"`              // 角色名称
	Status    string    `json:"status" gorm:"size:4;"`                  //
	RoleKey   string    `json:"roleKey" gorm:"size:128;"`               //角色代码
	RoleSort  int       `json:"roleSort" gorm:""`                       //角色排序
	Flag      string    `json:"flag" gorm:"size:128;"`                  //
	Remark    string    `json:"remark" gorm:"size:255;"`                //备注
	Admin     bool      `json:"admin" gorm:"size:4;"`
	DataScope string    `json:"dataScope" gorm:"size:128;"`
	SysMenu   []SysMenu `json:"sysMenu" gorm:"many2many:sys_role_menu;foreignKey:RoleId;joinForeignKey:role_id;references:MenuId;joinReferences:menu_id;"`
	ControlBy
	ModelTime
}

func (SysRole) TableName() string {
	return "sys_role"
}

type SysApi struct {
	Id     int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Handle string `json:"handle" gorm:"size:128;comment:handle"`
	Title  string `json:"title" gorm:"size:128;comment:标题"`
	Path   string `json:"path" gorm:"size:128;comment:地址"`
	Type   string `json:"type" gorm:"size:16;comment:接口类型"`
	Action string `json:"action" gorm:"size:16;comment:请求类型"`
	ModelTime
	ControlBy
}

func (SysApi) TableName() string {
	return "sys_api"
}

type ControlBy struct {
	CreateBy int `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy int `json:"updateBy" gorm:"index;comment:更新者"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}
