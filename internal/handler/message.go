package handler

import (
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sort"
)

var wsChan = make(chan model.MessagePayload)
var clients = make(map[model.WebSocketConnection]string)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) wsHandler(c *gin.Context) {
	ws, err := upgradeConnection.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		h.log.Errorf("can't upgrade to WS: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	h.log.Info("client connected to ws endpoint")

	response := model.MessageResponse{
		Message: `Connected to server!!!`,
	}

	conn := model.WebSocketConnection{
		Conn: ws,
	}
	clients[conn] = ""

	err = ws.WriteJSON(response)
	if err != nil {
		h.log.Errorf("can't send message: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, model.Response{
			IsSuccess: false,
			Message:   err.Error(),
		})
		return
	}

	go h.listenForWs(&conn)
}

func (h *Handler) listenForWs(conn *model.WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			h.log.Errorf("error: %v", r)
		}
	}()

	var payload model.MessagePayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func (h *Handler) listenToWsChannel() {
	var response model.MessageResponse

	for {
		e := <-wsChan

		switch e.Action {
		case "username":
			clients[e.Conn] = e.Username
			users := getUserList()
			response.Action = "list_users"
			response.ConnectedUsers = users
			h.broadcastToAll(response)

		case "left":
			response.Action = "list_users"
			delete(clients, e.Conn)
			users := getUserList()
			response.ConnectedUsers = users
			h.broadcastToAll(response)
		}

		//response.Action = "Got here"
		//response.Message = fmt.Sprintf("action %s", e.Action)
		//h.broadcastToAll(response)
	}
}

func getUserList() []string {
	var userList []string
	for _, x := range clients {
		if x != "" {
			userList = append(userList, x)
		}
	}

	sort.Strings(userList)
	return userList
}

func (h *Handler) broadcastToAll(response model.MessageResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			h.log.Errorf("websocker err: %s", err.Error())
			_ = client.Close()
			delete(clients, client)
		}
	}
}
