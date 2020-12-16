package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

// gorm的Callback
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error  {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// gorm的Callback
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error  {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}


// 查询文章标签
func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag)  {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

// 查询文章标签的总数
func GetTagTotal(maps interface {}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

// 根据标签名判断标签是否存在
func ExistTagByName(name string) bool  {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

// 根据标签id判断标签是否存在
func ExistTagByID(id int) bool  {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

// 添加新标签
func AddTag(name string, state int, createdBy string) bool  {
	//if err := db.Create(&Tag {
	//	Name : name,
	//	State : state,
	//	CreatedBy : createdBy,
	//}).Error;err != nil {
	//	fmt.Printf("failed to add:%v",err)
	//}
	db.Create(&Tag {
		Name : name,
		State : state,
		CreatedBy : createdBy,
	})

	return true
}

// 删除标签
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

// 更新标签
func EditTag(id int, data interface{}) bool  {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}