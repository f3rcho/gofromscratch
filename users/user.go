package users

import (
	"fmt"
	"time"

	"github.com/f3rcho/gofromscratch/models"
)

func HighUser() {
	u := new(models.User)
	u.AddUser(3, "fer", time.Now(), "example@e.com", "123", true)

	fmt.Println(u)
}
