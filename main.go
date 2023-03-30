package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
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

	messageA := MessageWrapper{
		MessageType:   "A",
		MessageType_A: MessageType_A{Name: "Pankhudi", Place: "India"},
	}

	bytes, err := json.Marshal(messageA)

	fmt.Println("Sending bytes :", string(bytes))
	err = websocketConn.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		fmt.Println("Error writing to ws connection !", err.Error())
	}
}

type MessageType_A struct {
	Name  string `json:"name"`
	Place string `json:"place"`
}

type MessageType_B struct {
	Animal string `json:"animal"`
	Thing  string `json:"thing"`
}

type MessageWrapper struct {
	MessageType   string `json:"message_type"`
	MessageType_A `json:"content"`
}
