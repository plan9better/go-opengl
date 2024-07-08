package main

import "github.com/veandco/go-sdl2/sdl"

type MyRect struct {
	Rect  *sdl.Rect
	Color uint32
}

func (r *MyRect) IsOnScreen(screenX, screenY int32) bool {
	if r.Rect.X < 0 {
		return false
	}
	if r.Rect.Y < 0 {
		return false
	}
	if (r.Rect.X + r.Rect.W) > screenX {
		return false
	}
	if (r.Rect.Y + r.Rect.H) > screenY {
		return false
	}
	return true
}
