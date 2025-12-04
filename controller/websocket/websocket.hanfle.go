package websocket

import (
	"encoding/json"
	"fmt"
	"guntingbatukertas/service"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	playService *service.PlayService
}

func NewWebSocketHandler(ps *service.PlayService) *WebSocketHandler {
	return &WebSocketHandler{
		playService: ps,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *WebSocketHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	type CreateRoomPayload struct {
		PlayerName string `json:"playerName"`
		RoomName   string `json:"roomName"`
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		var payload CreateRoomPayload
		if err := json.Unmarshal(message, &payload); err != nil {
			fmt.Println("Invalid JSON:", err)
			conn.WriteMessage(websocket.TextMessage, []byte("invalid json"))
			continue
		}

		res := h.playService.CreateRoom(payload.PlayerName, payload.RoomName)

		if res != nil {
			switch v := res.(type) {
			case error:
				conn.WriteMessage(websocket.TextMessage, []byte(v.Error()))
			case string:
				conn.WriteMessage(websocket.TextMessage, []byte(v))
			default:
				conn.WriteMessage(websocket.TextMessage, []byte("unknown response"))
			}
			continue
		}

		conn.WriteMessage(websocket.TextMessage, []byte("room created"))
	}
}
