package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const WINDOW_W = 1280
const WINDOW_H = 720

const RECT_W = 200
const RECT_H = 200

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Options bitmask
	windowopts := (uint32)(sdl.WINDOW_SHOWN)
	window, err := sdl.CreateWindow("Balls", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		WINDOW_W, WINDOW_H, windowopts)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{X: 0, Y: 0, W: RECT_W, H: RECT_H}
	rect2 := sdl.Rect{X: WINDOW_W - RECT_W, Y: WINDOW_H - RECT_H, W: RECT_W, H: RECT_H}
	myrect := MyRect{Rect: &rect}
	myrect2 := MyRect{Rect: &rect2}

	colour := sdl.Color{R: 0, G: 0, B: 0, A: 255}
	black := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)

	window.UpdateSurface()

	red := sdl.MapRGB(surface.Format, 255, 0, 0)
	green := sdl.MapRGB(surface.Format, 0, 255, 0)
	blue := sdl.MapRGB(surface.Format, 0, 0, 255)
	intersectColor := sdl.MapRGB(surface.Format, 255, 255, 0)

	surface.FillRect(myrect.Rect, red)
	surface.FillRect(myrect2.Rect, green)

	myrect.Color = red
	running := true
	move := false
	for running {

		if !myrect.IsOnScreen(1280, 720) {
			myrect.Color = blue
		} else {
			myrect.Color = red
		}
		if !myrect2.IsOnScreen(1280, 720) {
			myrect2.Color = blue
		} else {
			myrect2.Color = green
		}

		surface.FillRect(myrect.Rect, myrect.Color)
		surface.FillRect(myrect2.Rect, myrect2.Color)

		interRect, isIntersect := myrect.Rect.Intersect(myrect2.Rect)
		if isIntersect {
			surface.FillRect(&interRect, intersectColor)
		}
		window.UpdateSurface()
		time.Sleep(time.Millisecond * 5)

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch ev := event.(type) {
			case *sdl.QuitEvent:
				println("Exiting")
				running = false
				break
			case *sdl.KeyboardEvent:
				println(ev.Keysym.Sym, ": ", sdl.K_q)
				if event.GetType() == sdl.KEYDOWN && ev.Keysym.Sym == sdl.K_q {
					println("Exiting")
					running = false
					break
				}
			case *sdl.MouseButtonEvent:
				if event.GetType() == sdl.MOUSEBUTTONDOWN {
					move = !move
				}
			case *sdl.MouseMotionEvent:
				if !move {
					continue
				}
				surface.FillRect(myrect.Rect, black)
				surface.FillRect(myrect2.Rect, black)
				myrect.Rect.X += ev.XRel
				myrect.Rect.Y += ev.YRel
				myrect2.Rect.X -= ev.XRel
				myrect2.Rect.Y -= ev.YRel
			}
		}
	}
}
