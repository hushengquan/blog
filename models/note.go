package models

import (
	"fmt"
)

type Note struct {
	Model
	Key string `gorm:"unique_index;not null;"`
	UserID int
	User User
	Title string
	Summary string `gorm:"type:text"`
	Content string `gorm:"type:text"`
	//Source  string `gorm:"type:text" json:"source"`
	//Editor  string `gorm:"varchar(40)"`
	Visit int `gorm:"default:0"`
	Praise int `gorm:"default:0"`
}

func QueryNoteByKeyAndUserId(key string, userid int) (note Note, err error) {
	return note, db.Where("key = ? and user_id = ?", key, userid).Take(&note).Error
}

func QueryNoteByKey(key string) (note Note, err error) {
	return note, db.Where("key = ?", key).Take(&note).Error
}

func SaveNote(note *Note) error {
	return db.Save(note).Error
}

func QueryNotesCount(title string) (count int, err error) {
	return count, db.Where("title like ?", fmt.Sprintf("%%%s%%", title)).Model(&Note{}).Count(&count).Error
}

func QueryNotesByPage(page, limit int, title string) (note []*Note, err error) {
	return note, db.Where("title like ?", fmt.Sprintf("%%%s%%", title)).Order("updated_at desc").Offset((page - 1) * limit).Limit(limit).Find(&note).Error
}

func DeleteNoteByUserIdAndKey(key string, userid int) error {
	return db.Delete(&Note{}, "key = ? and user_id = ?", key, userid).Error
}
