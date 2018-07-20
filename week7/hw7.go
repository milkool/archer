package main

import (
	"archer"
)

type iPhoneＸ struct {
	Name string
	FaceID string
	PowerButton string
}

type iPhone8 struct {
	Name string
	TouchID string
	HomeButton string
}

func (p iPhoneＸ) Sensor() string {
	return p.FaceID
}

func (p iPhoneＸ) Monitor() string {
	return p.Name
}

func (p iPhoneＸ) Button() string {
	return p.PowerButton
}

func (p iPhone8) Sensor() string {
	return p.TouchID
}

func (p iPhone8) Monitor() string {
	return p.Name
}

func (p iPhone8) Button() string {
	return p.HomeButton
}

func main() {
	ix := iPhoneＸ{Name:"Archer", FaceID:"台中金城武的臉", PowerButton:"電源鍵"}
	i8 := iPhone8{Name:"Wesley", TouchID:"台中彭于晏的指紋", HomeButton:"故障的Home鍵"}
	archer.Lock(ix)
	archer.Unlock(ix)
	archer.Lock(i8)
	archer.Unlock(i8)
}