package dbURL

import "fmt"

func GetDbUrl(postgresHost, postgresUser, postgresPassword, postgresDatabase, postgresPort string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		postgresHost, postgresUser, postgresPassword, postgresDatabase, postgresPort)
}

func main() {
	fmt.Println(GetDbUrl("host", "user", "password", "db", "port"))

}
