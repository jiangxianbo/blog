package db

import (
	"blog/model"
	"github.com/jmoiron/sqlx"
)

// InsertCategory 添加分类
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	sqlStr := "INSERT INTO `category`(`category_name`,`category_no`) VALUES(?,?)"
	result, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryId, err = result.LastInsertId()
	return
}

// GetCategoryById 获取单个文章分类
func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlStr := "SELECT `id`,`category_name`,`category_no` FROM `category` WHERE `id` = ?"
	err = DB.Get(category, sqlStr, id)
	return
}

// GetCategoryList 查询多个
func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {
	// 构建sql
	sqlStr, args, err := sqlx.In("SELECT `id`,`category_name`,`category_no` FROM `category` WHERE `id` in (?)", categoryIds)
	if err != nil {
		return
	}
	// 查询
	err = DB.Select(&categoryList, sqlStr, args...)
	return
}

// 获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	// 构建sql
	sqlStr := "SELECT `id`,`category_name`,`category_no` FROM `category` ORDER BY category_no ASC"
	// 查询
	err = DB.Select(&categoryList, sqlStr)
	return
}
