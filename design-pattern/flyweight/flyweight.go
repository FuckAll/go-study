package flyweight

import "sync"

/*
	享元模式：所谓“享元”，顾名思义就是被共享的单元。享元模式的意图是复用对象，节省内存，前提是享元对象是不可变对象。
	// 设计一个棋盘游戏大厅，每个房间都有一个棋盘，每个棋盘都有30个棋子。
	// 如果每个房间都是创建一个棋盘和30个棋子，随着游戏大厅房间的增多，会有越来越多的
	// 棋盘对象和棋子对象，但是对于棋子来说，每个房间的30个棋子都是相同的。
*/

type Color string

const (
	Red   Color = "红"
	Black Color = "黑"
)

// 棋子是享元
type ChessPieceUnit struct {
	id    int
	text  string
	color Color
}

func NewChessPieceUnit(id int, text string, color Color) *ChessPieceUnit {
	return &ChessPieceUnit{
		id:    id,
		text:  text,
		color: color,
	}
}

var chessPiecesMap map[int]*ChessPieceUnit
var once sync.Once

func GetChessPiece(chessPieceId int) *ChessPieceUnit {
	once.Do(func() {
		//  初始化30个棋子
		chessPiecesMap[1] = &ChessPieceUnit{id: 1, text: "車", color: Red}
		chessPiecesMap[2] = &ChessPieceUnit{id: 2, text: "馬", color: Black}
		// ....
	})
	return chessPiecesMap[chessPieceId]
}

type ChessPiece struct {
	chessPieceUnit *ChessPieceUnit
	positionX      int
	positionY      int
}

func NewChessPiece(chessPieceUnit *ChessPieceUnit, positionX int, positionY int) *ChessPiece {
	return &ChessPiece{chessPieceUnit: chessPieceUnit, positionX: positionX, positionY: positionY}
}

type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

func NewChessBoard() *ChessBoard {
	// 初始化棋盘
	newBoard := &ChessBoard{}
	newBoard.chessPieces[1] = &ChessPiece{chessPieceUnit: GetChessPiece(1), positionX: 0, positionY: 0}
	newBoard.chessPieces[2] = &ChessPiece{chessPieceUnit: GetChessPiece(2), positionX: 0, positionY: 0}
	// ...
	return newBoard
}

// 移动棋子
func (c *ChessBoard) move(chessPieceId int, positionX int, positionY int) {
	c.chessPieces[chessPieceId].positionX = positionX
	c.chessPieces[chessPieceId].positionY = positionY
}
