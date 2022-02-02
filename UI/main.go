package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//TO USE IN LOCAL COMPUTER
//const conversationURL = "http://localhost:9155/api/v1/conversation"
//const aconversationURL = "http://localhost:9155/api/v1/aconversation"
//const replyURL = "http://localhost:9156/api/v1/reply"

//TO USE IN DOCKER
const conversationURL = "http://conversationapi15:9155/api/v1/conversation"
const aconversationURL = "http://conversationapi15:9155/api/v1/aconversation"
const replyURL = "http://repliesapi15:9156/api/v1/reply"

const studentURL = "http://10.31.11.12:9211/api/v1/students"
const tutorURL = "http://10.31.11.12:9181/api/v1/tutor/GetAllTutor"

//const studentURL = "http://localhost:9155/api/v1/students"

type Conversation struct {
	ID               int
	InitiatorID      int
	RecipientID      int
	StartTime        time.Time
	NoofMessages     int
	ConversationName string
}

type Replies struct {
	ID             int
	ConversationID int
	SenderID       int
	ReceiverID     int
	TimeSent       time.Time
	TimeRead       time.Time
	Header         string
	Content        string
}

type aConversation struct {
	Convo   Conversation
	Replies []Replies
}

type User struct {
	Name   string
	UserID string
}

type User1 struct {
	student_id    string
	name          string
	date_of_birth string
	address       string
	phone_number  string
}

type Student struct {
	Student_id    string
	Name          string
	Date_of_birth string
	Address       string
	Phone_number  string
}

type Tutor struct {
	Deleted      string
	TutorID      int
	Firstname    string
	Lastname     string
	Email        string
	Descriptions string
}

//RETRIEVE ALL Students
func GetUsers() []Student {
	url := studentURL
	fmt.Println("URL: ", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil
	} else {
		fmt.Println("PASS 1")
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		//var trip Trip
		var users []Student
		fmt.Println("PASS 2")
		err := json.Unmarshal([]byte(data), &users)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(users)
		fmt.Println("PASS 3")
		response.Body.Close()
		return users
	}
}

//RETRIEVE ALL Tutors
func GetTutor() []Tutor {
	url := tutorURL
	fmt.Println("URL: ", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil
	} else {
		fmt.Println("PASS 1")
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		//var trip Trip
		var users []Tutor
		fmt.Println("PASS 2")
		err := json.Unmarshal([]byte(data), &users)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(users)
		fmt.Println("PASS 3")
		response.Body.Close()
		return users
	}
}

//GET Conversation of Code ID
func GetConversation(code string) Conversation {
	url := replyURL
	fmt.Println("CODE: ", code)
	if code != "" {
		url = conversationURL + "/" + code
	}
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var driver Conversation
		err := json.Unmarshal([]byte(data), &driver)
		if err != nil {
			fmt.Println("ERROR1: ", err)
		}
		fmt.Println("ID            :", driver.ID)
		fmt.Println("InitiatorID    :", driver.InitiatorID)
		fmt.Println("RecipientID     :", driver.RecipientID)
		fmt.Println("StartTime :", driver.StartTime)
		fmt.Println("NoofMessages :", driver.NoofMessages)
		fmt.Println("ConversationName     :", driver.ConversationName)

		response.Body.Close()
		return driver
	}
	return Conversation{}
}

//Get Conversation and Replies
func GetaConversation(code string) (*aConversation, error) {
	url := replyURL
	if code != "" {
		url = aconversationURL + "/" + code
		fmt.Println(url)
	}
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var aconvo aConversation
		err := json.Unmarshal([]byte(data), &aconvo)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("GET ACONVERSATION")
		fmt.Println(aconvo)
		fmt.Println("ID            :", aconvo.Convo.ID)
		fmt.Println("InitiatorID    :", aconvo.Convo.InitiatorID)
		fmt.Println("RecipientID     :", aconvo.Convo.RecipientID)
		fmt.Println("StartTime :", aconvo.Convo.StartTime)
		fmt.Println("NoofMessages :", aconvo.Convo.NoofMessages)
		fmt.Println("ConversationName     :", aconvo.Convo.ConversationName)

		response.Body.Close()
		return &aconvo, nil
	}
}

//RETRIEVE ALL CONVERSATIONS of USER X
func GetConversationOfUser(ID string) (*[]Conversation, error) {
	url := conversationURL
	url = conversationURL + "/user/" + ID
	fmt.Println("URL: ", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err //return error
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		//var trip Trip
		var trips []Conversation
		err := json.Unmarshal([]byte(data), &trips)
		//fmt.Println(trips)
		fmt.Println(err)
		response.Body.Close()
		a := len(trips)
		for i := 0; i < a; i++ {
			fmt.Println("Trip-----------------------------")
			fmt.Println("ID           :", trips[i].ID)
			fmt.Println("InitiatorID      :", trips[i].InitiatorID)
			fmt.Println("RecipientID     :", trips[i].RecipientID)
			fmt.Println("StartTime:", trips[i].StartTime)
			fmt.Println("NoofMessages    :", trips[i].NoofMessages)
			fmt.Println("ConversationName  :", trips[i].ConversationName)
		}
		return &trips, nil
	}
}

//RETRIEVE ALL CONVERSATIONS
func GetConversations() {
	url := conversationURL
	url = conversationURL + "/"
	fmt.Println("URL: ", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		fmt.Println("PASS 1")
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		//var trip Trip
		var trips []Conversation
		fmt.Println("PASS 2")
		err := json.Unmarshal([]byte(data), &trips)
		//fmt.Println(trips)
		fmt.Println(err)
		fmt.Println("PASS 3")
		response.Body.Close()
		a := len(trips)
		for i := 0; i < a; i++ {
			fmt.Println("Trip-----------------------------")
			fmt.Println("ID           :", trips[i].ID)
			fmt.Println("InitiatorID      :", trips[i].InitiatorID)
			fmt.Println("RecipientID     :", trips[i].RecipientID)
			fmt.Println("StartTime:", trips[i].StartTime)
			fmt.Println("NoofMessages    :", trips[i].NoofMessages)
			fmt.Println("ConversationName  :", trips[i].ConversationName)
		}
	}

}

//POST
func addConversation(code string, p Conversation) {

	jsonValue, _ := json.Marshal(p)
	//response, err := http.Post(conversationURL+"/"+code+"?key="+key,
	//	"application/json", bytes.NewBuffer(jsonValue))
	response, err := http.Post(conversationURL+"/"+code,
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

//PUT
func updateConversation(code string, d Conversation) {
	jsonValue, _ := json.Marshal(d)

	request, err := http.NewRequest(http.MethodPut,
		conversationURL+"/"+code,
		bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

//DELETE
func DeleteConversation(code string) {
	request, err := http.NewRequest(http.MethodDelete,
		conversationURL+"/"+code, nil)
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

//GET
func GetReply(code string) Replies {
	url := replyURL
	if code != "" {
		url = replyURL + "/" + code
	}
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var driver Replies
		err := json.Unmarshal([]byte(data), &driver)
		if err != nil {
			fmt.Println("ERROR1: ", err)
		}
		fmt.Println("ID            :", driver.ID)
		fmt.Println("ConversationID    :", driver.ConversationID)
		fmt.Println("SenderID     :", driver.SenderID)
		fmt.Println("ReceiverID :", driver.ReceiverID)
		fmt.Println("TimeSent :", driver.TimeSent)
		fmt.Println("Header     :", driver.Header)
		fmt.Println("Content     :", driver.Content)
		fmt.Println("PASS 3")
		response.Body.Close()
		return driver
	}
	return Replies{}
}

//POST
func addReply(code string, p Replies) {
	fmt.Println("excalibur3")
	fmt.Println(p)
	jsonValue, _ := json.Marshal(p)
	//response, err := http.Post(conversationURL+"/"+code+"?key="+key,
	//	"application/json", bytes.NewBuffer(jsonValue))
	response, err := http.Post(replyURL+"/"+code,
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

//PUT
func updateReply(code string, d Replies) {
	jsonValue, _ := json.Marshal(d)

	request, err := http.NewRequest(http.MethodPut,
		replyURL+"/"+code,
		bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

//DELETE
func DeleteReply(code string) {
	request, err := http.NewRequest(http.MethodDelete,
		replyURL+"/"+code, nil)
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

type Data struct {
	DataFields []Conversation
	ID         string
}

type Data2 struct {
	DataFields aConversation
	ID         string
}

type Data3 struct {
	DataFields Conversation
	ID         string
}

type Data4 struct {
	DataFields Replies
	ConvoID    string
	UserID     string
}

type Data5 struct {
	Users  []User
	UserID string
}

type Data6 struct {
	Users  []User1
	UserID string
}

//backup for data9
type Data7 struct {
	Users  []Student
	UserID string
}

type Data9 struct {
	Users  []Student
	Tutors []Tutor
	UserID string
}

var templates = template.Must(template.ParseFiles("template/view.html", "template/home.html", "template/chat.html", "template/edit.html", "template/edit1.html", "template/write.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p Data) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/view/"):]
	p, _ := GetConversationOfUser(id)
	d := Data{
		DataFields: *p,
		ID:         id,
	}
	renderTemplate(w, "view", d)
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/chat/"):]
	fmt.Println(id)
	split := strings.Split(id, "/")
	fmt.Println(split)
	c, _ := GetaConversation(split[1])
	d := Data2{
		DataFields: *c,
		ID:         split[0],
	}
	err := templates.ExecuteTemplate(w, "chat.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/save/"):]
	fmt.Println(id)
	split := strings.Split(id, "/")
	fmt.Println(split)

	body := r.FormValue("body")
	header := r.FormValue("header")
	p := GetReply(split[2])
	p.Content = body
	p.Header = header
	updateReply(split[2], p)
	http.Redirect(w, r, "/chat/"+split[0]+"/"+split[1], http.StatusFound)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/create/"):]
	fmt.Println("PASS CREATE HANDLER: ID:", id)
	fmt.Println(id)
	split := strings.Split(id, "/")
	fmt.Println(split)

	body := r.FormValue("body")
	header := r.FormValue("header")
	ConvoID, _ := strconv.ParseInt(split[1], 10, 64)
	convo := GetConversation(split[1])
	a := convo.InitiatorID
	b := convo.RecipientID
	UserID, _ := strconv.ParseInt(split[0], 10, 64)

	if int(UserID) == a { //FIND SENDER OF THIS MESSAGE IF IS THE INITIATOR OR RECIPIENT
		addReply("1", Replies{Content: body, Header: header, ConversationID: int(ConvoID), ReceiverID: b, SenderID: int(UserID)}) //need sender and receiver
	} else {
		addReply("1", Replies{Content: body, Header: header, ConversationID: int(ConvoID), ReceiverID: a, SenderID: int(UserID)}) //need sender and receiver
	}
	//addReply("1", Replies{Content: body, Header: header, ConversationID: int(ConvoID), ReceiverID: receiver, SenderID: int(UserID)}) //need sender and receiver
	//p := GetReply(id)
	//p.Content = body
	//updateReply(id, p)
	c := GetConversation(split[1]) //RETRIEVE CONVERSATION TO UPDATE NUMBER OF MESSAGES
	c.NoofMessages = c.NoofMessages + 1
	updateConversation(split[1], c)
	http.Redirect(w, r, "/chat/"+id, http.StatusFound)
}

/*
func create1Handler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/write/"):]

	fmt.Println(id)
	split := strings.Split(id, "/")
	fmt.Println(split)

	fmt.Println("Conversation ID: ", split[1])
	c := GetConversation(split[1])
	fmt.Println(c)

	d := Data3{
		DataFields: c,
		ID:         split[0],
	}

	err := templates.ExecuteTemplate(w, "create.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}*/

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/delete/"):]
	DeleteReply(id)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func deleteHandler1(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/deleteconvo/"):]
	DeleteConversation(id)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/edit/"):]

	split := strings.Split(id, "/")

	fmt.Println("Reply ID: ", split[2])
	p := GetReply(split[2])
	d := Data4{
		DataFields: p,
		ConvoID:    split[1],
		UserID:     split[0],
	}

	err := templates.ExecuteTemplate(w, "edit.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func editHandler1(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/editconvo/"):]

	split := strings.Split(id, "/")

	fmt.Println("Convo ID: ", split[1])
	p := GetConversation(split[1])
	d := Data3{
		DataFields: p,
		ID:         split[0],
	}
	err := templates.ExecuteTemplate(w, "edit1.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler1(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/saveconvo/"):]
	fmt.Println(id)
	split := strings.Split(id, "/")
	fmt.Println(split)

	body := r.FormValue("body")
	p := GetConversation(split[1])
	p.ConversationName = body
	updateConversation(split[1], p)
	http.Redirect(w, r, "/view/"+split[0], http.StatusFound)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/writeconvo/"):]

	fmt.Println(id)
	var users1 = GetUsers()
	if len(users1) == 0 {
		users1 = []Student{
			{Student_id: "1", Name: "Wai Hou Man", Date_of_birth: "996076800000", Address: "BLK678B Jurong West, Singapore", Phone_number: "6511111111"},
			{Student_id: "2", Name: "Zachary Hong Rui Quan", Date_of_birth: "1007136000000", Address: "BLK123F Orchard Rd", Phone_number: "6512345678"},
			{Student_id: "3", Name: "Data rn is hard coded as", Date_of_birth: "912441600000", Address: "BLK666A Punggol", Phone_number: "6533333333"},
			{Student_id: "4", Name: "Student API Link Down", Date_of_birth: "912441600000", Address: "BLK666A Punggol", Phone_number: "6533333333"},
		}
	}
	var tutors = GetTutor()
	if len(tutors) == 0 {
		tutors = []Tutor{
			{Deleted: "", TutorID: 1, Firstname: "Tutor", Lastname: "1", Email: "email", Descriptions: "Relief Teacher"},
			{Deleted: "", TutorID: 2, Firstname: "Tutor", Lastname: "2", Email: "email", Descriptions: "Relief Teacher"},
		}
	}
	tutors = changetutorid(tutors)
	/*p2 := Data7{
		UserID: id,
		Users:  users1,
	}*/
	p2 := Data9{
		UserID: id,
		Users:  users1,
		Tutors: tutors,
	}
	fmt.Println(p2)

	fmt.Println("User ID: ", id)
	err := templates.ExecuteTemplate(w, "write.html", p2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Data8 struct {
	Users  Student
	UserID string
}

func changetutorid(u []Tutor) []Tutor { //GOLANG CANNOT PASS ATTRIBUTES THAT START WITH LOWERCASE LETTER TO SHOW IN HTML
	u2 := []Tutor{}
	for _, v := range u {
		var temp = Tutor{
			Deleted:      v.Deleted,
			TutorID:      v.TutorID + 100,
			Firstname:    v.Firstname,
			Lastname:     v.Lastname,
			Email:        v.Email,
			Descriptions: v.Descriptions,
		}
		//fmt.Println(temp)
		u2 = append(u2, temp)
	}
	return u2
}

func createconvoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/createconvo/"):]
	fmt.Println(id)

	convoname := r.FormValue("body")
	//RID := r.FormValue("recipientid")
	RID := r.FormValue("id")
	fmt.Println("RID: " + RID)

	RID1, err := strconv.ParseInt(RID, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	UserID, _ := strconv.ParseInt(id, 10, 64)
	fmt.Println(RID1)
	p := Conversation{
		InitiatorID:      int(UserID),
		RecipientID:      int(RID1),
		NoofMessages:     0,
		ConversationName: convoname,
	}
	fmt.Println("CHECK")
	fmt.Println(p)
	addConversation("1", p)
	http.Redirect(w, r, "/view/"+id, http.StatusFound)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	var users1 = GetUsers()
	if len(users1) == 0 {
		users1 = []Student{
			{Student_id: "1", Name: "Wai Hou Man", Date_of_birth: "996076800000", Address: "BLK678B Jurong West, Singapore", Phone_number: "6511111111"},
			{Student_id: "2", Name: "Zachary Hong Rui Quan", Date_of_birth: "1007136000000", Address: "BLK123F Orchard Rd", Phone_number: "6512345678"},
			{Student_id: "3", Name: "Data rn is hard coded as", Date_of_birth: "912441600000", Address: "BLK666A Punggol", Phone_number: "6533333333"},
			{Student_id: "4", Name: "Student API Link Down", Date_of_birth: "912441600000", Address: "BLK666A Punggol", Phone_number: "6533333333"},
		}
	}
	var tutors = GetTutor()
	if len(tutors) == 0 {
		tutors = []Tutor{
			{Deleted: "", TutorID: 1, Firstname: "Tutor", Lastname: "1", Email: "email", Descriptions: "Relief Teacher"},
			{Deleted: "", TutorID: 2, Firstname: "Tutor", Lastname: "2", Email: "email", Descriptions: "Relief Teacher"},
		}
	}
	tutors = changetutorid(tutors)
	/*p2 := Data7{
		UserID: id,
		Users:  users1,
	}*/
	p2 := Data9{
		UserID: "0",
		Users:  users1,
		Tutors: tutors,
	}

	fmt.Println(p2)
	//fmt.Println("User ID: ", 0)
	err := templates.ExecuteTemplate(w, "home.html", p2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//template := template.Must(template.ParseFiles("template/home.html"))
	//if err := template.ExecuteTemplate(w, "home.html", nil); err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//	}
}

func removespc(str string) string { //ALLOW ADDING OF < ' > TO DATABASE WITHOUT ERROR
	if str != "" {
		i := strings.Index(str, "'")
		if i > -1 {
			first := str[:i]
			second := str[i+1:]
			t := first + "\\'" + removespc(second)
			return t
		}
		if i == -1 {
			return str
		}
	}
	return ""
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/chat/", chatHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/create/", createHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.HandleFunc("/deleteconvo/", deleteHandler1)
	http.HandleFunc("/editconvo/", editHandler1)
	http.HandleFunc("/saveconvo/", saveHandler1)
	http.HandleFunc("/writeconvo/", writeHandler)
	http.HandleFunc("/createconvo/", createconvoHandler)
	http.ListenAndServe(":8000", nil)
}

/*

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TEST HANDLER")
	template := template.Must(template.ParseFiles("template/view.html"))
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//conversation := GetConversation("2")

	conversations := GetConversationOfUser("101")
	data := Data{DataFields: conversations}
	if err := template.ExecuteTemplate(w, "view.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := template.ExecuteTemplate(w, "view.html", conversation); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Redirect Called")
	http.Redirect(w, r, "/home", 301)
}

func chatHandler1(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/chat/"):]
	//GetConversationOfUser(id)
	c := GetConversation(id)
	//p, _ := loadPage(title) //Get Data
	x, _ := strconv.ParseInt(id, 10, 64)
	fmt.Fprintf(w, "<h1>%s</h1><div>ID: %d</div>", c.ConversationName, x)
}*/
