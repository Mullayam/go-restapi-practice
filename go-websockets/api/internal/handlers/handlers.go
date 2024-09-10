package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var wsChan = make(chan WsPayload)
var clients = make(map[WebSocketConection]string)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("../../html"),
	jet.InDevelopmentMode(),
)
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WsJsonResponse struct {
	Message        string   `json:"message"`
	Action         string   `json:"action"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}
type WebSocketConection struct {
	*websocket.Conn
}
type WsPayload struct {
	Message  string             `json:"message"`
	Username string             `json:"username"`
	Action   string             `json:"action"`
	Conn     WebSocketConection `json:"-"`
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected...")
	var response WsJsonResponse
	response.Message = "Hi Client"
	conn := WebSocketConection{Conn: ws}
	clients[conn] = ""

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}
	go ListenToWs(&conn)
}
func ListenToWs(conn *WebSocketConection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()
	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			log.Println(err)
			_ = conn.Close()
			delete(clients, *conn)
			break
		}
		payload.Conn = *conn
		wsChan <- payload

	}

}
func ListenToWsChannel() {
	var response WsJsonResponse

	for {
		e := <-wsChan
		switch e.Action {
		case "username":
			clients[e.Conn] = e.Username
			users := getUserList()
			response.Action = "list_users"
			response.ConnectedUsers = users
			broadCastToAll(response)

		case "left":
			response.Action = "list_users"
			delete(clients, e.Conn)
			users := getUserList()
			response.ConnectedUsers = users
			broadCastToAll(response)
		case "broadcast":
			response.Action = "list_users"
			response.Message = fmt.Sprintf("%s says %s", e.Username, e.Message)
			broadCastToAll(response)

		}
		// response.Message = e.Message
		// 		response.Action = "Got Here"
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
func broadCastToAll(response WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println(err)
			_ = client.Close()
			delete(clients, client)
		}
	}

}
func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)

		return err
	}
	w.Header().Set("Content-Type", "text/html")
	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)

		return err
	}
	return nil
}
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
