package env

import "os"

var MONGO_PORT, MONGO_HOST, MONGO_USER, MONGO_PASS, MONGO_DB, MONGO_URL string
func init() {
	MONGO_PORT = os.Getenv("MONGO_PORT")
	MONGO_HOST = os.Getenv("MONGO_HOST")
	MONGO_USER = os.Getenv("MONGO_USER")
	MONGO_PASS = os.Getenv("MONGO_PASS")
	MONGO_DB   = mongoDB()
	MONGO_URL  = os.Getenv("MONGO_URL")
}

func mongoDB() string {
	db := os.Getenv("MONGO_DB")
	if db != "" {
		return db
	}

	return "eucaprice"
}
