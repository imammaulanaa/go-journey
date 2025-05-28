package repository

import "fmt"

func GetUser(id int) string {
	return fmt.Sprintf("User-%d", id)
}
