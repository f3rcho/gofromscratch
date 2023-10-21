package models

import "time"

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Email     string
	Password  string
	Status    bool
}

func (u *User) AddUser(id int, name string, createdAt time.Time, email string, password string, status bool) {
	u.Id = id
	u.Name = name
	u.CreatedAt = createdAt
	u.Email = email
	u.Password = password
	u.Status = status
}
