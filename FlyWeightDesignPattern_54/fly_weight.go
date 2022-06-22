package FlyWeightDesignPattern_54

var chessPieceUnit = map[string]*ChessPieceUnit{
	"红兵1": &ChessPieceUnit{
		Id:    1,
		Name:  "兵",
		Color: "red",
	},
	"黑兵1": &ChessPieceUnit{
		Id:    2,
		Name:  "卒",
		Color: "black",
	},
}

type ChessPieceUnit struct {
	Id    int
	Name  string
	Color string
}

func NewChessPieceUnit(s string) *ChessPieceUnit {
	return chessPieceUnit[s]
}

type ChessPiece struct {
	ChessPiece ChessPieceUnit
	Px         int
	Py         int
	IsOut      bool // 是否被吃掉
}
type ChessBoard struct {
	chessPieces []*ChessPiece
}

func NewChessBoard(chessPieces []*ChessPiece) *ChessBoard {
	// 初始化各个棋子的位置信息和
	return &ChessBoard{chessPieces: chessPieces}
}
