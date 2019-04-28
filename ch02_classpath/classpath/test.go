package classpath

import (
	"log"
)

type User struct {
	name  string
	email string
}

func (u User) notify() {
	log.Printf("sending User Email to %s<%s>\n", u.name, u.email)
}

func (u *User) notifyPointer() {
	log.Printf("sending User Email to %s<%s>\n", u.name, u.email)
}

func test() {
	u := User{"Bill", "bill@email.com"}
	u.notify()
	u.notifyPointer()
}
