package main

import (
	"ebitengine-othello/src/config"
	"ebitengine-othello/src/domain"
	"ebitengine-othello/src/usecase"
	"ebitengine-othello/src/utils"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
	"os"
	"time"
)

type EbitenGame struct {
	GameStatus domain.GameStatus
	fontFace   *text.GoTextFace
	gameOver   bool
}

func (g *EbitenGame) Update() error {

	if g.GameStatus.Black+g.GameStatus.White == 64 || (g.GameStatus.PlayerPass && g.GameStatus.AiPass) {
		if !g.gameOver {
			fmt.Println("ゲーム終了")
			if g.GameStatus.Black > g.GameStatus.White {
				fmt.Println("黒の勝ち")
			} else if g.GameStatus.Black < g.GameStatus.White {
				fmt.Println("白の勝ち")
			} else {
				fmt.Println("引き分け")
			}
		}
		g.gameOver = true
		return nil
	}
	// ウィンドウ外にカーソルがある場合は何もしない
	cursorX, cursorY := ebiten.CursorPosition()
	targetCol := cursorY / (config.CELL_LENGTH + config.GRID_WIDTH)
	targetRow := cursorX / (config.CELL_LENGTH + config.GRID_WIDTH)
	if targetRow < 0 || targetRow >= 8 || targetCol < 0 || targetCol >= 8 {
		return nil
	}

	if g.GameStatus.Side == g.GameStatus.Ai {
		usecase.PlaceAi(&g.GameStatus, g.GameStatus.Ai)
		time.Sleep(700 * time.Millisecond)
		usecase.CountStones(&g.GameStatus)
		if g.GameStatus.Player == config.CELL_WHITE {
			found := usecase.CheckLegal(&g.GameStatus, g.GameStatus.Player)
			if !found {
				g.GameStatus.Side = g.GameStatus.Ai
				fmt.Printf("人間の手: パス\n")
				g.GameStatus.PlayerPass = true
			}
		}
		return nil
	} else {
		found := usecase.CheckLegal(&g.GameStatus, g.GameStatus.Player)
		if !found {
			g.GameStatus.Side = g.GameStatus.Ai
			fmt.Printf("人間の手: パス\n")
			g.GameStatus.PlayerPass = true
		}
	}

	if g.GameStatus.Score >= 1000 && !g.GameStatus.Help && g.GameStatus.Side == g.GameStatus.Player {
		g.GameStatus.Help = true
		fmt.Println("お助けくん")
	}

	// クリックした位置に駒がある場合は何もしない
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		return nil
	}
	usecase.Place(&g.GameStatus, targetCol, targetRow)
	return nil
}

func (g *EbitenGame) Draw(screen *ebiten.Image) {

	board := &g.GameStatus.Board
	turn := &g.GameStatus.Side

	// オセロ盤のベースの部分（緑の四角）
	vector.DrawFilledRect(
		screen,
		float32(config.OUTER_MARGIN)+10,
		float32(config.OUTER_MARGIN),
		float32(config.BOARD_LENGTH)+10,
		float32(config.BOARD_LENGTH),
		color.RGBA{0x00, 0x80, 0x00, 0xff},
		false,
	)

	// グリッドを引く
	for i := 1; i < 8; i++ {
		// 縦
		vector.StrokeLine(
			screen,
			float32(config.GRID_WIDTH*i+13+config.CELL_LENGTH*i),
			0,
			float32(config.GRID_WIDTH*i+13+config.CELL_LENGTH*i),
			float32(config.SCREEN_LENGTH),
			float32(config.GRID_WIDTH),
			color.Black,
			false,
		)

		// 横
		vector.StrokeLine(
			screen,
			0,
			float32(config.GRID_WIDTH*i+config.CELL_LENGTH*i),
			float32(config.SCREEN_LENGTH),
			float32(config.GRID_WIDTH*i+config.CELL_LENGTH*i),
			float32(config.GRID_WIDTH),
			color.Black,
			false,
		)
	}

	// オセロの石を描画
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			piece := board[y][x]
			if piece == config.CELL_EMPTY {
				continue
			}

			var pieceColor color.Color
			if piece == config.CELL_BLACK {
				pieceColor = color.Black
			} else if piece == config.CELL_WHITE {
				pieceColor = color.White
			} else {
				pieceColor = color.RGBA{0xff, 0xc0, 0xcb, 0xff}
			}

			vector.DrawFilledCircle(
				screen,
				float32(config.CELL_LENGTH/2+config.CELL_LENGTH*x+config.GRID_WIDTH*x+14),
				float32(config.CELL_LENGTH/2+config.CELL_LENGTH*y+config.GRID_WIDTH*y+2),
				float32(config.CELL_LENGTH/2)-2,
				pieceColor,
				true,
			)
		}
	}

	gray := color.RGBA{0x80, 0x80, 0x80, 0xff}

	if !g.gameOver {
		stoneOption := &text.GoTextFace{Source: g.fontFace.Source, Size: 30}
		w, h := text.Measure(fmt.Sprintf("黒 - %d vs 白 - %d", g.GameStatus.Black, g.GameStatus.White), stoneOption, 10)

		x := (config.SCREEN_LENGTH - w) / 2
		y := config.SCREEN_LENGTH - h

		vector.DrawFilledRect(screen, float32(x), float32(y), float32(w), float32(h), gray, false)

		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		op.LineSpacing = 10

		text.Draw(screen, fmt.Sprintf("黒 - %d vs 白 - %d", g.GameStatus.Black, g.GameStatus.White), stoneOption, op)
	} else {
		var result string
		result = fmt.Sprintf("黒 - %d vs 白 - %d\n\n", g.GameStatus.Black, g.GameStatus.White)
		if g.GameStatus.Black == g.GameStatus.White {
			result += "引き分け"
		} else if g.GameStatus.Black > g.GameStatus.White {
			if g.GameStatus.Ai == config.CELL_BLACK {
				result += "黒 (AI) の勝ち"
			} else {
				result += "黒 (Player) の勝ち"
			}
		} else {
			if g.GameStatus.Ai == config.CELL_WHITE {
				result += "白 (AI) の勝ち"
			} else {
				result += "白 (Player) の勝ち"
			}
		}

		stoneOption := &text.GoTextFace{Source: g.fontFace.Source, Size: 13}
		w, h := text.Measure(result, stoneOption, 10)

		x := (config.SCREEN_LENGTH - w) / 2
		y := config.SCREEN_LENGTH - h

		vector.DrawFilledRect(screen, float32(x), float32(y), float32(w), float32(h), gray, false)

		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		op.LineSpacing = 10

		text.Draw(screen, result, stoneOption, op)
	}

	// ホバーした位置に薄く駒を表示
	cursorX, cursorY := ebiten.CursorPosition()
	targetRow := cursorY / (config.CELL_LENGTH + config.GRID_WIDTH)
	targetCol := cursorX / (config.CELL_LENGTH + config.GRID_WIDTH)
	if targetRow < 0 || targetRow >= 8 || targetCol < 0 || targetCol >= 8 {
		return
	}

	if board[targetRow][targetCol] == config.CELL_EMPTY || board[targetRow][targetCol] == config.CELL_LEGAL {
		var pieceColor color.Color
		if *turn == config.CELL_BLACK {
			pieceColor = color.RGBA{0x00, 0x00, 0x00, 0x77}
		} else if *turn == config.CELL_WHITE {
			pieceColor = color.RGBA{0xaa, 0xaa, 0xaa, 0x77}
		}

		vector.DrawFilledCircle(
			screen,
			float32(config.CELL_LENGTH/2+config.CELL_LENGTH*targetCol+config.GRID_WIDTH*targetCol+14),
			float32(config.CELL_LENGTH/2+config.CELL_LENGTH*targetRow+config.GRID_WIDTH*targetRow+2),
			float32(config.CELL_LENGTH/2)-2,
			pieceColor,
			true,
		)
	}
}

func (g *EbitenGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.SCREEN_LENGTH, config.SCREEN_LENGTH
}

func main() {
	fmt.Print("AIの手番 1: 黒(先手) 2: 白(後手) : ")
	input := 0
	_, err := fmt.Scan(&input)
	if err != nil || (input != 1 && input != 2) {
		fmt.Println("1か2を入力してください")
		os.Exit(1)
	}

	player := utils.ConvertPlayer(input)

	// 石の数を表示
	f, err := os.Open("NotoSansJP-Medium.ttf")
	if err != nil {
		return
	}
	defer f.Close()

	src, err := text.NewGoTextFaceSource(f)
	if err != nil {
		return
	}

	ebiten.SetWindowSize(config.WINDOW_WIDTH, config.WINDOW_HEIGHT+100)
	ebiten.SetWindowTitle("Hello, World!")

	game := &EbitenGame{
		GameStatus: domain.GameStatus{
			Board: [8][8]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 2, 1, 0, 0, 0},
				{0, 0, 0, 1, 2, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			Side:   config.CELL_BLACK,
			Player: player,
			Ai:     input,
		},
		fontFace: &text.GoTextFace{Source: src, Size: 30},
	}

	fmt.Println("Player vs AI")
	fmt.Println(" ")
	if input == 2 {
		fmt.Println("Player: 黒")
		fmt.Println("AI: 白")
	} else {
		fmt.Println("Player: 白")
		fmt.Println("AI: 黒")
	}
	fmt.Println(" ")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
