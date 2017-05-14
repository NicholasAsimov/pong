package main

import (
	"image"

	"azul3d.org/engine/gfx"
	"azul3d.org/engine/gfx/window"
	"azul3d.org/engine/keyboard"
)

type Vector struct {
	X, Y int
}

type Entity struct {
	Position Vector
	Velocity Vector
	Color    gfx.Color
}

type Player struct {
	Entity
	Width  int
	Height int
}

type Ball struct {
	Entity
	Radius int
}

func (p Player) Bounds() image.Rectangle {
	return image.Rect(
		p.Position.X,
		p.Position.Y,
		p.Position.X+p.Width,
		p.Position.Y+p.Height,
	)
}

func (b Ball) Bounds() image.Rectangle {
	// TODO change rectangle to circle
	return image.Rect(
		b.Position.X,
		b.Position.Y,
		b.Position.X+b.Radius,
		b.Position.Y+b.Radius,
	)
}

func (b *Ball) HandleCollision(windowBounds image.Rectangle) {
	ballBounds := b.Bounds()

	if ballBounds.Min.X < windowBounds.Min.X || ballBounds.Max.X >= windowBounds.Max.X {
		b.Velocity.X = -b.Velocity.X
	}

	if ballBounds.Min.Y < windowBounds.Min.Y || ballBounds.Max.Y >= windowBounds.Max.Y {
		b.Velocity.Y = -b.Velocity.Y
	}
}

func gfxLoop(w window.Window, d gfx.Device) {
	windowWidth, windowHeight := w.Props().Size()

	player1 := Player{
		Entity: Entity{
			Position: Vector{100, 100},
			Velocity: Vector{1, 1},
			Color:    gfx.Color{0, 0, 1, 1},
		},
		Width:  20,
		Height: 50,
	}

	ball := Ball{
		Entity: Entity{
			Position: Vector{windowWidth / 2, windowHeight / 2},
			Velocity: Vector{1, 1},
			Color:    gfx.Color{1, 0, 0, 1},
		},
		Radius: 5,
	}

	for {
		d.Clear(d.Bounds(), gfx.Color{1, 1, 1, 1})

		if w.Keyboard().Down(keyboard.W) {
			player1.Position.Y -= player1.Velocity.Y
		}
		if w.Keyboard().Down(keyboard.S) {
			player1.Position.Y += player1.Velocity.Y
		}

		ball.HandleCollision(d.Bounds())
		ball.Position.X += ball.Velocity.X
		ball.Position.Y += ball.Velocity.Y

		d.Clear(player1.Bounds(), player1.Color)
		d.Clear(ball.Bounds(), ball.Color)

		d.Render()
	}
}

func main() {
	window.Run(gfxLoop, nil)
}
