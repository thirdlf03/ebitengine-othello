package domain

type GameStatus struct {
	Board      [8][8]int
	Side       int
	Player     int
	Ai         int
	Black      int
	White      int
	PlayerPass bool
	AiPass     bool
}
