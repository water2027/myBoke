package main

import (
	"fmt"
	"regexp"
	"gopkg.in/gomail.v2"
)

var pwd string = "pwd"
var sendBy string = "123456@qq.com"
var mailHost string = "smtp.qq.com"
var mailPort int = 465

func verifyEmail(email string) bool {
	var emailReg = `^([\w\.\_]{2,10})@(\w{2,10})\.([a-z]{2,4})$`
	re := regexp.MustCompile(emailReg)
	return re.MatchString(email)
}

func sendEmail(sendTo string, subject string, body string) bool {
	if !verifyEmail(sendTo) {
		panic("email format error")
	}
	m := gomail.NewMessage()
	m.SetHeader("From", sendBy)
	m.SetHeader("To", sendTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(mailHost, mailPort, sendBy, pwd)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}




