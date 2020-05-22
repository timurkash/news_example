package pg

import "github.com/jinzhu/gorm"

type Postgres struct {
	Name   string
	Schema string
	Args   string
	Model  interface{}
}

func (p *Postgres) GetDb() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", p.Args)
	if err != nil {
		return nil, err
	}
	if p.Schema != "" && p.Schema != "public" {
		if err := db.Exec("create schema if not exists " + p.Schema).Error; err != nil {
			return nil, err
		}
		if err := db.Exec("SET search_path TO " + p.Schema).Error; err != nil {
			return nil, err
		}
	}
	db = db.Table(p.Name)
	if db.HasTable(p.Name) {
		if err = db.AutoMigrate(p.Model).Error; err != nil {
			return nil, err
		}
	} else {
		if err = db.CreateTable(p.Model).Error; err != nil {
			return nil, err
		}
	}
	return db, nil
}
