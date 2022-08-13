package controller

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// IndexHandler 访问主页控制器
func IndexHandler(c *gin.Context) {
	// 从service取数据
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 加载分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//var ret map[string]interface{} = make(map[string]interface{})
	//ret["articleRecordList"] = articleRecordList
	//ret["categoryList"] = categoryList
	//c.HTML(http.StatusOK, "views/index.html", ret)

	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}

// CategoryList 点击分类云进行分类
func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	// 转int
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 根据分类id, 获取文章列表
	articleRecordList, err := service.GetArticleRecordListById(categoryId, 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 再次加载所有分类
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}

func ArticleDetail(c *gin.Context) {
	articleIdStr := c.Query("article_id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 获取文章详情
	articleDetail, err := service.GetArticleDetail(articleId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/detail.html", gin.H{
		"detail": articleDetail,
	})
}
