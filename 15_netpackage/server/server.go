package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/google/uuid"
)

type CmdSession int

type ChatSession struct {
	ChatID       string
	Participants [2]string
}

const (
	Register CmdSession = iota
	QueryParticipant
	Invite
	AcceptInvite
	SendMessage
	TerminateChat
)

type RegisterRequest struct {
	UUID   string `json:UUID`
	Mobile string `json:Mobile`
}

type RegisterResponse struct {
	UUID               string `json:UUID`
	RegistrationStatus string `json:RegistrationStatus`
}

type QueryParticipantRequest struct {
	UUID string `json:UUID`
}

type QueryParticipantResponse struct {
	UUID     string   `json:UUID`
	UserList []string `json:UserList`
}

type InviteRequest struct {
	FromUUID string `json:From`
	ToUUID string `json:To`
}

type InviteResponse struct {
	SessionID string `json:SessionID`
	SessionStatus string `json:SessionStatus`
}

type AcceptInviteRequest struct {
	
}

type 

type SessionCommand struct {
	Command           CmdSession `json:ChatCommand`
	RegisterRequest   `json:SendMessage`
	QueryParticipantRequest `json:EndChat`
	InviteData        `json:Invite`
}

type user struct {
	username   string
	connection net.Conn
}

type App struct {
	server_add      string
	protocol        string
	listner         net.Listener
	users           []user
	ChatSessionList []ChatSession
}

var app App

func init() {
	uuid.New()

	users := make([]user, 100)
	app = App{server_add: "0.0.0.0:5221", protocol: "tcp", users: users}
}

func (a *App) Listen() error {
	listner, err := net.Listen(a.protocol, a.server_add)
	a.listner = listner
	return err
}

func (a *App) Accept() (user, error) {
	conn, err := a.listner.Accept()
	usr := user{connection: conn}
	return usr, err
}

func handleRead(usr user) {
	for {
		data := make([]byte, 2048)
		n, err := usr.connection.Read(data)
		if n == 0 || err != nil {
			continue
		}
		fmt.Println("\nReceived:" + string(data))
	}

}

func (usr *user) handleConnection() {
	for {
		command := RequestUUIDData{Command: RequestUUID}
		data, err := json.Marshal(command)
		if err != nil {
			log.Fatal("failed to marshal command")
		}
		usr.connection.Write(data)
		n, err := usr.connection.Read(data)

		if n == 0 || err != nil {
			continue
		}
		usr.username = string(data)
		app.users = append(app.users, *usr)
		break
	}

	go handleRead(*usr)
}

func readFromStdin() []byte {
	var reader io.Reader = os.Stdin
	data := make([]byte, 512)
	fmt.Print("Sent:")
	reader.Read(data)

	//fmt.Println(string(data))
	return data
}

func main() {

	fmt.Println("Welcome to TCP chat server")

	fmt.Println("Read data from screen")

	if app.Listen() != nil {
		log.Fatal("Failed to start server")
	}

	for {
		usr, err := app.Accept()

		if err != nil {
			log.Print("Failed to connect...")
			continue
		}
		go usr.handleConnection()
	}
}
