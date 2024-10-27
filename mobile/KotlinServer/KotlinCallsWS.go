package KotlinServer

import (
	"encoding/hex"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}
var url = "ws://localhost:8080"

func requestFile(filehash string) {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	// Request for a file
	err = c.WriteMessage(websocket.TextMessage, []byte(filehash))
	if err != nil {
		log.Println("write:", err)
		return
	}
	//for {
	messageType, message, err := c.ReadMessage()
	if err != nil {
		log.Println("Error during message reading:", err)
		//break
	}
	if messageType == websocket.BinaryMessage {
		fileContent := hex.EncodeToString(message)
		fmt.Println(fileContent)
	} else {
		log.Println("Received non-binary message")
	}
	log.Printf("Received: %s", message)
	//}
}
func receiveFile(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

	for {
		messageType, message, _ := conn.ReadMessage() // error ignored for simplicity

		// Now "message" should contain the file content.
		if messageType == websocket.BinaryMessage {
			fileContent := hex.EncodeToString(message)
			fmt.Println(fileContent)
		} else {
			log.Println("Received non-binary message")
		}
	}
}

func main() {
	http.HandleFunc("/readFile", receiveFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
