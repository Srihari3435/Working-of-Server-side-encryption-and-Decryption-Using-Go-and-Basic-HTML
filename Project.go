package main

import (
	"fmt"
	"html/template"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template

func main() {

	tpl, _ = template.ParseGlob("templates/*.html")
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loginauth", loginAuthHandler)
    http.HandleFunc("/description", descriptionHandler)
	http.ListenAndServe("localhost:8000", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****loginHandler running*****")
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func generatePasswordHash(password string) ([]byte, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    return hashedPassword, nil
}

func createUserMap(usernames []string, passwords []string) (map[string][]byte, error) {
    userMap := make(map[string][]byte)
    for i, username := range usernames {
        hashedPassword, err := generatePasswordHash(passwords[i])
        if err != nil {
            return nil, err
        }
        userMap[username] = hashedPassword
    }
    return userMap, nil
}

func checkPassword(username string, password string, userMap map[string][]byte) int {
    hashedPassword, ok := userMap[username]
    fmt.Println("hashed Password " ,string(hashedPassword))
    if !ok {
        return 0
    }
    err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
    if(err == nil){
					return 1
				}
				return 0
}

func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
usernames := []string{"john", "sam", "johny"}
    passwords := []string{"hello", "sam123", "hello"}

    userMap, err := createUserMap(usernames, passwords)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
	fmt.Println("*****loginAuthHandler running*****")
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("username:", username, "password:", password)
	ok := checkPassword(username,password,userMap)
	if ok != 0 {
		tpl.ExecuteTemplate(w,"Welcome.html",nil)
		return
	}
        fmt.Fprintf(w,"Wrong Password Please get Back to Login Page")
	fmt.Println("incorrect password")
}



func descriptionHandler(w http.ResponseWriter, r *http.Request){
fmt.Println("*Desciption Handler running")
 book := r.URL.Query().Get("book")
var description string
    switch book {
    case "One piece":
        description = "Autho : Eiichiro oda . Description : The series focuses on Monkey D. Luffy—a young man made of rubber after unintentionally eating a Devil Fruit—who sets off on a journey from the East Blue Sea to find the deceased King of the Pirates Gol D. Roger's ultimate treasure known as the ONE PIECE, and take over his prior title. In an effort to organize his own crew, the Straw Hat Pirates,[Jp 15] Luffy rescues and befriends a pirate hunter and swordsman named Roronoa Zoro, and they head off in search of the titular treasure. "
    case "JJK":
        description = "Author : Gege Akutami.  . Description : Yuji Itadori is an unnaturally physically strong high school student living in Sendai. On his deathbed in June 2018, his grandfather instills two powerful messages within Yuji: always help others and die surrounded by people. Yuji's friends at the Occult Club attract Curses to their school when they unseal a rotten finger talisman. Yuji swallows the finger to protect Jujutsu Sorcerer Megumi Fushiguro, becoming host to a powerful Curse named Ryomen Sukuna"
    case "Bleach":
        description = "Author : Tite Kubo.  Description : Ichigo Kurosaki is a teenager from Karakura Town who can see ghosts, a talent allowing him to meet Rukia Kuchiki, a Soul Reaper who enters the town in search of a Hollow, a kind of monstrous lost soul who can harm both ghosts and hcd umans. Rukia is one of the Soul Reapers (死神, Shinigami, literally 'Death Gods'), soldiers trusted with ushering the souls of the dead from the World of the Living to the Soul Society (尸魂界ソウル・ソサエティ, lit. Dead Spirit World), the afterlife realm from which she originates and with fighting Hollows. "
    case "sherlock":
        description = "Author :Arthur Conan Doyle . Description : Sherlock Holmes is a collection of short stories by British writer Arthur Conan Doyle, first published on 14 October 1892. It contains the earliest short stories featuring the consulting detective Sherlock Holmes, which had been published in twelve monthly issues of The Strand Magazine from July 1891 to June 1892. "
    case "Crime":
        description = "Author : Fyodor Dostoevsky . Description : Crime and Punishment follows the mental anguish and moral dilemmas of Rodion Raskolnikov, an impoverished ex-student in Saint Petersburg who plans to kill an unscrupulous pawnbroker, an old woman who stores money and valuable objects in her flat."
    case "400 days":
        description = "Author :Chetan Bhagat . Description : 12-year-old Siya has been missing nine months. It’s a cold case, but Keshav wants to help her mother, Alia, who refuses to give up. Welcome to 400 Days―a mystery and romance story like no other.‘My daughter Siya was kidnapped. Nine months ago,’ Alia said.The police had given up. They called it a cold case.  "
    case "Meditations":
        description = "Author : Marcus Aurelius . Description : Marcus teaches that our mind is a thing that controls itself completely and is separated from the world; it cannot be affected by events unless it makes itself be affected. Every appearance is the result of what the mind wills it to appear to be and the mind makes itself exactly what it is."
    case "Atomioc Habbits":
        description = "Author : James Clear . Description : Atomic Habits by James Clear is a comprehensive, practical guide on how to change your habits and get 1% better every day. Using a framework called the Four Laws of Behavior Change, Atomic Habits teaches readers a simple set of rules for creating good habits and breaking bad ones."
    case "Power":
        description = "Author : Robert Green . Description : The 48 Laws of Power by Robert Greene is a self-help book offering advice on how to gain and maintain power, using lessons drawn from parables and the experiences of historical figures. Power depends on the relationships between a person and those he or she seeks to control."
    default:
        description = "No book found"
    }
    fmt.Fprintf(w, description)

}