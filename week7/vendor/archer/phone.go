package archer

type phone interface {
	Sensor() string
	Monitor() string
	Button() string
}

func Unlock(p phone) {
	println("已使用" + p.Sensor() + "解鎖")
	println("Hi~ " + p.Monitor())
}

func Lock(p phone) {
	println("已用" + p.Button() + "上鎖")
}