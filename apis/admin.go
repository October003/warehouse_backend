package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"warehouse/service"
)

type Admin struct {
	api.Api
}

// 添加系统设置
func (e Admin) Add(c *gin.Context) {
	inboundPersons := c.Query("inbound_persons")
	outboundPersons := c.Query("outbound_persons")
	units := c.Query("units")
	storageLocations := c.Query("storage_locations")
	s := service.Admin{}
	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
	}
	if err := s.Add(inboundPersons, 1); err != nil {
		e.Error(500, err, err.Error())
	}
	if err := s.Add(outboundPersons, 2); err != nil {
		e.Error(500, err, err.Error())
	}
	if err := s.Add(units, 3); err != nil {
		e.Error(500, err, err.Error())
	}
	if err := s.Add(storageLocations, 4); err != nil {
		e.Error(500, err, err.Error())
	}
	config, err := s.GetAll()
	if err != nil {
		e.Logger.Errorf("s.GetAll() failed,db err:%s", err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(config, "添加系统设置成功")
}

// 删除系统设置
func (e Admin) Delete(c *gin.Context) {
	inboundPersons := c.Query("inbound_persons")
	outboundPersons := c.Query("outbound_persons")
	units := c.Query("units")
	storageLocations := c.Query("storage_locations")
	s := service.Admin{}
	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	if err := s.Delete(inboundPersons, 1); err != nil {
		e.Error(500, err, err.Error())
	}
	if err := s.Delete(outboundPersons, 2); err != nil {
		e.Error(500, err, err.Error())
	}
	if err := s.Delete(units, 3); err != nil {
		e.Error(500, err, err.Error())
	}
	if err := s.Delete(storageLocations, 4); err != nil {
		e.Error(500, err, err.Error())
	}
	config, err := s.GetAll()
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(config, "删除系统设置成功")
}

// 获取系统配置
func (e Admin) Get(c *gin.Context) {
	s := service.Admin{}
	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	config, err := s.GetAll()
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(config, "获取系统设置成功")
}
