package WrdEngine

type Screen struct {
	Width, Height int32
	// TODO...
}

var screen *Screen

type Physic struct {
	GravtyPower float64
	// TODO...
}

var physic Physic = Physic{
	GravtyPower: 10,
}
