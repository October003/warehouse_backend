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
		if err := a.Orm.Create(&models.InBoundPersons{Name: name}).Error; err != nil {
			return err
		}
	case 2:
		if err := a.Orm.Create(&models.OutBoundPersons{Name: name}).Error; err != nil {
			return err
		}
	case 3:
		if err := a.Orm.Create(&models.Units{Name: name}).Error; err != nil {
			return err
		}
	case 4:
		if err := a.Orm.Create(&models.StrongLocation{Name: name}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (a *Admin) GetAll() (*models.AdminConfig, error) {
	var (
		inboundPersons   []models.InBoundPersons
		outboundPersons  []models.OutBoundPersons
		units            []models.Units
		storageLocations []models.StrongLocation
		ItemIDs          []models.ItemIDs
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
	config.StrongLocation = storageLocations
	if err := a.Orm.Find(&ItemIDs).Error; err != nil {
		return nil, err
	}
	config.ItemIDs = ItemIDs
	return &config, nil
}

func (a *Admin) Delete(name string, num int) error {
	switch num {
	case 1:
		if err := a.Orm.Where("Name = ?", name).Unscoped().Delete(&models.InBoundPersons{}).Error; err != nil {
			return err
		}
	case 2:
		if err := a.Orm.Where("Name = ?", name).Unscoped().Delete(&models.OutBoundPersons{}).Error; err != nil {
			return err
		}
	case 3:
		if err := a.Orm.Where("Name = ?", name).Unscoped().Delete(&models.Units{}).Error; err != nil {
			return err
		}
	case 4:
		if err := a.Orm.Where("Name = ?", name).Unscoped().Delete(&models.StrongLocation{}).Error; err != nil {
			return err
		}
	}
	return nil
}
