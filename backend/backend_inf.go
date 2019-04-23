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

type RegisterEntry struct {
	Name    string
	Servers []string
	Tag     []string
}

type Register interface {
	RegisterSelf(name string, serAddr string, endPoint string)
	GetRegisterGateway() []string
	GetRegisters(name string) *RegisterEntry
	GetAllService() []string
	//0 db //1 services
	RegisterWatch(ty int, change chan string)
}
