package db

import (
	"blog/model"
	"testing"
	"time"
)

func init() {
	// parseTime=true 将mysql中时间类型，自动解析为go结构体中的时间类型
	// 不加报错
	dns := "root:jiang123@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertArticle(t *testing.T) {
	// 构建对象
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.Content = "abc"
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Title = "test001"
	article.ArticleInfo.Username = "jxb"
	article.ArticleInfo.Summary = "a"
	article.ArticleInfo.ViewCount = 1
	articleId, err := InsertArticle(article)
	if err != nil {
		panic(err)
	}
	t.Logf("articleId:%d\n", articleId)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		panic(err)
	}
	for _, v := range articleList {
		t.Logf("articleList:%#v\n", v)
	}
}

func TestGetArticleDetail(t *testing.T) {
	articleDetail, err := GetArticleDetail(1)
	if err != nil {
		panic(err)
	}
	t.Logf("%#v\n", articleDetail)
}

func TestGetArticleListByCategoryId(t *testing.T) {
	articleList, err := GetArticleListByCategoryId(1, 1, 10)
	if err != nil {
		panic(err)
	}
	for _, v := range articleList {
		t.Logf("articleList:%v\n", v)
	}
}
