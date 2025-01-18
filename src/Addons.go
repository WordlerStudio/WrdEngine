package WrdEngine

type Addon interface {
	Start(obj *BaseObject)
	Tick(obj *BaseObject)
}
