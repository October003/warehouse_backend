package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"
	"warehouse/models"
)

type ItemSql struct {
	service.Service
}

func (i *ItemSql) CreateItemAndItemID(req *models.Item) error {
	return i.Orm.Transaction(func(tx *gorm.DB) error {
		if err := i.Orm.Create(&models.ItemSql{Item: *req}).Error; err != nil {
			return err
		}
		return i.Orm.Create(&models.ItemIDs{Name: req.ItemID}).Error
	})
}

func (i *ItemSql) DeleteItemAndDetailsByID(id string) error {
	return i.Orm.Transaction(func(tx *gorm.DB) error {
		var itemSql models.ItemSql
		if err := tx.Where("id = ?", id).First(&itemSql).Error; err != nil {
			return err
		}
		if err := tx.Where("item_id = ?", itemSql.ItemID).Unscoped().Delete(&models.InboundSql{}).Error; err != nil {
			return err
		}

		if err := tx.Where("item_id = ?", itemSql.ItemID).Unscoped().Delete(&models.OutboundSql{}).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", id).Unscoped().Delete(&models.ItemSql{}).Error; err != nil {
			return err
		}
		return tx.Where("name = ?", itemSql.ItemID).Unscoped().Delete(&models.ItemIDs{}).Error
	})
}

func (i *ItemSql) Update(newItem *models.Item) error {
	return i.Orm.Transaction(func(tx *gorm.DB) error {
		return tx.Where("item_id = ?", newItem.ItemID).Updates(&models.ItemSql{Item: *newItem}).Error
	})
}

func (i *ItemSql) QueryAll() ([]models.ItemSql, error) {
	var itemSqls []models.ItemSql
	if err := i.Orm.Find(&itemSqls).Error; err != nil {
		return itemSqls, err
	}
	return itemSqls, nil
}
