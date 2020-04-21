package main

import (
	"fmt"
	"github.com/SnakeRoyale/snake-royale-backend/repository"
	"github.com/SnakeRoyale/snake-royale-backend/service"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func main() {
	Register()

}

func Register() {
	// init repositories & facades
	gameRepository := repository.NewGameRepository()

	// api for frontend
	gameApi := service.NewGame(gameRepository)

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	gameApi.RegisterEvents(server)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		s.Join("snakeRoom")
		return nil
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})



	/*

	server.OnEvent("/notice", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

*/




	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}