package container

type Container interface {
	BuildUseCase(code string) (interface{}, error)

	Get(code string) (interface{}, bool)

	Put(code string, value interface{})
}
