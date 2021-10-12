package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                 string
	Age                  uint16
	Money                int16
	Avg_grades, Happynes float64
	Hobbies              []string
}

func (u User) getAllinfo() string {
	return fmt.Sprintf("User name is: %s. He is %d "+
		"and he has money equal: %d", u.Name, u.Age, u.Money)
}
func (u *User) setNewName(newName string) {
	u.Name = newName
}
func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{
		Name:       "Bob",
		Age:        25,
		Money:      12,
		Avg_grades: 4.3,
		Happynes:   1,
		Hobbies:    []string{"Cyber Sport", "Footbal", "Eat"},
	}
	//fmt.Fprintf(w," " + bob.getAllinfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	/*fmt.Fprintf(w, `<h1>main text</h1>
	<b>main text</b>`)*/
	tmpl.Execute(w, bob)

}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts Page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()

}
