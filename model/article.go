package model

import "time"

// ArticleInfo 文章结构体
type ArticleInfo struct {
	Id         int64 `db:"id"`
	CategoryId int64 `db:"category_id"`
	// 文章摘要
	Summary   string `db:"summary"`
	Title     string `db:"title"`
	ViewCount uint32 `db:"view_count"`
	// 时间
	CreateTime   time.Time `db:"create_time"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
}

// ArticleDetail 文章详情页
type ArticleDetail struct {
	ArticleInfo
	// 文章内容
	Content string `db:"content"`
	Category
}

// ArticleRecord 用于文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}
