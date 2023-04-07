package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"websocket-server/model"
)

func main() {
	httpEngine := gin.Default()

	fmt.Println("Starting WebSocket Server...")

	httpEngine.GET("/talk-to-server", handleWebsocket)
	fmt.Println()

	err := http.ListenAndServe(":8080", httpEngine)
	if err != nil {
		fmt.Println("Error starting server!")
		return
	}
}

func handleWebsocket(context *gin.Context) {
	upgrader := websocket.Upgrader{}
	websocketConn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading connection !", err.Error())
	}

	if write(websocketConn) != nil {
		fmt.Println("Error writing to ws connection !", write(websocketConn).Error())
	}
}

func write(websocketConn *websocket.Conn) error {
	messageA := model.MessageWrapper{
		MessageType:   "A",
		MessageType_A: model.MessageType_A{Name: "Pankhudi", Place: "India"},
	}
	bytes, err := json.Marshal(messageA)
	if err != nil {
		return err
	}
	return websocketConn.WriteMessage(websocket.TextMessage, bytes)
}
