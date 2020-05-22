package paging

import (
	"errors"
	"github.com/jinzhu/gorm"
	"gitlab.mcsolutions.ru/lib/common/json"
)

type (
	Paging struct {
		From       int64 `json:"from"`
		To         int64 `json:"to"`
		Count      int64 `json:"count"`
		HasNext    bool  `json:"hasNext"`
		TotalCount int64 `json:"totalCount"`
	}

	PagingOut struct {
		Paging *Paging `json:"paging,omitempty"`
	}
)

func GetItems(model *gorm.DB, page *[]int64, item interface{}) (*[][]byte, *Paging, error) {
	if len(*page) != 2 {
		return nil, nil, errors.New("wrong page")
	}
	var totalCount int64
	model.Count(&totalCount)
	from := (*page)[0]
	if from > 0 {
		model = model.Offset(from)
	}
	limit := (*page)[1]
	if limit > 0 {
		model = model.Limit(limit)
	}
	rows, err := model.Rows()
	defer rows.Close()
	if err != nil {
		return nil, nil, err
	}
	items := [][]byte{}
	var count int64
	for rows.Next() {
		if err := model.ScanRows(rows, item); err != nil {
			return nil, nil, err
		}
		bytes, err := json.EncodeByte(item)
		if err != nil {
			return nil, nil, err
		}
		items = append(items, *bytes)
		count++
	}
	paging := Paging{
		From:       from,
		To:         from + count,
		Count:      count,
		HasNext:    from+count < totalCount,
		TotalCount: totalCount,
	}
	return &items, &paging, nil
}
