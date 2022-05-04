package database

import (
	"engine_app/database/model"
	"gorm.io/gorm"
)

type IMigrator interface {
	MigrateUser() error
	MigrateProjectDoc() error
	MigrateProjectDescription() error
	//MigrateConfiguration() error
	//MigrateLanguage() error
	//MigrateBuilder() error
	MigrateData() error
	MigrateAll() error
}

type Migrator struct {
	db *gorm.DB
}

func (m *Migrator) MigrateAll() error {

	err := m.MigrateUser()
	if err != nil {
		return err
	}
	err = m.MigrateProjectDoc()
	if err != nil {
		return err
	}

	err = m.MigrateProjectDescription()
	if err != nil {
		return err
	}

	err = m.MigrateData()
	if err != nil {
		return err
	}
	return nil
}

func MakeMigrator(Db *gorm.DB) IMigrator {
	return &Migrator{db: Db}
}

func (m *Migrator) MigrateUser() error {
	err := m.db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	return nil
}

func (m *Migrator) MigrateProjectDoc() error {
	err := m.db.AutoMigrate(&model.ProjectConfig{})
	if err != nil {
		return err
	}
	return nil
}

func (m *Migrator) MigrateProjectDescription() error {
	err := m.db.AutoMigrate(&model.Project{})
	if err != nil {
		return err
	}
	return nil
}

func (m *Migrator) MigrateData() error {
	err := m.db.AutoMigrate(&model.Data{})
	if err != nil {
		return err
	}
	return nil
}

//func (m *Migrator) MigrateConfiguration() error {
//	err := m.db.AutoMigrate(&model.Configuration{})
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (m *Migrator) MigrateLanguage() error {
//	err := m.db.AutoMigrate(&model.Language{})
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (m *Migrator) MigrateBuilder() error {
//	err := m.db.AutoMigrate(&model.Builder{})
//	if err != nil {
//		return err
//	}
//	return nil
//}
