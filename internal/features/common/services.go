package common

type IdProvider interface {
	GetId() (string, error)
}
