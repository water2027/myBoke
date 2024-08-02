package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Jmjx      string `json:"jmjx"`
    Body      string `json:"body"`
	Tag       string `json:"tag"`
	CreatedAt string `json:"created_at"`
}

type UserInfo struct {
	Time     string `json:"time"`
	Week     string `json:"week"`
	Ip       string `json:"ip"`
	Location string `json:"location"`
	Browser  string `json:"browser"`
	System   string `json:"system"`
}

func getPosts(c *gin.Context) {
    rows, err := db.Query("SELECT * FROM posts")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "查询失败", "error": err.Error()})
        log.Printf("查询失败: %v", err)
        return
    }
    defer rows.Close()

    c.Header("Content-Type", "application/json")
    c.Status(http.StatusOK)

    var posts []Post
    for rows.Next() {
        var post Post
        if err := rows.Scan(&post.ID, &post.Title, &post.Jmjx, &post.Body, &post.Tag, &post.CreatedAt); err != nil {
            log.Printf("行扫描错误: %v", err)
            continue
        }
        posts = append(posts, post)
    }
    c.JSON(http.StatusOK, gin.H{"message": "success", "data": posts})
}

func getReply(c *gin.Context) {
	title := c.Query("title")
	rows, err := db.Query("SELECT * FROM reply WHERE title = ?", title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取reply失败", "error": err.Error()})
	}
	defer rows.Close()
    var reply []Reply
    for rows.Next() {
        var r Reply
        if err := rows.Scan(&r.ID, &r.Title, &r.Name, &r.Reply, &r.Time); err != nil {
            log.Printf("行扫描错误: %v", err)
            continue
        }
        reply = append(reply, r)
    }
	c.JSON(http.StatusOK, gin.H{"message": "获取reply成功", "data": reply})
}

func getUserInfo(c *gin.Context) {
	var userInfo UserInfo
	if err := c.BindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误", "error": err.Error()})
		return
	}
	_, err := db.Exec("INSERT INTO user_info (time, week, ip, location, browser, deviceSystem) VALUES (?, ?, ?, ?, ?, ?)", userInfo.Time, userInfo.Week, userInfo.Ip, userInfo.Location, userInfo.Browser, userInfo.System)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "插入失败", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})

}
