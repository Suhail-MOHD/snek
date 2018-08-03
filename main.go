package snek

type vector struct {
	x int
	y int
}

type coord struct {
	x int
	y int
}

type snake struct {
	pos       []coord
	direction vector
	dir       int
}

type level struct {
	snake *snake
}
