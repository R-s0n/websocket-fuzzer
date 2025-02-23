package main
import (
	"flag"
	"log"
	"os"
	"github.com/gorilla/websocket"
)
func main() {
	uri := flag.String("uri", "", "WebSocket URI")
	message := flag.String("message", "", "Message to send")
	flag.Parse()
	if *uri == "" || *message == "" {
		log.Println("Usage: go run main.go -uri ws://localhost:1234 -message YOUR_MESSAGE")
		os.Exit(1)
	}
	c, _, err := websocket.DefaultDialer.Dial(*uri, nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer c.Close()
	for {
		err := c.WriteMessage(websocket.TextMessage, []byte(*message))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
