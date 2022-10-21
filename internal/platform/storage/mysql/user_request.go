package mysql

const (
	sqlUserRequestTable = "user_requests"
)

type sqlUserRequest struct {
	ApiKey  string `db:"api_key"`
	Url     string `db:"url"`
	Headers string `db:"headers"`
}
