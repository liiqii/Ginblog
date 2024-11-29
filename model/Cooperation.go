package model

import (
	"time"

	"github.com/wejectchen/ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Cooperation struct {
	ID              uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name            string    `gorm:"type:varchar(255);not null" json:"name"`
	Phone           string    `gorm:"type:varchar(255);not null" json:"phone"`
	Content         string    `gorm:"type:varchar(255);not null" json:"content"`
	Platform        int       `gorm:"type:int;not null" json:"platform"`
	PlatformContent string    `gorm:"type:varchar(255)" json:"platform_content"`
	Status          int       `gorm:"type:int;not null" json:"status"`
	Remark          string    `gorm:"type:varchar(255);not null" json:"remark"`
	CreatedTime     time.Time `json:"created_time"` // 创建时间
	UpdatedTime     time.Time `json:"updated_time"` // 更新时间
}

// AddCooperation 添加合作模式
func CreateCooperation(data *Cooperation) (int, error) {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR, err // 500
	}
	return errmsg.SUCCESS, nil
}

// GetCooperation 查看合作模式
func GetCooperation(pageSize int, pageNum int) ([]Cooperation, int64) {
	var cooperation []Cooperation
	var total int64
	err = db.Find(&cooperation).Where("hidden = ?", 0).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&cooperation).Where("hidden = ?", 0).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cooperation, total
}

// EditCooperation // EditCooperation 编辑合作模式
func EditCooperation(id uint, data *Cooperation) (int, error) {
	var cooperation Cooperation
	var maps = make(map[string]interface{})
	maps["status"] = data.Status
	maps["remark"] = data.Remark

	err = db.Model(&cooperation).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR, err
	}
	return errmsg.SUCCESS, nil
}

func DeleteCooperation(id uint) (int, error) {
	var cooperation Cooperation
	var maps = make(map[string]interface{})
	maps["hidden"] = 1

	err = db.Model(&cooperation).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR, err
	}
	return errmsg.SUCCESS, nil
}
