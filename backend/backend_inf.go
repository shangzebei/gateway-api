package backend

type DB interface {
	Put(key string, data []byte) error

	Get(key string) ([]byte, error)

	Del(key string) error
}

type Register interface {
	RegisterSelf()
	RegisterWatch()
}
