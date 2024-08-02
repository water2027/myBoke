package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func isAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		token = md5V(token)
		if token != "token" {
			c.JSON(403, gin.H{
				"message": "you are not admin",
			})
			c.Abort()
		}
		c.Next()
	}
}

func writePost(c *gin.Context) {
	//验证身份
	title := c.PostForm("title")
	jmjx := c.PostForm("jmjx")
	body := c.PostForm("body")
	tag := c.PostForm("tag")
	if tag == "" {
		tag = "默认"
	}
	_, err := db.Exec("INSERT INTO posts (title, jmjx, body, tag) VALUES (?, ?, ?, ?)", title, jmjx, body, tag)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "write post failed",
		})
	}
	rows, err := db.Query("SELECT email FROM users WHERE follow = true")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			log.Println(err)
		}
		sendEmail(email, "新文章", "新文章"+title+"发布了")
	}
	c.JSON(200, gin.H{
		"message": "write post success",
	})
}

func editPost(c *gin.Context) {
	id := c.PostForm("id")
	title := c.PostForm("title")
	jmjx := c.PostForm("jmjx")
	body := c.PostForm("body")
	tag := c.PostForm("tag")
	if tag == "" {
		tag = "默认"
	}
	_, err := db.Exec("UPDATE posts SET title = ?, jmjx = ?, body = ?, tag = ? WHERE id = ?", title, jmjx, body, tag, id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "edit post failed",
		})
	}
	c.JSON(200, gin.H{
		"message": "edit post success",
	})
}

func seeUserInfos(c *gin.Context) {
	//验证身份
	rows, err := db.Query("SELECT * FROM userinfo")
	if err != nil {
		c.JSON(500, gin.H{
			"message": "get user info failed",
		})
	}
	defer rows.Close()
	c.JSON(200, gin.H{
		"message": "get user info success",
		"data":    rows,
	})
}
