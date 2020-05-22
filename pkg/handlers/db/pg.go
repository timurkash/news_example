package db

import (
	"bytes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.mcsolutions.ru/find-psy/back/admins/news/pkg/models"
	"gitlab.mcsolutions.ru/find-psy/common/dbmodels"
	"gitlab.mcsolutions.ru/lib/common/config"
	"gitlab.mcsolutions.ru/lib/common/gorm/pg"
	"gitlab.mcsolutions.ru/lib/common/json"
	"gitlab.mcsolutions.ru/lib/common/paging"
	"time"
)

var (
	POSTGRES_ARGS = config.GetEnv("POSTGRES_ARGS", "postgresql://doadmin@db-postgresql-ams3-60546-do-user-6494983-0.db.ondigitalocean.com:25060/defaultdb?sslmode=require")
)

type News struct {
	db *gorm.DB
}

func (t *News) checkDb() error {
	if t.db == nil {
		postgres := pg.Postgres{
			Name:   models.NEWS,
			Schema: models.NEWS,
			Args:   POSTGRES_ARGS,
			Model:  &models.NewsTable{},
		}
		var err error
		t.db, err = postgres.GetDb()
		if err != nil {
			return err
		}
	}
	//return t.db.Exec("SET search_path TO " + NEWS).Error
	return nil
}

func (t *News) getModel(lang, region string) *gorm.DB {
	where := &models.NewsTable{
		Lang: dbmodels.Lang{lang},
	}
	regionWhere := &models.NewsTable{
		OneNews: models.OneNews{
			Region: dbmodels.Region{region},
		},
	}
	model := t.db.Where(where).Order("date desc").Order("id desc")
	if region != "all" {
		model = model.Where(regionWhere)
	}
	return model
}

func (t *News) Get(lang, region string, page *[]int64) (*[]models.OneNews, *paging.Paging, error) {
	if err := t.checkDb(); err != nil {
		return nil, nil, err
	}
	items, paging, err := paging.GetItems(t.getModel(lang, region), page, &models.OneNews{})
	if err != nil {
		return nil, nil, err
	}
	news := []models.OneNews{}
	oneNews := &models.OneNews{}
	for _, item := range *items {
		if err := json.Decode(bytes.NewReader(item), oneNews); err != nil {
			return nil, nil, err
		}
		news = append(news, *oneNews)
	}
	return &news, paging, nil
}

func (t *News) GetFull(lang, region string, page *[]int64) (*[]models.OneNewsFull, *paging.Paging, error) {
	if err := t.checkDb(); err != nil {
		return nil, nil, err
	}
	items, paging, err := paging.GetItems(t.getModel(lang, region), page, &models.OneNewsFull{})
	if err != nil {
		return nil, nil, err
	}
	news := []models.OneNewsFull{}
	for _, item := range *items {
		oneNewsFull := &models.OneNewsFull{}
		if err := json.Decode(bytes.NewReader(item), oneNewsFull); err != nil {
			return nil, nil, err
		}
		news = append(news, *oneNewsFull)
	}
	return &news, paging, nil
}

func (t *News) Insert(lang, login string, oneNews *models.OneNews) (*models.IdOut, error) {
	if err := t.checkDb(); err != nil {
		return nil, err
	}
	oneNewsFull := &models.OneNewsFull{
		Lang:    dbmodels.Lang{lang},
		OneNews: *oneNews,
		Login:   dbmodels.Login{login},
	}
	err := t.db.Create(oneNewsFull).Error
	if err != nil {
		return nil, err
	}
	return &models.IdOut{
		Id:         oneNewsFull.Id,
		CreateDate: oneNewsFull.CreateDate,
	}, nil
}

func (News) getIdWhere(id int64) *models.NewsTable {
	return &models.NewsTable{
		Id: dbmodels.Id{id},
	}
}

func (t *News) Update(id int64, login string, oneNews *models.OneNews) (*models.IdOut, error) {
	if err := t.checkDb(); err != nil {
		return nil, err
	}
	now := time.Now()
	oneNewsFull := &models.OneNewsFull{
		OneNews:    *oneNews,
		Login:      dbmodels.Login{login},
		UpdateDate: dbmodels.UpdateDate{&now},
	}
	err := t.db.Where(t.getIdWhere(id)).Update(oneNewsFull).Error
	if err != nil {
		return nil, err
	}
	return &models.IdOut{
		UpdateDate: oneNewsFull.UpdateDate,
	}, nil
}

func (t *News) Delete(id int64) error {
	if err := t.checkDb(); err != nil {
		return err
	}
	return t.db.Delete(t.getIdWhere(id)).Error
}
