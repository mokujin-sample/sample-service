package service

type ObjectService interface {
	Get(id uint64) (Object, error)
	GetAll(as []uint32, ip []string, limit uint32) ([]Object, int64, error)
}

type ObjectWorker interface {
	ProcessObjects()
	ProcessNotifications()
}
