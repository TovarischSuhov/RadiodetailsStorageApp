package domain

type Object interface {
	GetType() string
	GetDescription() string
}
