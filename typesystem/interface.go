package typesystem

type Speaker interface {
	Speak()
}

type MySpeaker struct{}

func (ms *MySpeaker) Speak() {
	print("now Las is speaking")
}
