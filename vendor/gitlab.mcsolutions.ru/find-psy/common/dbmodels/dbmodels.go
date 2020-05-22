package dbmodels

import "time"

type (
	Id struct {
		Id int64 `json:"id,omitempty" gorm:"size:32;primary_key;not null;auto_increment" valid:"numeric,required"`
	}
	Lang struct {
		Lang string `json:"lang" gorm:"size:2;not null;index" valid:"required,stringlength(2|2)"`
	}
	LangPrimary struct {
		Lang string `json:"lang" gorm:"size:2;primary_key" valid:"required,stringlength(2|2)"`
	}
	Key struct {
		Key string `json:"key" gorm:"size:32;not null;index" valid:"required,stringlength(2|32)"`
	}
	KeyPrimary struct {
		Key string `json:"key" gorm:"size:32;primary_key" valid:"required,stringlength(2|32)"`
	}
	Region struct {
		Region string `json:"region" gorm:"size:16;index;not null" valid:"required,stringlength(5|16)"`
	}
	Tags struct {
		Tags string `json:"tags" gorm:"type:text"`
	}
	Login struct {
		Login string `json:"login" gorm:"size:32;index;not null" valid:"alphanumeric,required,stringlength(5|32)"`
	}
	LoginPrimary struct {
		Login string `json:"login" gorm:"size:32;primary_key" valid:"alphanumeric,required,stringlength(5|32)"`
	}
	Password struct {
		Password string `json:"password" gorm:"size:32;not null" valid:"alphanumeric,required,stringlength(5|32)"`
	}
	Name struct {
		Name string `json:"name" gorm:"size:50;not null" valid:"alphanumeric,required,stringlength(5|50)"`
	}
	Email struct {
		Email string `json:"email" gorm:"size:50;unique;not null" valid:"email,required,stringlength(5|50)"`
	}
	Msisdn struct {
		Msisdn string `json:"msisdn" gorm:"unique;size:15"`
	}
	IsActive struct {
		IsActive bool `json:"isActive" gorm:"column:isactive;not null;default:true" valid:"required"`
	}
	Date struct {
		Date string `json:"date" gorm:"size:10;not null" valid:"alphanumeric,required,stringlength(10|10)"`
	}
	CreateDate struct {
		CreateDate *time.Time `json:"createDate,omitempty" gorm:"column:createdate;not null;default:now()"`
	}
	UpdateDate struct {
		UpdateDate *time.Time `json:"updateDate,omitempty" gorm:"column:updatedate"`
	}
	Ord struct {
		Ord float32 `json:"ord" gorm:"" valid:""`
	}
)
