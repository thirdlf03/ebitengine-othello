package config

const (
	WINDOW_WIDTH  = 640
	WINDOW_HEIGHT = 480

	SCREEN_LENGTH = 380

	BOARD_LENGTH = 330 - OUTER_MARGIN*2
	OUTER_MARGIN = 1
	GRID_WIDTH   = 1
	CELL_LENGTH  = (330 - OUTER_MARGIN*2) / 8

	CELL_EMPTY = 0
	CELL_BLACK = 1
	CELL_WHITE = 2
	CELL_LEGAL = 3

	INF = 1000000000
)
