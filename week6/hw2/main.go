package main

import (
	archer "archer"
)

func main()  {
	rocket1 := archer.Rocket{Name:"洽洽號", Fuel:3}
	rocket2 := archer.Rocket{Name:"粉紅惡魔兔", Fuel:5}
	rocket1.Launch(rocket1)
	rocket2.Launch(rocket2)
}
