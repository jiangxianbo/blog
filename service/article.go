package service

import (
	"blog/dao/db"
	"blog/model"
)

// GetArticleRecordList 获取文章和对应的分类
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章列表
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}

	// 2.获取文章分类
	categoryIds := getCategoryIs(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}

	// 聚合
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		// 文章分类id
		categoryId := article.CategoryId
		// 遍历分类
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

// getCategoryIs 根据多个文章的id，获取多个分类id的集合
func getCategoryIs(articleInfoList []*model.ArticleInfo) (ids []int64) {
LABEL:
	// 遍历文章
	for _, article := range articleInfoList {
		categoryId := article.CategoryId
		// 去重，防止重复
		for _, id := range ids {
			if id == categoryId {
				continue LABEL
			}
		}
		ids = append(ids, categoryId)
	}
	return
}

// GetArticleRecordListById 根据id获取该类文章和对应分类信息
func GetArticleRecordListById(categoryId int64, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	articleInfoList, err := db.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}

	// 2.获取文章分类
	categoryIds := getCategoryIs(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}

	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		// 文章分类id
		categoryId := article.CategoryId
		// 遍历分类
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

// GetArticleDetail 获取文章详情
func GetArticleDetail(id int64) (articleDetail *model.ArticleDetail, err error) {
	articleDetail, err = db.GetArticleDetail(id)
	if err != nil {
		return
	}
	return
}
