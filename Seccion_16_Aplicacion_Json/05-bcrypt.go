package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `clave123`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)
	fmt.Println(string(bs))

	claveLogin := `clave123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(claveLogin))
	if err != nil {
		fmt.Println("No te puedes loguear")
		return
	}

	fmt.Println("Te haz logueado")

}
