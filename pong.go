// Hayes 2018
// Born-On Date: 07-26-18
// Version 1.0
// Create te games pieces like screen, paddle, and Ball

package main

import "strconv"
import "fmt"
import "github.com/gen2brain/beeep"
import "math/rand"
import "time"

const (
	gname = "Hayes Pong 2018"
	pspeed = 2
	bspeed = 2
	statsHeight = 20
)

// Paddle
type Paddle struct {
	x      int
	y      int
	w      int
	h      int
	vel    float32
	hits int
	misses int
}

// Paddle Up
func (p *Paddle) Up() {
	p.y -= pspeed
}

// Paddle Down
func (p *Paddle) Down() {
	p.y += pspeed
}

// The Ball
type Ball struct {
	x     int
	y     int
	w     int
	h     int
	xVel  int
	yVel  int
	speed int
}

// The game object
// This stores the scores, players, ball
type Pong struct {
	player   *Paddle
	ball     *Ball
}

// Load the game
// Create the necessary objects
func (p *Pong) Load(s Screen) {
	p.player = &Paddle{
		x:      10,
		w:      10,
		h:      50,
		vel:    1,
		hits: 	0,
		misses: 0,
	}

	p.player.y = s.Height()/2 - p.player.h/2

	p.ball = &Ball{
		w:     10,
		h:     10,
		xVel:  1,
		yVel:  -1,
		speed: bspeed,
	}

	p.reset(s)
}

// Reserves the ball after a miss
// Its randomly positioned between half to the right and full right and somewhere from top to bottom
func (p *Pong) reset(s Screen) {
	p.ball.x = p.random(s.Width()/2, s.Width() - p.ball.w)
	p.ball.y = p.random(p.ball.h, s.Height() - p.ball.h)
	fmt.Println(fmt.Sprintf("x:%d y: %d", p.ball.x, p.ball.y ))
}

// Beeps when ball reaches top, back, and bottom walls or the players paddle
func (p *Pong) beep() {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}
}

// Generate Randum Number
// This is much more randomized using nana seconds
func (p *Pong) random(min, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min) + min
}

// Update, Check for collisions and update the game state
// this is
func (p *Pong) Update(s Screen, keys Keys) {
	// Update Ball position
	p.ball.x = p.ball.x + p.ball.speed*p.ball.xVel
	p.ball.y = p.ball.y + p.ball.speed*p.ball.yVel

	// Check For Collisions
	// Collisions occur when the ball hits the top,bottom,back wall, and player Paddle

	// Bottom Wall
	if (p.ball.y + p.ball.h) >= s.Height() {
		p.ball.yVel = -1
		p.beep()
		fmt.Println("hit bottom")
	}

	// Back Wall
	if (p.ball.x + p.ball.w) >= s.Width() {
		p.ball.xVel = -1
		p.beep()
		fmt.Println("hit wall")
	}

	// Top Wall
	if p.ball.y <= statsHeight {
		p.ball.yVel = 1
		p.beep()
		fmt.Println("hit top")
	}

	// Player Paddle
	if p.ball.x <= (p.player.x+p.player.w) &&
		(p.ball.y+p.ball.h/2) >= p.player.y &&
		p.ball.y <= (p.player.y+p.player.h) {
		p.player.hits++
		p.ball.xVel = 1
		p.beep()
		fmt.Println("hit paddle")
	}

	// Missed Ball
	// Ball crosses the front wall
	if p.ball.x <= 0 {
		p.player.misses++
		fmt.Println("missed ball")
		p.reset(s)
	}

	// Keyboard Detection
	if keys.IsDown(keyUp) {
		if p.player.y > 22 {
			p.player.Up()
		}
	}

	if keys.IsDown(keyDown) {
		if p.player.y <= s.Height()-p.player.h - 2 {
			p.player.Down()
		}
	}
}

// Draw the objects on the screen
func (p *Pong) Draw(screen Screen) {

	//black := Color{0, 0, 0}
	white := Color{255, 255, 255}
	green := Color{0, 200, 0}
	blue := Color{0, 0, 255}
	yellow := Color{255, 255, 0}
	black := Color{0, 0, 0}
	//red := Color{255, 0, 0}

	// Stats Board
	screen.DrawRect(0, 0, screen.Width(), statsHeight, black)

	// Main Board
	screen.DrawRect(0, statsHeight, screen.Width(), screen.Height()-statsHeight, blue)

	// Player
	screen.DrawRect(p.player.x, p.player.y, p.player.w, p.player.h, white)

	// Ball
	screen.DrawRect(p.ball.x, p.ball.y, p.ball.w, p.ball.h, yellow)

	// Game Name
	screen.DrawText(gname, 15, 15, white)

	// Hits
	hits := fmt.Sprintf("Hits: %s",strconv.Itoa(p.player.hits))
	screen.DrawText(hits, screen.Width()/2-50, 15, green)

	// Misses
	misses := fmt.Sprintf("Misses: %s",strconv.Itoa(p.player.misses))
	screen.DrawText(misses, screen.Width()/2+20, 15, green)

	// Ball Telemetry
	ball := fmt.Sprintf("Ball %d:%d",p.ball.x, p.ball.y)
	screen.DrawText(ball, screen.Width()-80, 15, yellow)
}
