package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Follow   bool   `json:"follow"`
}

func md5V(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return md5str1
}

func getRandomCode(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func sendCode(c *gin.Context) {
	// get email or send email
	session, _ := Store.Get(c.Request, "userInfo")
	email := c.PostForm("email")
	rowA, error := db.Query("SELECT username FROM users WHERE email = ?", email)
	if error != nil {
		fmt.Println(error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "数据库错误"})
		return // 发生错误时提前返回
	}
	defer rowA.Close() // 确保rowA被关闭
	if rowA.Next() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "邮箱已存在"})
		return // 邮箱已存在时提前返回
	}

	username := c.PostForm("username")
	row, err := db.Query("SELECT email FROM users WHERE username = ?", username)
	if err != nil {
		fmt.Println(error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "数据库错误"})
		return
	}
	if row.Next() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "用户名已存在"})
		return
	}

	row.Close()
	password := c.PostForm("password")
	password = md5V(password)
	if email == "" || username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "信息不完整，请检查邮箱、用户名、密码是否填写完整"})
		return
	}

	//默认验证码为空。如果再次发送验证码，会覆盖之前的验证码
	randCode := getRandomCode(6)
	session.Values["email"] = email
	session.Values["username"] = username
	session.Values["password"] = password
	session.Values["code"] = randCode

	session.Options.MaxAge = 300
	session.Save(c.Request, c.Writer)
	if sendEmail(email, "register", randCode) {
		log.Println("sendCode", email, username)
		c.JSON(http.StatusOK, gin.H{"message": "验证码已发送"})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证码发送失败，请检查邮箱是否正确"})
		return
	}
}

func register(c *gin.Context) {// get email or send email
	session, err := Store.Get(c.Request, "userInfo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请先发送验证码，如果已经发送，请检查是否禁用了cookie"})
		return
	}
	email := session.Values["email"].(string)
	username := session.Values["username"].(string)
	password := session.Values["password"].(string)
	code := c.PostForm("code")
	if email == "" || username == "" || password == "" || code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "信息不完整，请检查邮箱、用户名、密码、验证码是否填写完整"})
		return
	}
	if session.Values["code"] == code {
		_, err := db.Exec("INSERT INTO users (email,username,password,follow) VALUES (?, ?, ?,?)", email, username, password,false)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "注册失败，数据库错误"})
			return
		}
		session.Options.MaxAge = 0
		session.Save(c.Request, c.Writer)
		log.Println("register", email, username)
		c.JSON(http.StatusOK, gin.H{"message": "注册成功", "code": 1})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证码错误"})
		return
	}
}

func login(c *gin.Context) {
	// examine email and password
	session, _ := Store.Get(c.Request, "userInfo")
	email := c.PostForm("email")
	password := c.PostForm("password")
	password = md5V(password)
	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "信息不完整，请检查邮箱、密码是否填写完整"})
		return
	}
	//database action
	rows, err := db.Query("SELECT * FROM users WHERE email = ? AND password = ?", email, password)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"message": "请检查账号密码是否正确"})
		return
	}
	defer rows.Close()
	if rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Follow); err != nil {
			log.Panicln("读取数据时发生错误:", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "数据库错误"})
			return
		}
		session.Values["id"] = user.Id
		session.Values["email"] = user.Email
		session.Values["username"] = user.Username
		session.Values["follow"] = user.Follow
		session.Options.MaxAge = 0
		session.Save(c.Request, c.Writer)
		c.JSON(http.StatusOK, gin.H{"message": "success", "code": 1})
		return
	}
}
