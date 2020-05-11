package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserDTO struct {
	Email string `json:"email"`
}

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func userAuth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == http.MethodOptions {
		return
	}
	fmt.Printf("Body : %v", r.Body)
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(err)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewEncoder(w).Encode(UserDTO{Email: user.UserName})
}
