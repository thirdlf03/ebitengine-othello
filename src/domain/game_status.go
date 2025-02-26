package domain

import "github.com/gorilla/websocket"

type GameStatus struct {
	Board      [8][8]int
	Side       int
	Player     int
	Ai         int
	Black      int
	White      int
	PlayerPass bool
	AiPass     bool
	Score      int
	Help       bool
	Conn       *websocket.Conn
	Channel    chan int
}
