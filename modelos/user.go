package modelos

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	status    bool
}

func (this *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	this.Id = id
	this.Name = name
	this.status = status
	this.CreatedAt = createdAt
}
