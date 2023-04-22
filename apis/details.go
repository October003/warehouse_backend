package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"net/http"
	"warehouse/models"
	"warehouse/service"
	"warehouse/service/dto"
)

type Details struct {
	api.Api
}

// 创建入库记录
func (e Details) CreateInbound(c *gin.Context) {
	s := service.Inbound{}
	req := models.Details{}
	err := e.MakeContext(c).MakeOrm().Bind(&req, binding.JSON).MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.CreateInbound(&req)
	if err != nil {
		e.Logger.Errorf("s.CreateInbound() failed,err:%s", err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "添加成功")
}

// 查询入库记录
func (e Details) QueryInbound(c *gin.Context) {
	s := service.Inbound{}
	datalist := []dto.DetailQueryReq{}
	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	datalist, err = s.QueryAllInbound(&datalist)
	if err != nil {
		e.Logger.Errorf("s.QueryAllInbound() failed,err:%s", err)
		e.Error(500, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, &gin.H{
		"details": datalist,
		"msg":     "查询成功",
	})
}

// 删除入库记录
func (e Details) DeleteInbound(c *gin.Context) {
	id := c.Query("id")
	s := service.Inbound{}
	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.DeleteInboundById(id)
	if err != nil {
		e.Logger.Errorf("s.DeleteInboundById() failed.err:%s", err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "删除成功")
}

// 添加出库记录
func (e Details) CreateOutbound(c *gin.Context) {
	s := service.Outbound{}
	req := models.Details{}
	err := e.MakeContext(c).MakeOrm().Bind(&req, binding.JSON).MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.CreateOutbound(&req)
	if err != nil {
		e.Logger.Errorf("s.CreateInbound failed,err:%s", err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "添加成功")
}

// 查询出库记录
func (e Details) QueryOutbound(c *gin.Context) {
	s := service.Outbound{}
	datalist := []dto.DetailQueryReq{}
	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	datalist, err = s.QueryAllOutbound(&datalist)
	if err != nil {
		e.Logger.Errorf("s.QueryAllInbound() failed,err:%s", err)
		e.Error(500, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, &gin.H{
		"details": datalist,
		"msg":     "查询成功",
	})
}

// 删除出库记录
func (e Details) DeleteOutbound(c *gin.Context) {
	id := c.Query("id")
	s := service.Outbound{}
	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.DeleteOutboundById(id)
	if err != nil {
		e.Logger.Errorf("s.DeleteInboundById() failed.err:%s", err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "删除成功")
}

func (e Details) QueryByTimestamp(c *gin.Context) {
	start := c.Query("start")
	end := c.Query("end")
	s := service.QueryByTime{}
	//req := dto.DetailQueryReq{}
	err := e.MakeContext(c).MakeOrm().MakeService(&s.Service).Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	inboundDetails, outboundDetails, err := s.QueryByTimestamp(start, end)
	if err != nil {
		e.Logger.Errorf("s.QueryByTime() failed,err:%s", err)
		e.Error(500, err, err.Error())
		return
	}
	//datalist := [][]dto.DetailQueryReq{inboundDetails, outboundDetails}
	//e.OK(datalist, "查询成功")
	c.JSON(http.StatusOK, gin.H{
		"msg":      "查询成功",
		"inbound":  inboundDetails,
		"outbound": outboundDetails,
	})
}
