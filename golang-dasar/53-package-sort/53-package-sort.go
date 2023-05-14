package main

import (
	"fmt"
	"sort"
)

/*
	Package sort adalah package yang berisikan utilitas untuk proses pengurutan
	Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interfac
*/

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []User{
		{Name: "Ilham", Age: 12},
		{Name: "Abyan", Age: 13},
		{Name: "Farhan", Age: 14},
		{Name: "Ramy", Age: 15},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
