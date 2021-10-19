package service

type Object struct {
	ID      uint64 `json:"id,omitempty"`
	IP      string `json:"ip,omitempty"`
	Meta    string `json:"meta"`
	HasMeta uint8  `json:"has_meta"`
}

type Repository interface {
	Create(a Object) (uint64, error)
	Update(a Object) error
	Get(id uint64) (Object, error)
	GetAll(as []uint32, ip []string, limit uint32) ([]Object, int64, error)
	GetTargets(id uint64) (list []map[string]interface{}, count int64, err error)
	UnderObject(as []uint32, ip []string) (int64, error)
}
