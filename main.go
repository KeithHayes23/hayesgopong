// Hayes 2018
// Born-On Date: 07-26-18
// Version 1.0
// Sets up the Game

package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	// Get the document
	document := js.Global.Get("document")

	// Create the canvas element
	canvas := document.Call("createElement", "canvas")
	canvas.Set("width", 400)
	canvas.Set("height", 300)

	body := document.Get("body")

	// append the canvas element to the html page body
	body.Call("appendChild", canvas)

	// Create the game
	pong := Pong{}

	// Create the Keys Struct to capture the up and down key signals
	keys := NewKeys()

	// This is the main screen
	screen := &CanvasScreen{Canvas: canvas}

	// Initialize the players paddle and the ball
	pong.Load(screen)

	// The loop variable
	var loop func()

	// Called every time frame
	loop = func() {
		pong.Update(screen, keys)
		pong.Draw(screen)
		js.Global.Call("requestAnimationFrame", loop)
	}

	loop()

	// Attach keyboard events
	body.Call("addEventListener", "keydown", func(e *js.Object) {
		keys.OnKeyDown(e.Get("keyCode").Int())
	})

	body.Call("addEventListener", "keyup", func(e *js.Object) {
		keys.OnKeyUp(e.Get("keyCode").Int())
	})

}
