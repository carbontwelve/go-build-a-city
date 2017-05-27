package main

import (
	"github.com/faiface/pixel"
	"fmt"
)

type AnimationHandler struct {
	// Array of animations
	animations []Animation

	// Current time since the animation loop started
	t float64

	currentAnim int

	// Current section of the texture that should be displayed
	bounds pixel.Rect

	// Pixel dimensions of each individual frame
	frameSize pixel.Rect
}

// Add a new animation
func (aH *AnimationHandler) AddAnimation(animation Animation) {
	aH.animations = append(aH.animations, animation)
}

func (aH *AnimationHandler) ChangeAnimation(animID int) {
	if aH.currentAnim == animID || animID >= len(aH.animations) || animID < 0 {
		return
	}

	// Set the current animations
	aH.currentAnim = animID

	// Update the animation bounds

	//sf::IntRect rect = this->frameSize;
	//rect := aH.frameSize

	//rect.top = rect.height * animID;
	//rect.Min.Y = rect.H() * float64(animID)
	//rect.Max.Y = rect.H() * float64(animID)

	//this->bounds = rect;
	aH.bounds = pixel.R(aH.frameSize.Min.X, aH.frameSize.H() * float64(animID), aH.frameSize.Max.X, aH.frameSize.H() * float64(animID) + aH.frameSize.H())
	aH.t = 0.0

	fmt.Println("animID: ", animID, "framesize Max X:", aH.frameSize.Max.X)
}

// Update the current frame of animation. dt is the time since the update
// was last called (i.e. the time for one frame to be executed)
func (aH *AnimationHandler) Update(dt float64) {

	if aH.currentAnim >= len(aH.animations) || aH.currentAnim < 0 {
		return
	}

	duration := aH.animations[aH.currentAnim].Duration

	// Check to see if the animation has progressed to a new frame and if so
	// change to the next frame
	if int((aH.t+dt)/duration) > int(aH.t/duration) {
		// Calculate the frame number
		frame := int((aH.t + dt) / duration)

		// Adjust for looping
		frame %= int(aH.animations[aH.currentAnim].GetLength())

		// Set the sprite to the new frame
		rect := aH.frameSize

		aH.bounds.Min.Y = rect.H() * float64(aH.currentAnim) + rect.H()
		aH.bounds.Min.X = rect.W() * float64(frame)
		aH.bounds.Max.X = rect.W() * float64(frame) + rect.W()
		aH.bounds.Max.Y = rect.H() * float64(aH.currentAnim)

		// fmt.Println("FrameSize W: ", aH.frameSize.W(), "H: ", aH.frameSize.H(), "Min (", aH.frameSize.Min.X, ",", aH.frameSize.Min.Y, ") Max (", aH.frameSize.Max.X, ",", aH.frameSize.Max.Y, ")" )
		// fmt.Println("Frame: ", frame, "W: ", aH.bounds.W(), "H: ", aH.bounds.H(), "Min (", aH.bounds.Min.X, ",", aH.bounds.Min.Y, ") Max (", aH.bounds.Max.X, ",", aH.bounds.Max.Y, ")" )

	}

	aH.t += dt

	// Adjust for looping

	if aH.t > (duration * float64( aH.animations[aH.currentAnim].GetLength())){
		aH.t = 0.0
	}
	return

}

func NewAnimationHandler(rect pixel.Rect) *AnimationHandler {
	aH := AnimationHandler{frameSize: rect, t: 0.0, currentAnim: -1}
	return &aH
}
