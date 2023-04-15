package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"
	"warehouse/models"
)

type MaterialSql struct {
	service.Service
}

func (m *MaterialSql) Create(req *models.MaterialSql) error {
	if err := m.Orm.Create(req).Error; err != nil {
		m.Log.Errorf("Create failed , db err :%s", err)
		return err
	}
	return nil
}

func (m *MaterialSql) DeleteMaterialDetailById(id string) error {
	if err := m.Orm.Where("material_id = ?", id).Unscoped().Delete(&models.MaterialSql{}).Error; err != nil {
		m.Log.Errorf("delete material failed,db err:%s", err)
		return err
	}
	return nil
}

func (m *MaterialSql) Update(data *models.Material) error {
	return m.Orm.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("material_id = ?", data.MaterialID).
			Updates(&models.MaterialSql{
				Material: *data,
			}).Error; err != nil {
			m.Log.Errorf("updates failed,db err:%s", err)
			return err
		}
		return nil
	})
}

func (m *MaterialSql) Query(datalist *[]models.MaterialSql) (*[]models.MaterialSql, error) {
	err := m.Orm.Find(datalist).Limit(-1).Offset(-1).Error
	if err != nil {
		m.Log.Errorf("material query failed,db err:%s", err)
		return nil, err
	}
	return datalist, nil
}
