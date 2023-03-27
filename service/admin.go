package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"warehouse/models"
)

type Admin struct {
	service.Service
}

func (a *Admin) Add(name string, num int) error {
	if name == "" {
		return nil
	}
	switch num {
	case 1:
		if err := a.Orm.Create(&models.InBoundPerson{Name: name}).Error; err != nil {
			return err
		}
	case 2:
		if err := a.Orm.Create(&models.OutBoundPerson{Name: name}).Error; err != nil {
			return err
		}
	case 3:
		if err := a.Orm.Create(&models.Unit{Name: name}).Error; err != nil {
			return err
		}
	case 4:
		if err := a.Orm.Create(&models.StorageLocation{Name: name}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (a *Admin) GetAll() (*models.AdminConfig, error) {
	var (
		inboundPersons   []models.InBoundPerson
		outboundPersons  []models.OutBoundPerson
		units            []models.Unit
		storageLocations []models.StorageLocation
		materialIds      []models.MaterialID
		config           models.AdminConfig
	)
	if err := a.Orm.Find(&inboundPersons).Error; err != nil {
		return nil, err
	}
	config.InBoundPersons = inboundPersons
	if err := a.Orm.Find(&outboundPersons).Error; err != nil {
		return nil, err
	}
	config.OutBoundPersons = outboundPersons
	if err := a.Orm.Find(&units).Error; err != nil {
		return nil, err
	}
	config.Units = units
	if err := a.Orm.Find(&storageLocations).Error; err != nil {
		return nil, err
	}
	config.StorageLocation = storageLocations
	if err := a.Orm.Find(&materialIds).Error; err != nil {
		return nil, err
	}
	config.MaterialIds = materialIds
	return &config, nil
}

func (a *Admin) Delete(name string, num int) error {
	switch num {
	case 1:
		if err := a.Orm.Where("Name = ?", name).Unscoped().Delete(&models.InBoundPerson{}).Error; err != nil {
			return err
		}
	case 2:
		if err := a.Orm.Where("Name = ?", name).Unscoped().Delete(&models.OutBoundPerson{}).Error; err != nil {
			return err
		}
	case 3:
		if err := a.Orm.Where("Name = ?", name).Unscoped().Delete(&models.Unit{}).Error; err != nil {
			return err
		}
	case 4:
		if err := a.Orm.Where("Name = ?", name).Unscoped().Delete(&models.StorageLocation{}).Error; err != nil {
			return err
		}
	}
	return nil
}
