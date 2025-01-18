package WrdEngine

type Object interface {
	Tick()
	Render() error
	Connect(event Event, handler func())
	Attach(addon Addon)
	ChangePos(x, y int32)
	Emit(event Event)
}
