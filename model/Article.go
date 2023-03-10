package model

import (
	"ginVueBlog/utils/errmsg"
	"gorm.io/gorm"
)

// 文章
type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `grom:"type:varchar(100);not null" json:"title"`
	Cid     int    `grom:"type:int;not null" json:"cid"`
	Desc    string `grom:"type:varchar(200)" json:"desc"`
	Content string `grom:"type:longtext" json:"content"`
	Img     string `grom:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// 查询文章列表
func GetArt(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	if title == "" {
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Preload("Category").Find(&articleList).Error
		// 单独计数
		db.Model(&articleList).Count(&total)
		if err != nil {
			return nil, errmsg.ERROR, 0
		}
		return articleList, errmsg.SUCCSE, total
	}
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Preload("Category").Where("title LIKE ?", title+"%").Find(&articleList).Error
	// 单独计数
	db.Model(&articleList).Where("title LIKE ?", title+"%").Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total
}

// 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&art).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ? ", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id=?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCSE
}

// 查询分类下的所有文章
func GetcateArt(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64
	db.Preload("Category").Where("cid=?", id).First(&cateArtList).Count(&total)
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", id).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCSE, total
}
