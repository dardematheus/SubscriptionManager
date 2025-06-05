package handlers

import (
	"database/sql"
)

type Env struct {
	DB *sql.DB
	//Session e essas coisa
}
