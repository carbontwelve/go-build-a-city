package gobuildacity

import "github.com/faiface/pixel"

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
	rect := aH.frameSize

	//rect.top = rect.height * animID;
	rect.Min.Y = rect.H() * float64(animID)
	rect.Max.Y = rect.H() * float64(animID)

	//this->bounds = rect;
	aH.bounds = rect




	aH.t = 0.0
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
		// rect := aH.frameSize

		// @todo
	}

	aH.t += dt

	// Adjust for looping

	if aH.t > (duration * float64( aH.animations[aH.currentAnim].GetLength())){
		aH.t = 0.0
	}
	return

}

func NewAnimationHandler() *AnimationHandler {
	aH := AnimationHandler{t: 0.0, currentAnim: -1}
	return &aH
}

func NewAnimationHandlerWithFrameSize(rect pixel.Rect) *AnimationHandler {
	aH := AnimationHandler{frameSize: rect, t: 0.0, currentAnim: -1}
	return &aH
}
