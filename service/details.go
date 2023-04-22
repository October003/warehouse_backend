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
	var ItemSql models.ItemSql
	return i.Orm.Transaction(func(tx *gorm.DB) error {
		// 根据物料ID 在数据库中查找物料信息
		err := tx.Where("item_id = ? ", data.ItemID).First(&ItemSql).Error
		if err != nil {
			i.Log.Errorf("query item information failed,db err:%s", err)
			return err
		}
		// 物料数量
		materialNumber, err := decimal.NewFromString(data.Number)
		if err != nil {
			return err
		}
		// 物料结存数量
		materialResNumber, err := decimal.NewFromString(ItemSql.Number)
		if err != nil {
			return err
		}
		//物料入库数量
		materialInNumber, err := decimal.NewFromString(ItemSql.InNumber)
		if err != nil {
			return err
		}
		//物料结存数量 =  结存数 + 入库数量
		ItemSql.Number = materialResNumber.Add(materialNumber).String()
		//物料总入库数 = 入库数 + 入库数量
		ItemSql.InNumber = materialInNumber.Add(materialNumber).String()
		// 更新 物料信息
		err = tx.Where("item_id = ? ", data.ItemID).Updates(&ItemSql).Error
		if err != nil {
			i.Log.Errorf("Updates material information failed,db err:%s", err)
			return err
		}
		return tx.Create(&models.InboundSql{Details: *data}).Error
	})
}
func (o *Outbound) CreateOutbound(data *models.Details) error {
	var materialSql models.ItemSql
	return o.Orm.Transaction(func(tx *gorm.DB) error {
		// 根据物料ID 在item_sql 中查找该物料的信息
		err := tx.Where("item_id = ? ", data.ItemID).First(&materialSql).Error
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
		materialResNumber, err := decimal.NewFromString(materialSql.Number)
		if err != nil {
			return err
		}
		//物料入库数量
		materialOutNumber, err := decimal.NewFromString(materialSql.OutNumber)
		if err != nil {
			return err
		}
		//物料结存数量 =  结存数 - 出库数量
		materialSql.Number = materialResNumber.Sub(materialNumber).String()
		//物料总出库数 = 出库数 + 出库数量
		materialSql.OutNumber = materialOutNumber.Add(materialNumber).String()
		// 更新 物料信息
		err = tx.Where("item_id = ? ", data.ItemID).Updates(&materialSql).Error
		if err != nil {
			o.Log.Errorf("Updates item information failed,db err:%s", err)
			return err
		}
		return tx.Create(&models.OutboundSql{Details: *data}).Error
	})
}
func (i *Inbound) QueryAllInbound(datalist *[]dto.DetailQueryReq) ([]dto.DetailQueryReq, error) {
	//根据 QueryReq 中的字段  去 inbound_sql 和 item_sql  表中查询 material_id 相等的字段
	i.Orm.Model(&models.InboundSql{}).
		Select("inbound_sql.*,item_sql.item_id,item_sql.item_name,item_sql.specification,item_sql.unit,item_sql.storage_location").
		Joins("JOIN item_sql ON inbound_sql.item_id = item_sql.item_id").Scan(datalist)
	return *datalist, nil
}

func (o *Outbound) QueryAllOutbound(datalist *[]dto.DetailQueryReq) ([]dto.DetailQueryReq, error) {
	//根据 QueryReq 中的字段  去 outbound_detail_sql 和 material_sql  表中查询 material_id 相等的字段
	o.Orm.Model(&models.OutboundSql{}).
		Select("outbound_sql.*,item_sql.item_id,item_sql.item_name,item_sql.specification,item_sql.unit,item_sql.storage_location").
		Joins("JOIN item_sql ON outbound_sql.item_id = item_sql.item_id").Scan(datalist)
	return *datalist, nil
}

func (i *Inbound) DeleteInboundById(id string) error {
	return i.Orm.Transaction(func(tx *gorm.DB) error {
		var (
			Item    models.Item
			Inbound models.InboundSql
		)
		//根据 id 在 inbound_sql 表中查找到这条入库记录
		if err := tx.Where("id = ? ", id).First(&Inbound).Error; err != nil {
			return err
		}
		// 根据 item_id 在item_sql 表中找到 该物料的信息
		if err := tx.Where("item_id = ? ", Inbound.ItemID).First(&Item).Error; err != nil {
			return err
		}
		// 入库的数量
		Number, err := decimal.NewFromString(Inbound.Number)
		if err != nil {
			return err
		}
		ResNumber, err := decimal.NewFromString(Item.Number)
		if err != nil {
			return err
		}
		InboundNumber, err := decimal.NewFromString(Item.InNumber)
		if err != nil {
			return err
		}
		Item.InNumber = InboundNumber.Sub(Number).String()
		Item.Number = ResNumber.Sub(Number).String()
		if err := tx.Where("item_id = ? ", Item.ItemID).Updates(&Item).Error; err != nil {
			return err
		}
		return tx.Where("id = ? ", id).Unscoped().Delete(&Inbound).Error
	})
}

func (o *Outbound) DeleteOutboundById(id string) error {
	return o.Orm.Transaction(func(tx *gorm.DB) error {
		var (
			Item     models.ItemSql
			Outbound models.OutboundSql
		)
		//根据 id 在 inbound_detail_sql 表中查找到这条入库记录
		if err := tx.Where("id = ? ", id).First(&Outbound).Error; err != nil {
			return err
		}
		// 根据 material_id 在material_sql 表中找到 该物料的信息
		if err := tx.Where("item_id = ? ", Outbound.ItemID).First(&Item).Error; err != nil {
			return err
		}
		// 出库的数量
		Number, err := decimal.NewFromString(Outbound.Number)
		if err != nil {
			return err
		}
		ResNumber, err := decimal.NewFromString(Item.Number)
		if err != nil {
			return err
		}
		OutboundNumber, err := decimal.NewFromString(Item.OutNumber)
		if err != nil {
			return err
		}
		// 总出库数量  = 总出库数 - 出库的数量
		Item.OutNumber = OutboundNumber.Sub(Number).String()
		// 结存数量 = 结存数 + 出库数
		Item.Number = ResNumber.Add(Number).String()
		if err := tx.Where("item_id = ? ", Item.ItemID).Updates(&Item).Error; err != nil {
			return err
		}
		return tx.Where("id = ? ", id).Unscoped().Delete(&Outbound).Error
	})
}

// 时间查询 通过时间戳 去 inbound_sql outbound_sql 表中查询出入库明细

func (q *QueryByTime) QueryByTimestamp(start, end string) ([]dto.DetailQueryReq, []dto.DetailQueryReq, error) {
	var (
		inBoundDetails  []dto.DetailQueryReq
		outBoundDetails []dto.DetailQueryReq
	)
	q.Orm.Model(&models.InboundSql{}).
		Select("inbound_sql.*,item_sql.item_id,item_sql.item_name,item_sql.specification,item_sql.unit,item_sql.storage_location").
		Joins("JOIN item_sql ON inbound_sql.item_id = item_sql.item_id").
		Where("timestamp >= ?", start).
		Where("timestamp <= ?", end).
		Scan(&inBoundDetails)

	q.Orm.Model(&models.OutboundSql{}).
		Select("outbound_sql.*,item_sql.item_id,item_sql.item_name,item_sql.specification,item_sql.unit,item_sql.storage_location").
		Joins("JOIN item_sql ON outbound_sql.item_id = item_sql.item_id").
		Where("timestamp >= ?", start).
		Where("timestamp <= ?", end).
		Scan(&outBoundDetails)

	return inBoundDetails, outBoundDetails, nil
}
