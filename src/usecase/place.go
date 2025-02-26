package usecase

import (
	"connectrpc.com/connect"
	"context"
	"ebitengine-othello/src/config"
	"ebitengine-othello/src/domain"
	othellov1 "ebitengine-othello/src/gen/othello/v1"
	"ebitengine-othello/src/gen/othello/v1/othelloconnect"
	"fmt"
	"net/http"
)

func Place(g *domain.GameStatus, y int, x int) {

	if g.Board[y][x] == config.CELL_EMPTY {
		return
	}

	if g.Board[y][x] == config.CELL_LEGAL {
		// ひっくり返し処理
		g.PlayerPass = false
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				if g.Board[i][j] == config.CELL_LEGAL {
					g.Board[i][j] = config.CELL_EMPTY
				}
			}
		}
		g.Board[y][x] = g.Side

		currentSide := g.Side
		var enemySide int
		if currentSide == config.CELL_BLACK {
			enemySide = config.CELL_WHITE
		} else {
			enemySide = config.CELL_BLACK
		}

		// 各方向ごとにひっくり返しを実施する
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}

				flipPositions := []map[string]int{}
				px, py := x+dx, y+dy
				for {
					// 盤外なら終了
					if px < 0 || px >= 8 || py < 0 || py >= 8 {
						break
					}

					// 空白なら終了
					if g.Board[py][px] == config.CELL_EMPTY {
						break
					}

					// 敵の駒なら記録して次へ
					if g.Board[py][px] == enemySide {
						flipPositions = append(flipPositions, map[string]int{"x": px, "y": py})
						px += dx
						py += dy
						continue
					}

					// 自分の駒なら、これまでの駒をひっくり返す
					if g.Board[py][px] == currentSide {
						if len(flipPositions) > 0 {
							for _, pos := range flipPositions {
								g.Board[pos["y"]][pos["x"]] = currentSide
							}
						}
						break
					}
				}
			}
		}

		fmt.Println(" ")
		fmt.Println("Playerの手: ", y, x)
		g.Board[y][x] = currentSide
		CountStones(g)
		fmt.Println("黒: ", g.Black, "白: ", g.White)
		fmt.Println(" ")
		g.Side = enemySide

		return
	}
	// クリックした位置に駒がある場合は何もしない
	if g.Board[y][x] != config.CELL_EMPTY {
		return
	}

	// ひっくり返しが発生しないなら、設置しない
	if canPlace(g.Board, y, x, g.Side) {
		return
	}
}

func PlaceAi(g *domain.GameStatus, player int) {
	client := othelloconnect.NewOthelloServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		connect.WithGRPC(),
	)

	var board []int32
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board = append(board, int32(g.Board[i][j]))
		}
	}
	res, err := client.GetAIMove(
		context.Background(),
		connect.NewRequest(&othellov1.GetAIMoveRequest{
			Board:  board,
			Player: int32(player),
		}),
	)

	if err != nil {
		fmt.Println("AIの手: パス")
		g.Side = g.Player
		g.AiPass = true
		return
	}

	y, x := int(res.Msg.GetY()), int(res.Msg.GetX())
	g.AiPass = false
	currentSide := g.Side
	var enemySide int
	if currentSide == config.CELL_BLACK {
		enemySide = config.CELL_WHITE
	} else {
		enemySide = config.CELL_BLACK
	}

	// 各方向ごとにひっくり返しを実施する
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			flipPositions := []map[string]int{}
			px, py := x+dx, y+dy
			for {
				// 盤外なら終了
				if px < 0 || px >= 8 || py < 0 || py >= 8 {
					break
				}

				// 空白なら終了
				if g.Board[py][px] == config.CELL_EMPTY {
					break
				}

				// 敵の駒なら記録して次へ
				if g.Board[py][px] == enemySide {
					flipPositions = append(flipPositions, map[string]int{"x": px, "y": py})
					px += dx
					py += dy
					continue
				}

				// 自分の駒なら、これまでの駒をひっくり返す
				if g.Board[py][px] == currentSide {
					if len(flipPositions) > 0 {
						for _, pos := range flipPositions {
							g.Board[pos["y"]][pos["x"]] = currentSide
						}
					}
					break
				}
			}
		}
	}

	g.Board[y][x] = currentSide
	g.Side = g.Player
	fmt.Println(" ")
	fmt.Println("AIの手: ", y, x)
	CountStones(g)
	fmt.Println("黒: ", g.Black, "白: ", g.White)
	fmt.Println("")
}

func canPlace(board [8][8]int, y int, x int, currentSide int) bool {
	// 指定セルが空かチェック
	if board[y][x] != config.CELL_EMPTY {
		return false
	}

	var enemySide int
	if currentSide == config.CELL_BLACK {
		enemySide = config.CELL_WHITE
	} else {
		enemySide = config.CELL_BLACK
	}

	// ひっくり返す可能性がある方向があれば false を返す
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			canTurnOver := false
			px, py := x+dx, y+dy
			i := 1
			for {
				// 盤外なら終了
				if px < 0 || px >= 8 || py < 0 || py >= 8 {
					break
				}

				// 空または1マス目で自分の駒なら終了
				if board[py][px] == config.CELL_EMPTY || (board[py][px] == currentSide && i == 1) {
					break
				}

				// 相手の駒なら次へ
				if board[py][px] == enemySide {
					px += dx
					py += dy
					i++
					continue
				}

				// 自分の駒であれば、ひっくり返しが発生する
				if board[py][px] == currentSide {
					canTurnOver = true
					break
				}
			}
			if canTurnOver {
				return false
			}
		}
	}

	return true
}

func inside(y int, x int) bool {
	return 0 <= y && y < 8 && 0 <= x && x < 8
}

func CheckLegal(b *domain.GameStatus, player int) bool {
	found := false
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if b.Board[y][x] == config.CELL_LEGAL {
				b.Board[y][x] = config.CELL_EMPTY
			}
		}
	}
	var dy = [8]int{0, 1, 0, -1, 1, 1, -1, -1}
	var dx = [8]int{1, 0, -1, 0, 1, -1, 1, -1}

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if b.Board[y][x] != config.CELL_EMPTY {
				continue
			}
			for d := 0; d < 8; d++ {
				ny, nx := y+dy[d], x+dx[d]
				if !inside(ny, nx) || b.Board[ny][nx] != b.Ai {
					continue
				}
				for steps := 2; ; steps++ {
					ny, nx = y+dy[d]*steps, x+dx[d]*steps
					if !inside(ny, nx) || b.Board[ny][nx] == config.CELL_EMPTY {
						break
					}
					if b.Board[ny][nx] == player {
						b.Board[y][x] = config.CELL_LEGAL
						found = true
						break
					}
				}
			}
		}
	}
	return found
}

func CountStones(b *domain.GameStatus) {
	b.Black = 0
	b.White = 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if b.Board[i][j] == config.CELL_BLACK {
				b.Black++
			} else if b.Board[i][j] == config.CELL_WHITE {
				b.White++
			}
		}
	}
}
