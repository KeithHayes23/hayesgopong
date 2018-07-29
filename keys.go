// Hayes 2018
// Born-On Date: 07-26-18
// Version 1.0
// This file sets up the keyboard up and down keys to control the paddle

package main

const (
	keyUp   = 38
	keyDown = 40
)

// Keys Struct
type Keys struct {
	Pressed map[int]bool
}

// Create NewKeys
func NewKeys() Keys {
	keys := Keys{}
	keys.Pressed = make(map[int]bool)
	return keys
}

// IsDown verify if the given key is pressed
func (k *Keys) IsDown(key int) bool {
	stat, exists := k.Pressed[key]
	return exists && stat
}

// OnKeyDown Attached to OnKeyDown event
func (k *Keys) OnKeyDown(key int) {
	k.Pressed[key] = true
}

// OnKeyUp Attached to OnKeyUp event
func (k *Keys) OnKeyUp(key int) {
	k.Pressed[key] = false
}
