package backend

type Backend interface {
	GetDB() DB
	GetRegister() Register
}

type DB interface {
	Put(key string, data []byte) error
	Get(key string) ([]byte, error)
	Del(key string) error
}

type Register interface {
	RegisterSelf(regAddr string, serAddr string)
	GetRegisterGateway() []string
	GetRegisters(name string) []string
	RegisterWatch()
}
