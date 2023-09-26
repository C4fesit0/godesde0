package users

import (
	"D/Workspace/Go/godesde0/modelos"
	"fmt"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
