package objects

type ObjectResponse struct {
	ID      uint64 `db:"id" json:"id,omitempty"`
	IP      string `db:"ip" json:"ip,omitempty"`
	HasMeta uint8  `db:"has_meta" json:"has_meta,omitempty"`
}
