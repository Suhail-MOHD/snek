package common

type VectorDat struct {
	X int
	Y int
}

type CoordDat struct {
	X int
	Y int
}

type SnakeDat struct {
	Pos []CoordDat
}

type WallBlockDat struct {
	Pos CoordDat
}

type FruitDat struct {
	Pos CoordDat
}

type BoardDat struct {
	Width  int64
	Length int64
}
