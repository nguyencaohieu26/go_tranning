package main

import "fmt"

//Method in structs
func (u User) printUserEmail() {
	fmt.Println("Your email is: ", u.Email)
}

//Pass the copy of the object into the method -> only change the object's copy email
// func (u User) setNewEmail() {
// 	u.Email = "hieu@gmail.com"
// 	fmt.Println("New email: ", u.Email)
// }

//Pass the reference of the object into method -> change successfully email
func (u *User) setNewEmail() {
	u.Email = "hieu@gmail.com"
	fmt.Println("New email: ", u.Email)
}

func (s Student) getStudentAge() int {
	return s.Age
}

func (s Student) getAverageGrade() float32 {
	var sum float32
	for _, v := range s.Grades {
		sum += float32(v)
	}
	return sum / float32(len(s.Grades))
}

func (a Address) getAddress() {
	fmt.Println("Street", a.Street+" "+a.Province)
}
