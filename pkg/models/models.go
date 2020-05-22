package models

import (
	"gitlab.mcsolutions.ru/find-psy/common/dbmodels"
	"gitlab.mcsolutions.ru/lib/common/paging"
)

const NEWS = "news"

type NewsTable OneNewsFull

func (NewsTable) TableName() string {
	return NEWS
}

type (
	Title struct {
		Title string `json:"title" gorm:"size:200;not null" valid:"required,stringlength(5|200)"`
	}
	Text struct {
		Text string `json:"text" gorm:"type:text" valid:""`
	}
)

type (
	OneNews struct {
		dbmodels.Date
		Title
		Text
		dbmodels.Region
		dbmodels.Tags
	}
	OneNewsFull struct {
		dbmodels.Id
		dbmodels.Lang
		OneNews
		dbmodels.Login
		dbmodels.CreateDate
		dbmodels.UpdateDate
	}
)

type (
	NewsOut struct {
		News *[]OneNews `json:"news"`
		paging.PagingOut
	}
	NewsFullOut struct {
		News *[]OneNewsFull `json:"news"`
		paging.PagingOut
	}
	IdOut struct {
		dbmodels.Id
		dbmodels.CreateDate
		dbmodels.UpdateDate
	}
)
