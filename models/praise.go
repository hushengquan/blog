package models

import (
	"github.com/jinzhu/gorm"
	"gostudy/newstudy/blog/syserror"
)

type PraiseLog struct {
	Model
	UserID int
	Key string
	Table string
	Flag bool
}

type TempPraise struct {
	Praise int
}

func UpdatePraise(table, key string, userid int) (pcnt int, err error) {
	d := db.Begin()
	defer func() {
		if err != nil {
			d.Rollback()
		} else {
			d.Commit()
		}
	}()

	var p PraiseLog
	err = d.Model(&PraiseLog{}).Where("`key` = ? and `table` =? and user_id = ?", key, table, userid).Take(&p).Error
	if err == gorm.ErrRecordNotFound {
		p = PraiseLog{
			UserID: userid,
			Key:    key,
			Table:  table,
			Flag:   false,
		}
	} else if err != nil {
		return 0, err
	}
	if p.Flag {
		return 0, syserror.HasPraiseError{}
	}
	p.Flag = true

	if err = d.Save(&p).Error; err != nil {
		return 0, err
	}

	var ppp TempPraise
	err = d.Table(table).Where("key = ?", key).Select("praise").Scan(&ppp).Error
	if err != nil {
		return 0, err
	}
	pcnt = ppp.Praise + 1
	err = d.Table(table).Where("key = ? ", key).Update("praise", pcnt).Error
	if err != nil {
		return 0, err
	}

	return pcnt, nil
}