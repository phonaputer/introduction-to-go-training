package itgerr

type Kind int

const (
	KindNotFound Kind = iota
	KindInvalidInput
	KindInternalServer
)
