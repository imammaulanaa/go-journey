package handler

import (
	"fmt"
	"net/http"
	"middleware-context/middleware"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUserFromContext(r.Context())
	fmt.Fprintf(w, "Hello, %s! Welcome to the users page.\n", user)
}
