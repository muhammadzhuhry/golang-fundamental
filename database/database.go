package database

var connection string

func init() {
	connection = "MySQL"
}

func getDatabase() string {
	return connection
}
