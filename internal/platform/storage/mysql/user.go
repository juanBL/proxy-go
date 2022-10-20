package mysql

const (
	sqlUserTable = "users"
)

type sqlUser struct {
	ApiKey         string `db:"api_key"`
	ExpirationDate string `db:"expiration_date"`
}
