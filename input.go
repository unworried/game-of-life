package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Dir int

const (
	DirRight Dir = iota
	DirLeft
)

func (d Dir) String() string {
	switch d {
	case DirRight:
		return "Right"
	case DirLeft:
		return "Left"
	}
	panic("not reach")
}

type Input struct{}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Dir() (Dir, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		return DirLeft, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		return DirRight, true
	}
	return 0, false
}
