package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"net/http"
	"warehouse/models"
	"warehouse/service"
)

type Material struct {
	api.Api
}

// 添加物料信息
func (e Material) Create(c *gin.Context) {
	s := service.ItemSql{}
	req := models.Item{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	if err := s.CreateItemAndItemID(&req); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "添加成功")
}

// 获取物料信息
func (e Material) Query(c *gin.Context) {
	s := service.ItemSql{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var list []models.ItemSql
	list, err = s.QueryAll()
	if err != nil {
		e.Logger.Errorf("s.Query() failed,db err:%s", err)
		e.Error(500, err, "查询失败")
		return
	}
	c.JSON(http.StatusOK, &gin.H{
		"items": list,
		"msg":   "查询成功",
	})
}

// 修改物料信息
func (e Material) Update(c *gin.Context) {
	s := service.ItemSql{}
	req := models.Item{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	if err := s.Update(&req); err != nil {
		e.Logger.Errorf("s.Update() failed,db err:%s", err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "更新成功")
}

// 删除物料信息
func (e Material) Delete(c *gin.Context) {
	id := c.Query("id")
	s := service.ItemSql{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	if err := s.DeleteItemAndDetailsByID(id); err != nil {
		e.Logger.Errorf("s.DeleteMaterialDetailById() failed,err:%s", err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "删除成功")
}
