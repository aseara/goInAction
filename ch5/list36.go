package main

import "fmt"

type notifier interface {
	notify()
}

type user36 struct {
	name  string
	email string
}

func (u *user36) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name, u.email)
}

type admin struct {
	user36
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending Admin Email To %s<%s>\n",
		a.name, a.email)
}

func main() {
	u := &user36{"Lisa", "lisa@email.com"}
	sendNotify(u)

	a := admin{user36{"Bill", "bill@email.com"}, "1"}
	sendNotify(&a)

	a.notify()
	a.user36.notify()
}

func sendNotify(n notifier) {
	n.notify()
}
