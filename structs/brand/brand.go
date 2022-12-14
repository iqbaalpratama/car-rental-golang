package brand

import "database/sql"

type Brand struct {
	ID        int
	Name      string `json:"name"`
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
