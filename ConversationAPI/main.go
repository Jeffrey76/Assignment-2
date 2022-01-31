package main

import (
	//"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"strconv"

	"github.com/blockloop/scan"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"io/ioutil"
	"net/http"
)

//TO USE IN DOCKER
//const replyURL = "http://repliesapi15:9156/api/v1/reply"
//const repliesURL = "http://repliesapi15:9156/api/v1/replies"

//TO USE IN LOCAL COMPUTER + CHANGE DB URL TO LOCALHOST:3306 INSTEAD OF DATABASE15
const replyURL = "http://localhost:9156/api/v1/reply"
const repliesURL = "http://localhost:9156/api/v1/replies"

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

func validKey(r *http.Request) bool {
	v := r.URL.Query()
	if key, ok := v["key"]; ok {
		if key[0] == "2c78afaf-97da-4816-bbee-9ad239abb297" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!")
}

//Provide Valid ID
func GetIDConversation(db *sql.DB) int {
	query := fmt.Sprintf("SELECT COUNT(*) FROM Conversation")

	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	//fmt.Println(results)
	for results.Next() {
		// map this type to the record in the table
		var count = 0
		err = results.Scan(&count)
		fmt.Println(count)
		return count + 1
	}
	return 0
}

func GetConversation(db *sql.DB, ID int) Conversation {
	var query string
	//query = fmt.Sprintf("SELECT * FROM Conversation WHERE InitiatorID='%d' OR RecipientID='%d' ORDER BY ID DESC;", ID, ID)
	query = fmt.Sprintf("SELECT * FROM Conversation WHERE ID='%d' ORDER BY ID DESC;", ID)

	fmt.Println(query)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		// map this type to the record in the table
		var c Conversation
		err = results.Scan(&c.ID, &c.InitiatorID,
			&c.RecipientID, &c.StartTime, &c.NoofMessages, &c.ConversationName)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(c.ID, c.InitiatorID,
			c.RecipientID, c.StartTime, c.NoofMessages, c.ConversationName)
		return c
	}
	return Conversation{}
}

func GetAllConversation(db *sql.DB) []Conversation {
	var query string
	query = fmt.Sprintf("SELECT * FROM Conversation ORDER BY ID DESC;")

	fmt.Println(query)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	var conversations []Conversation
	scan.Rows(&conversations, results)
	fmt.Println(conversations)
	return conversations
}

//CURRENTLY UNUSED
func GetAllConversationOfUser(db *sql.DB, ID int) []Conversation {
	var query string
	query = fmt.Sprintf("SELECT * FROM Conversation WHERE InitiatorID='%d' OR RecipientID='%d' ORDER BY ID DESC;", ID, ID)

	fmt.Println(query)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var conversations []Conversation
	scan.Rows(&conversations, results)

	fmt.Println(conversations)
	return conversations
}

func EditConversation(db *sql.DB, ID int, Initiator int, Recipient int, NoofMessage int, ConversationName string) {
	ConversationName = removespc(ConversationName)
	query := fmt.Sprintf(
		"UPDATE Conversation SET InitiatorID='%d', RecipientID='%d',NoofMessages='%d', ConversationName='%s' WHERE ID=%d",
		Initiator, Recipient, NoofMessage, ConversationName, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func InsertConversation(db *sql.DB, ID int, Initiator int, Recipient int, StartTime time.Time, NoofMessage int, ConversationName string) {
	ConversationName = removespc(ConversationName)
	query := fmt.Sprintf("INSERT INTO Conversation VALUES (%d, '%d', '%d', now(), '%d','%s')",
		ID, Initiator, Recipient, NoofMessage, ConversationName)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func DeleteConversation(db *sql.DB, ID int) {
	query := fmt.Sprintf(
		"DELETE FROM Conversation WHERE ID='%d'", ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func GetRepliesbyConversation(code string) []Replies {
	url := replyURL
	if code != "" {
		//url = baseURL + "/" + code + "?key=" + key
		url = repliesURL + "/" + code
	}
	fmt.Println("URL: ", url)
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		//var trip Trip
		var trips []Replies
		err := json.Unmarshal([]byte(data), &trips)
		fmt.Println(trips)
		fmt.Println(err)
		response.Body.Close()
		return trips
	}
	return []Replies{}
}

func GetaConversation(db *sql.DB, code string) aConversation {
	url := replyURL
	var aConvo aConversation
	if code != "" {
		//url = baseURL + "/" + code + "?key=" + key
		url = repliesURL + "/" + code
	}
	fmt.Println("URL: ", url)
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		//var trip Trip
		var replies []Replies
		err := json.Unmarshal([]byte(data), &replies)
		fmt.Println(err)
		response.Body.Close()
		aConvo.Replies = replies
		var convo Conversation
		x, _ := strconv.ParseInt(code, 10, 64)
		fmt.Println("Selected Conversation ID: ", x)
		convo = GetConversation(db, int(x))
		aConvo.Convo = convo
		return aConvo
	}
	return aConversation{}
}

//ROUTING CONFIGURATION
func conversation(w http.ResponseWriter, r *http.Request) {
	/*if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key"))
		return
	}*/
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/edufi?parseTime=true")
	//db, err := sql.Open("mysql", "root:password@tcp(database15:3306)/edufi?parseTime=true")
	// handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened")
	}

	params := mux.Vars(r)
	var ID = params["conversation_id"]
	x, _ := strconv.ParseInt(ID, 10, 64)
	fmt.Println("Selected ID: ", x)

	if r.Method == "GET" {
		fmt.Println("Get Request Called")
		p := GetConversation(db, int(x))
		json.NewEncoder(w).Encode(
			p)
		defer db.Close()
	}
	//DELETE
	if r.Method == "DELETE" {
		fmt.Println("Delete Request Called")
		fmt.Println("Delete Conversation ID: ", ID)
		DeleteConversation(db, int(x)) //DELETE DRIVER WITH ID OF VAR X
		defer db.Close()
	}

	if r.Header.Get("Content-type") == "application/json" {
		fmt.Println("Content Type is application/json")

		// POST is for creating new course
		if r.Method == "POST" {
			fmt.Println("Post Request Called")
			var newconversation Conversation
			reqBody1, err := ioutil.ReadAll(r.Body)
			if err == nil {
				var ID = GetIDConversation(db) // GET SUITABLE ID FOR CONVERSATION
				json.Unmarshal(reqBody1, &newconversation)
				if newconversation.InitiatorID == 0 {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply course " +
							"information " + "in JSON format"))
					return
				}
				//Automatically Assigns ID AND CREATE DRIVER
				InsertConversation(db, ID, newconversation.InitiatorID, newconversation.RecipientID, time.Now(), newconversation.NoofMessages, newconversation.ConversationName)
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply course information " +
					"in JSON format"))
			}
		}

		//---PUT is for creating or updating
		// existing course---
		if r.Method == "PUT" {
			fmt.Println("Put Request Called")

			var newconversation Conversation
			reqBody1, err := ioutil.ReadAll(r.Body)
			if err == nil {
				json.Unmarshal(reqBody1, &newconversation)
				//UPDATE DRIVER
				EditConversation(db, int(x), newconversation.InitiatorID, newconversation.RecipientID, newconversation.NoofMessages, newconversation.ConversationName)
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " +
					"course information " +
					"in JSON format"))
			}
		}

	}
}

func aconversation(w http.ResponseWriter, r *http.Request) {
	/*if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key"))
		return
	}*/
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/edufi?parseTime=true")
	//db, err := sql.Open("mysql", "root:password@tcp(database15:3306)/edufi?parseTime=true")
	// handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened")
	}

	params := mux.Vars(r)
	var ID = params["conversation_id"]
	x, _ := strconv.ParseInt(ID, 10, 64)
	fmt.Println("Selected ID: ", x)

	if r.Method == "GET" {
		fmt.Println("Get Request Called")
		p := GetaConversation(db, ID)
		json.NewEncoder(w).Encode(
			p)
		defer db.Close()
	}
}

func allconversation(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to the REST API!")
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/edufi?parseTime=true")
	//db, err := sql.Open("mysql", "root:password@tcp(database15:3306)/edufi?parseTime=true")
	// handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened")
	}
	if r.Method == "GET" {
		fmt.Println("Get Request Called")
		p := GetAllConversation(db) //RETRIEVE PASSENGER AND RETURN TO UI
		json.NewEncoder(w).Encode(
			p)
		defer db.Close()
	}
}

func removespc(str string) string {
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

func user(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to the REST API!")
	params := mux.Vars(r)
	var ID = params["user_id"]
	x, _ := strconv.ParseInt(ID, 10, 64)
	fmt.Println("Selected User ID: ", x)
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/edufi?parseTime=true")
	//db, err := sql.Open("mysql", "root:password@tcp(database15:3306)/edufi?parseTime=true")
	// handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened")
	}
	if r.Method == "GET" {
		fmt.Println("Get Request Called")
		p := GetAllConversationOfUser(db, int(x)) //RETRIEVE PASSENGER AND RETURN TO UI
		json.NewEncoder(w).Encode(
			p)
		defer db.Close()
	}
}

func student(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to the REST API!")
	//params := mux.Vars(r)
	//var ID = params["user_id"]
	//x, _ := strconv.ParseInt(ID, 10, 64)
	//fmt.Println("Selected User ID: ", x)
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/edufi?parseTime=true")
	//db, err := sql.Open("mysql", "root:password@tcp(database15:3306)/edufi?parseTime=true")
	// handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened")
	}
	if r.Method == "GET" {
		fmt.Println("Get Request Called")
		p := GetStudent(db) //RETRIEVE PASSENGER AND RETURN TO UI
		json.NewEncoder(w).Encode(
			p)
		defer db.Close()
	}
}

type Student struct {
	Student_id    string
	Name          string
	Date_of_birth string
	Address       string
	Phone_number  string
}

/*
func GetStudent(db *sql.DB) Student {
	var query string
	//query = fmt.Sprintf("SELECT * FROM Conversation WHERE InitiatorID='%d' OR RecipientID='%d' ORDER BY ID DESC;", ID, ID)
	query = fmt.Sprintf("SELECT * FROM Student WHERE student_id=1")

	fmt.Println(query)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		// map this type to the record in the table
		var c Student
		err = results.Scan(&c.Student_id, &c.Name,
			&c.Date_of_birth, &c.Address, &c.Phone_number)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(c.Student_id, c.Name,
			c.Date_of_birth, c.Address, c.Phone_number)
		return c
	}
	return Student{}
}*/

func GetStudent(db *sql.DB) []Student {
	var query string
	query = fmt.Sprintf("SELECT * FROM Student;")

	fmt.Println(query)
	results, err := db.Query(query)

	fmt.Println(results)
	if err != nil {
		panic(err.Error())
	}
	var students []Student
	scan.Rows(&students, results)
	fmt.Println(students)
	return students
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/conversation/", allconversation)
	router.HandleFunc("/api/v1/conversation/user/{user_id}", user)
	router.HandleFunc("/api/v1/conversation/{conversation_id}", conversation).Methods(
		"GET", "PUT", "POST", "DELETE")
	router.HandleFunc("/api/v1/aconversation/{conversation_id}", aconversation).Methods(
		"GET", "PUT", "POST", "DELETE")

	router.HandleFunc("/api/v1/students", student).Methods(
		"GET", "PUT", "POST", "DELETE")

	fmt.Println("Listening at port 9155")
	log.Fatal(http.ListenAndServe(":9155", router))

}
