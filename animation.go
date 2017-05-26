package main

type Animation struct {
	StartFrame uint
	EndFrame   uint
	Duration   float64
}

func (a Animation) GetLength() (uint) {
	return a.EndFrame - a.StartFrame + 1
}

func NewAnimation(startFrame uint, endFrame uint, duration float64) *Animation {
	a := Animation{StartFrame:startFrame, EndFrame: endFrame, Duration: duration}
	return &a
}
