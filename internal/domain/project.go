package domain

type Project struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	TypeID    string      `json:"type_id"`
	Type      ProjectType `json:"type"`
	Users     []User      `json:"users"`
	CreatedAt string      `json:"created_at"`
	UpdateAt  string      `json:"update_at"`
}
