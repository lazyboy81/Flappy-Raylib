package entity

type Entity interface {
	Update(dt float32)
	Render()
}
