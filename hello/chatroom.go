package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var (
	connections = make(map[*websocket.Conn]interface{})
	mu          sync.RWMutex
)

var Messages []map[string]interface{}

func broadcastMessage(ctx context.Context, msg interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	for conn := range connections {
		err := wsjson.Write(ctx, conn, msg)
		if err != nil {
			log.Printf("Error broadcasting message: %v", err)
		}
	}
}

func websocketChat(c *gin.Context) {
	options := &websocket.AcceptOptions{
		OriginPatterns: []string{"watering.top"},
	}

	conn, err := websocket.Accept(c.Writer, c.Request, options)
	if err != nil {
		log.Printf("WebSocket accept error: %v", err)
		if strings.Contains(err.Error(), "not authorized") {
			c.String(http.StatusForbidden, "Forbidden")
		} else {
			c.String(http.StatusInternalServerError, "Failed to establish WebSocket connection")
		}
		return
	}
    session,err := Store.Get(c.Request, "userInfo")
    if err != nil {
        log.Printf("Error getting session: %v", err)
        c.JSON(http.StatusProxyAuthRequired, gin.H{"error": "please login first"})
    }

	username := session.Values["username"]

	mu.Lock()
	connections[conn] = username
	mu.Unlock()

	firstMsg := map[string]interface{}{
		"username": "system",
		"message":  "欢迎来到聊天室",
		"time":     time.Now().Format("2006-1-2 15:04"),
		"formerMsg": Messages,
	}

	err = wsjson.Write(c.Request.Context(), conn, firstMsg)
	if err != nil {
		log.Printf("Error writing welcome message: %v", err)
	}

	defer func() {
		mu.Lock()
		delete(connections, conn)
		mu.Unlock()
		conn.Close(websocket.StatusNormalClosure, "")
	}()

	ctx := c.Request.Context()

	for {
		var v map[string]interface{}
		err := wsjson.Read(ctx, conn, &v)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		msg, ok := v["message"].(string)
		if !ok || len(msg) > 1000 {
			continue
		}

        now := time.Now()
        bjTime := now.In(time.FixedZone("CST", 8*3600))

		newMsg := map[string]interface{}{
			"username": username,
			"message":  msg,
			"time":     bjTime.Format("2006-1-2 15:04"),
		}

		if len(Messages) <= 10 {
			Messages = append(Messages, newMsg)
		} else {
			Messages = Messages[1:]
			Messages = append(Messages, newMsg)
		}

		broadcastMessage(ctx, newMsg)
	}
}