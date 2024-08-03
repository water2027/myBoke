package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var Store = sessions.NewCookieStore([]byte("a secret"))
var adminToken string

var db *sql.DB

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	adminToken = os.Getenv("ADMIN_TOKEN")

	r := gin.Default()
	r.LoadHTMLGlob("dist/*.html") 
	r.Static("/asset", "dist/asset") 
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", user, pass, host, dbName)
	db, err = sql.Open("mysql", dsn)
	if db == nil {
		fmt.Println("db is nil")
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()	
	if err = db.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://game.watering.top") // 允许的源
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	// public
	r.GET("/", homePage)
	r.GET("/api/getPosts", getPosts)
	r.GET("/api/getReply", getReply)
	r.POST("/api/getUserInfo", getUserInfo)

	// user
	r.POST("/api/sendCode", sendCode)
	r.POST("/api/register", register)
	r.POST("/api/login", login)
	r.POST("/api/addReply", addReply)
	r.GET("/api/delReply", delReply)
	r.GET("/api/follow",follow)

	//admin
	r.POST("/api/writePost",isAdmin(), writePost)
	r.POST("/api/editPost",isAdmin(), editPost)
	r.GET("/api/seeUserInfos",isAdmin(), seeUserInfos)

	// websocket
	r.GET("/api/chatroom", websocketChat)

	r.Run(":3000")

	fmt.Println("Server is running on port 3000")
}
