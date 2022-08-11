package db

import (
	"blog/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// InsertArticle 插入文章
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	// 验证
	if article == nil {
		return
	}
	sqlStr := `INSERT INTO article(content, summary, title, username, category_id, view_count, comment_count) 
				VALUES(?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlStr, article.Content, article.Summary, article.Title, article.Username, article.ArticleInfo.CategoryId, article.ViewCount, article.CommentCount)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

// GetArticleList 获取文章列表
func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageSize <= 0 || pageNum < 0 {
		return
	}
	sqlStr := `SELECT 
				id, summary, title, username, category_id, view_count, comment_count
			FROM
				article
			WHERE
				status = 1
			ORDER BY create_time DESC
			LIMIT ?,?`
	err = DB.Select(&articleList, sqlStr, pageNum, pageSize)
	return
}

// GetArticleDetail 根据文章id查询文章详情
func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId <= 0 {
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlstr := `select 
							id, summary, title, view_count, content,
							 create_time, comment_count, username, category_id
						from 
							article 
						where 
							id = ?
						and
							status = 1
						`
	err = DB.Get(articleDetail, sqlstr, articleId)
	return
}

// GetArticleListByCategoryId 根据分类id查询
func GetArticleListByCategoryId(categoryId int64, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if categoryId <= 0 || pageSize <= 0 || pageNum < 0 {
		return
	}
	sqlStr := `SELECT 
				id, summary, title, view_count, create_time, comment_count, username, category_id
			FROM
				article
			WHERE 
				category_id = ?
			AND 
			    status = 1
			ORDER BY create_time DESC
			limit ?,?`
	fmt.Println(sqlStr)
	err = DB.Select(&articleList, sqlStr, categoryId, pageNum, pageSize)
	return
}
