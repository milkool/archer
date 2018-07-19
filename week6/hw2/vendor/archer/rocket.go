package archer


type RocketInterface interface {
	Launch ()
}

func Start(r RocketInterface) {
	r.Launch()
}