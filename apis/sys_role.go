package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"net/http"
	"warehouse/common/global"
	"warehouse/models"
	"warehouse/service"
	"warehouse/service/dto"
)

type SysRole struct {
	api.Api
}

// Insert
// @Summary 创建角色
func (e SysRole) Insert(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleInsertReq{}
	err := e.MakeContext(c).MakeOrm().Bind(&req, binding.JSON).MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.CreateBy = user.GetUserId(c)
	if req.Status == "" {
		req.Status = "2"
	}
	cb := sdk.Runtime.GetCasbinKey(c.Request.Host)
	err = s.Insert(&req, cb)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "创建失败,"+err.Error())
		return
	}
	_, err = global.LoadPolicy(c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "创建失败,"+err.Error())
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户角色
// @Summary 修改用户角色
func (e SysRole) Update(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleUpdateReq{}
	err := e.MakeContext(c).MakeOrm().Bind(&req, nil, binding.JSON).MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	cb := sdk.Runtime.GetCasbinKey(c.Request.Host)
	req.SetUpdateBy(user.GetUserId(c))
	if err = s.Update(&req, cb); err != nil {
		e.Logger.Error(err)
		return
	}
	if _, err = global.LoadPolicy(c); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "更新失败,"+err.Error())
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除用户角色
func (e SysRole) Delete(c *gin.Context) {
	s := new(service.SysRole)
	req := dto.SysRoleDeleteReq{}
	err := e.MakeContext(c).MakeOrm().Bind(&req, binding.JSON).MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, fmt.Sprintf("删除角色 %v 失败，\r\n失败信息 %s", req.Ids, err.Error()))
		return
	}
	cb := sdk.Runtime.GetCasbinKey(c.Request.Host)
	if err = s.Remove(&req, cb); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "")
		return
	}
	e.OK(req.GetId(), fmt.Sprintf("删除角色角色 %v 成功！", req.GetId()))
}

// Update2Status 修改用户角色状态
// @Summary 修改用户角色
func (e SysRole) Update2Status(c *gin.Context) {
	s := service.SysRole{}
	req := dto.UpdateStatusReq{}
	err := e.MakeContext(c).MakeOrm().Bind(&req, binding.JSON, nil).MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, fmt.Sprintf("更新角色状态失败，失败原因：%s ", err.Error()))
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	if err = s.UpdateStatus(&req); err != nil {
		e.Error(500, err, fmt.Sprintf("更新角色状态失败，失败原因：%s ", err.Error()))
		return
	}
	e.OK(req.GetId(), fmt.Sprintf("更新角色 %v 状态成功！", req.GetId()))
}

// Update2DataScope 更新角色数据权限
// @Summary 更新角色数据权限
func (e SysRole) Update2DataScope(c *gin.Context) {
	s := service.SysRole{}
	req := dto.RoleDataScopeReq{}
	err := e.MakeContext(c).MakeOrm().Bind(&req, binding.JSON, nil).MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	data := &models.SysRole{
		RoleId:    req.RoleId,
		DataScope: req.DataScope,
		DeptIds:   req.DeptIds,
	}
	data.UpdateBy = user.GetUserId(c)
	if err = s.UpdateDataScope(&req).Error; err != nil {
		e.Error(500, err, fmt.Sprintf("更新角色数据权限失败！错误详情：%s", err.Error()))
		return
	}
	e.OK(nil, "操作成功")
}

// GetPage
// @Summary 角色列表数据
func (e SysRole) GetPage(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleGetPageReq{}
	err := e.MakeContext(c).MakeOrm().Bind(&req, binding.Form).MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	list := make([]models.SysRole, 0)
	var count int64
	if err = s.GetPage(&req, &list, &count); err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get
// @Summary 获取Role数据
func (e SysRole) Get(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleGetReq{}
	err := e.MakeContext(c).MakeOrm().Bind(&req, nil).MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, fmt.Sprintf(" %s ", err.Error()))
		return
	}
	var object models.SysRole
	if err = s.Get(&req, &object); err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}
