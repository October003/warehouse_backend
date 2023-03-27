package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"warehouse/models"
	"warehouse/service/dto"
)

type Inbound struct {
	service.Service
}
type Outbound struct {
	service.Service
}
type QueryByTime struct {
	service.Service
}

func (i *Inbound) CreateInbound(data *models.Details) error {
	var materialSql models.MaterialSql
	return i.Orm.Transaction(func(tx *gorm.DB) error {
		// 根据物料ID 在数据库中查找物料信息
		err := tx.Where("material_id = ? ", data.MaterialID).First(&materialSql).Error
		if err != nil {
			i.Log.Errorf("query material information failed,db err:%s", err)
			return err
		}
		// 物料数量
		materialNumber, err := decimal.NewFromString(data.Number)
		if err != nil {
			return err
		}
		// 物料结存数量
		materialResNumber, err := decimal.NewFromString(materialSql.ResNumber)
		if err != nil {
			return err
		}
		//物料入库数量
		materialInNumber, err := decimal.NewFromString(materialSql.InNumber)
		if err != nil {
			return err
		}
		//物料结存数量 =  结存数 + 入库数量
		materialSql.ResNumber = materialResNumber.Add(materialNumber).String()
		//物料总入库数 = 入库数 + 入库数量
		materialSql.InNumber = materialInNumber.Add(materialNumber).String()
		// 更新 物料信息
		err = tx.Where("material_id = ? ", data.MaterialID).Updates(&materialSql).Error
		if err != nil {
			i.Log.Errorf("Updates material information failed,db err:%s", err)
			return err
		}
		return tx.Create(&models.InboundDetailSql{Details: *data}).Error
	})
}
func (o *Outbound) CreateOutbound(data *models.Details) error {
	var materialSql models.MaterialSql
	return o.Orm.Transaction(func(tx *gorm.DB) error {
		// 根据物料ID 在material_sql 中查找该物料的信息
		err := tx.Where("material_id = ? ", data.MaterialID).First(&materialSql).Error
		if err != nil {
			o.Log.Errorf("query material information failed,db err:%s", err)
			return err
		}
		// 物料数量
		materialNumber, err := decimal.NewFromString(data.Number)
		if err != nil {
			return err
		}
		// 物料结存数量
		materialResNumber, err := decimal.NewFromString(materialSql.ResNumber)
		if err != nil {
			return err
		}
		//物料入库数量
		materialOutNumber, err := decimal.NewFromString(materialSql.OutNumber)
		if err != nil {
			return err
		}
		//物料结存数量 =  结存数 - 出库数量
		materialSql.ResNumber = materialResNumber.Sub(materialNumber).String()
		//物料总出库数 = 出库数 + 出库数量
		materialSql.OutNumber = materialOutNumber.Add(materialNumber).String()
		// 更新 物料信息
		err = tx.Where("material_id = ? ", data.MaterialID).Updates(&materialSql).Error
		if err != nil {
			o.Log.Errorf("Updates material information failed,db err:%s", err)
			return err
		}
		return tx.Create(&models.OutboundDetailSql{Details: *data}).Error
	})
}
func (i *Inbound) QueryAllInbound(datalist *[]dto.DetailQueryReq) ([]dto.DetailQueryReq, error) {
	//根据 QueryReq 中的字段  去 inbound_detail_sql 和 material_sql  表中查询 material_id 相等的字段
	i.Orm.Model(&models.InboundDetailSql{}).
		Select("inbound_detail_sql.*,material_sql.material_id,material_sql.material_name,material_sql.specification,material_sql.unit,material_sql.storage_location").
		Joins("JOIN material_sql ON inbound_detail_sql.material_id = material_sql.material_id").Scan(datalist)
	return *datalist, nil
}

func (o *Outbound) QueryAllOutbound(datalist *[]dto.DetailQueryReq) ([]dto.DetailQueryReq, error) {
	//根据 QueryReq 中的字段  去 outbound_detail_sql 和 material_sql  表中查询 material_id 相等的字段
	o.Orm.Model(&models.OutboundDetailSql{}).
		Select("outbound_detail_sql.*,material_sql.material_id,material_sql.material_name,material_sql.specification,material_sql.unit,material_sql.storage_location").
		Joins("JOIN material_sql ON outbound_detail_sql.material_id = material_sql.material_id").Scan(datalist)
	return *datalist, nil
}
func (i *Inbound) DeleteInboundById(id string) error {
	return i.Orm.Transaction(func(tx *gorm.DB) error {
		var (
			material models.MaterialSql
			Inbound  models.InboundDetailSql
		)
		//根据 id 在 inbound_detail_sql 表中查找到这条入库记录
		if err := tx.Where("id = ? ", id).First(&Inbound).Error; err != nil {
			return err
		}
		// 根据 material_id 在material_sql 表中找到 该物料的信息
		if err := tx.Where("material_id = ? ", Inbound.MaterialID).First(&material).Error; err != nil {
			return err
		}
		// 入库的数量
		Number, err := decimal.NewFromString(Inbound.Number)
		if err != nil {
			return err
		}
		ResNumber, err := decimal.NewFromString(material.ResNumber)
		if err != nil {
			return err
		}
		InboundNumber, err := decimal.NewFromString(material.InNumber)
		if err != nil {
			return err
		}
		material.InNumber = InboundNumber.Sub(Number).String()
		material.ResNumber = ResNumber.Sub(Number).String()
		if err := tx.Where("material_id = ? ", material.MaterialID).Updates(&material).Error; err != nil {
			return err
		}
		return tx.Where("id = ? ", id).Unscoped().Delete(&Inbound).Error
	})
}

func (o *Outbound) DeleteOutboundById(id string) error {
	return o.Orm.Transaction(func(tx *gorm.DB) error {
		var (
			material models.MaterialSql
			Outbound models.OutboundDetailSql
		)
		//根据 id 在 inbound_detail_sql 表中查找到这条入库记录
		if err := tx.Where("id = ? ", id).First(&Outbound).Error; err != nil {
			return err
		}
		// 根据 material_id 在material_sql 表中找到 该物料的信息
		if err := tx.Where("material_id = ? ", Outbound.MaterialID).First(&material).Error; err != nil {
			return err
		}
		// 出库的数量
		Number, err := decimal.NewFromString(Outbound.Number)
		if err != nil {
			return err
		}
		ResNumber, err := decimal.NewFromString(material.ResNumber)
		if err != nil {
			return err
		}
		OutboundNumber, err := decimal.NewFromString(material.OutNumber)
		if err != nil {
			return err
		}
		// 总出库数量  = 总出库数 - 出库的数量
		material.OutNumber = OutboundNumber.Sub(Number).String()
		// 结存数量 = 结存数 + 出库数
		material.ResNumber = ResNumber.Add(Number).String()
		if err := tx.Where("material_id = ? ", material.MaterialID).Updates(&material).Error; err != nil {
			return err
		}
		return tx.Where("id = ? ", id).Unscoped().Delete(&Outbound).Error
	})
}

// 时间查询 通过时间戳 去 inbound_detail_sql outbound_detail_sql 表中查询出入库明细

func (q *QueryByTime) QueryByTime(start, end string) ([]dto.DetailQueryReq, []dto.DetailQueryReq, error) {
	var (
		inboundDetails  []dto.DetailQueryReq
		outboundDetails []dto.DetailQueryReq
	)
	if err := q.Orm.Model(&models.InboundDetailSql{}).
		Select("inbound_detail_sql.*,material_sql.material_id,material_sql.material_name,material_sql.specification,material_sql.unit,material_sql.storage_location").
		Joins("JOIN material_sql ON inbound_detail_sql.material_id = material_sql.material_id").
		Where("timestamp >= ?", start).
		Where("timestamp <= ?", end).
		Scan(&inboundDetails).Error; err != nil {
		q.Log.Errorf("query inboundDetail failed,db err:%s", err)
		return nil, nil, err
	}
	if err := q.Orm.Model(&models.OutboundDetailSql{}).
		Select("outbound_detail_sql.*,material_sql.material_id,material_sql.material_name,material_sql.specification,material_sql.unit,material_sql.storage_location").
		Joins("JOIN material_sql ON outbound_detail_sql.material_id = material_sql.material_id").
		Where("timestamp >= ? AND timestamp <= ?", start, end).Error; err != nil {
		q.Log.Errorf("query outboundDetail failed,db err:%s", err)
		return nil, nil, err
	}
	return inboundDetails, outboundDetails, nil
}
