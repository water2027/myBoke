package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Reply struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
	Reply string `json:"reply"`
	Time  string `json:"time"`
}

func addReply(c *gin.Context) {
	session,error := Store.Get(c.Request,"userInfo")
	if error != nil {
		fmt.Println(error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "please login first"})
		return
	}
	title := c.PostForm("title")
	name := session.Values["username"]
	reply := c.PostForm("reply")
	fmt.Println("title,name,reply",title, name, reply)
	_, err := db.Exec("INSERT INTO reply (title,name,reply) VALUES (?, ?, ?)", title, name, reply)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": "add reply failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "add reply success",
	})
}

func delReply(c *gin.Context) {
	sessions, _ := Store.Get(c.Request, "userInfo")
	id := c.Query("id")
	rows, err := db.Query("SELECT name FROM reply WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "delete reply failed"})
		return
	}
	defer rows.Close()

	if rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "delete reply failed"})
			return
		}
		if name != sessions.Values["username"] {
			c.JSON(http.StatusForbidden, gin.H{"message": "not authorized to delete this reply"})
			return
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "reply not found"})
		return
	}
	if _, err := db.Exec("DELETE FROM reply WHERE id = ?", id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "delete reply failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete reply success"})
}

func follow(c *gin.Context) {
	session, _ := Store.Get(c.Request, "userInfo")
	id := session.Values["id"]
	if id != 0 {
		row,error := db.Query("SELECT follow FROM users WHERE id = ?",id)
		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "失败，数据库错误"})
			return
		}
		defer row.Close()
		var follow bool
		if row.Next() {
			if err:=row.Scan(&follow);err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "失败，数据库错误"})
				return
			}
			follow = !follow
		}
        sqlStatement := `UPDATE users SET follow = ? WHERE id = ?`
		_, err := db.Exec(sqlStatement, follow, id)
        session.Values["follow"] = follow
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "失败，数据库错误"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success","code":true})
	}else{
        c.JSON(http.StatusForbidden,gin.H{"message":"请先登录！","code":false})
    }
}
