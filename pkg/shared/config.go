package shared

type Handler func(args any) error

type Config struct {
	Url        string
	HandlerMap map[string]Handler
}
