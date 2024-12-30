package model

type Goods struct {
	ID    uint   `gorm:"primary_key;auto_increment" json:"id"`
	Title string `gorm:"type:varchar(255);not null" json:"title"`
	Mid   uint   `gorm:"column:mid;not null" json:"mid"` // Foreign key to Merchant
}

type Merchant struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}

type Result struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Name   string `json:"name"`
	Mid    uint   `json:"mid"`
	Remark string `json:"remark"`
}

// GetCooperation 查看合作模式
func GetGoods(page, size int) ([]Result, int64, error) {
	var results []Result
	var total int64

	// Get total count of goods
	err := db.Model(&Goods{}).
		Joins("left join merchant on goods.mid = merchant.id").
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Fetch the goods with pagination
	err = db.Model(&Goods{}).
		Select("goods.id AS id, goods.title AS title, goods.mid AS mid, merchant.name AS name").
		Joins("left join merchant on goods.mid = merchant.id").
		Offset((page - 1) * size). // Skip items for pagination
		Limit(size).               // Limit the number of items
		Find(&results).Error
	if err != nil {
		return nil, 0, err
	}

	// Set the Remark field to merchant's name for each result
	for k, _ := range results {
		results[k].Remark = "asdas"
	}

	return results, total, nil
}

// CREATE TABLE `goods` (
// 	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
// 	`title` varchar(255) DEFAULT '' COMMENT '商品名称',
// 	`mid` varchar(255) DEFAULT '' COMMENT '商家id',
// 	`hidden` tinyint(4) DEFAULT '0' COMMENT '是否隐藏 0:否 1:是',
// 	`created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
// 	`updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
// 	PRIMARY KEY (`id`) USING BTREE
//   ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='商品';

//   CREATE TABLE `merchant` (
// 	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
// 	`name` varchar(255) DEFAULT '' COMMENT '商家名称',
// 	`hidden` tinyint(4) DEFAULT '0' COMMENT '是否隐藏 0:否 1:是',
// 	`created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
// 	`updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
// 	PRIMARY KEY (`id`) USING BTREE
//   ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='商家';
