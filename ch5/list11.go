package main

import (
	"fmt"
	"os"
)

type user11 struct {
	name  string
	email string
}

func (u user11) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name, u.email)
}

func (u *user11) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user11{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user11{"Lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("bill@nodomain.com")
	bill.notify()

	lisa.changeEmail("lisa@nodomain.com")
	lisa.notify()

	_, _ = os.Open("D:\\test.file")

}
