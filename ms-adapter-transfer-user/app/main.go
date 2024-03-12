package main

import (
	"fmt"

	u "ms-adapter-transfer-user/domain/model"
)

func main() {

	//documents := []string{"http://fexample.com/document1", "http://example.com/document2"}
	document := map[string]string{
		"url1: ": "http://fexample.com/document1",
		"url2: ": "http://example.com/document2",
	}

	user := u.User{
		ID:       1102399932,
		Nombre:   "Jose Oliva",
		Email:    "jose.oliva@micarpeta.com",
		Document: u.Document{document}}
	//user := u.MiMetodo("hola")

	fmt.Println(user.Document)
}
