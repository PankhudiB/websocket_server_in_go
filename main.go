package main

import (
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

	_, message, _ := websocketConn.ReadMessage()
	fmt.Println("Message: ", string(message))

	err = websocketConn.WriteMessage(websocket.TextMessage, []byte("Hello from server!\n"))
	if err != nil {
		fmt.Println("Error writing to ws connection !", err.Error())
	}
}
