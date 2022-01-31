package main

//CHANGE DB URL TO LOCALHOST:3306 INSTEAD OF DATABASE15 TO USE IN LOCAL COMPUTER

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

//const key = "2c78afaf-97da-4816-bbee-9ad239abb210"

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
func GetIDReplies(db *sql.DB) int {
	query := fmt.Sprintf("SELECT COUNT(*) FROM Replies")

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

func GetReplies(db *sql.DB, ID int) Replies {
	var query string
	query = fmt.Sprintf("SELECT ID,ConversationID,SenderID,ReceiverID,TimeSent,Header,Content FROM Replies WHERE ID='%d' ORDER BY ID DESC;", ID)

	fmt.Println(query)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		// map this type to the record in the table
		var r Replies
		err = results.Scan(&r.ID, &r.ConversationID, &r.SenderID,
			&r.ReceiverID, &r.TimeSent, &r.Header, &r.Content)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(r.ID, r.ConversationID, r.SenderID,
			r.ReceiverID, r.TimeSent, r.Header, r.Content)
		return r
	}
	return Replies{}
}

func GetRepliesbyConversation(db *sql.DB, ID int) []Replies {
	var query string
	query = fmt.Sprintf("SELECT ID,ConversationID,SenderID,ReceiverID,TimeSent,Header,Content FROM Replies WHERE ConversationID='%d' ORDER BY ID DESC;", ID)

	fmt.Println(query)
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	var Replies []Replies
	scan.Rows(&Replies, results)

	fmt.Println(Replies)
	return Replies
}

func GetAllReplies(db *sql.DB) []Replies {
	var query string
	query = fmt.Sprintf("SELECT ID,ConversationID,SenderID,ReceiverID,TimeSent,Header,Content FROM Replies ORDER BY ID DESC;")

	fmt.Println(query)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var Replies []Replies
	scan.Rows(&Replies, results)

	fmt.Println(Replies)
	return Replies
}

//CURRENTLY UNUSED
func GetAllRepliesOfUser(db *sql.DB, ID int) []Replies {
	var query string
	query = fmt.Sprintf("SELECT ID,ConversationID,SenderID,ReceiverID,TimeSent,Header,Content FROM Replies WHERE SenderID='%d' OR ReceiverID='%d' ORDER BY ID DESC;", ID, ID)

	fmt.Println(query)
	results, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var Repliess []Replies
	scan.Rows(&Repliess, results)

	fmt.Println(Repliess)
	return Repliess
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

func EditReplies(db *sql.DB, ID int, Initiator int, Recipient int, Header string, Content string) {
	Header = removespc(Header)
	Content = removespc(Content)
	query := fmt.Sprintf(
		"UPDATE Replies SET SenderID='%d', ReceiverID='%d',Header='%s', Content='%s' WHERE ID=%d",
		Initiator, Recipient, Header, Content, ID)
	fmt.Println("QUERY: ", query)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func InsertReplies(db *sql.DB, ID int, ConversationID int, Initiator int, Recipient int, StartTime time.Time, Header string, Content string) {
	Header = removespc(Header)
	Content = removespc(Content)
	query := fmt.Sprintf("INSERT INTO Replies(ID,ConversationID,SenderID,ReceiverID,TimeSent,Header,Content) VALUES ('%d','%d', '%d', '%d', now(), '%s','%s')",
		ID, ConversationID, Initiator, Recipient, Header, Content)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func DeleteReplies(db *sql.DB, ID int) {
	query := fmt.Sprintf(
		"DELETE FROM Replies WHERE ID='%d'", ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

//ROUTING CONFIGURATION
func reply(w http.ResponseWriter, r *http.Request) {
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
	var ID = params["reply_id"]
	x, _ := strconv.ParseInt(ID, 10, 64)
	fmt.Println("Selected Conversation ID: ", x)

	if r.Method == "GET" {
		fmt.Println("Get Request Called")
		p := GetReplies(db, int(x)) //RETRIEVE PASSENGER AND RETURN TO UI
		//p := GetRepliesbyConversation(db, int(x))
		json.NewEncoder(w).Encode(
			p)
		defer db.Close()
	}
	//DELETE
	if r.Method == "DELETE" {
		fmt.Println("Delete Request Called")
		fmt.Println("Delete Replies ID: ", ID)
		DeleteReplies(db, int(x)) //DELETE DRIVER WITH ID OF VAR X
		defer db.Close()
	}

	if r.Header.Get("Content-type") == "application/json" {
		fmt.Println("Content Type is application/json")

		// POST is for creating new course
		if r.Method == "POST" {
			fmt.Println("Post Request Called")
			var newReplies Replies
			reqBody1, err := ioutil.ReadAll(r.Body)
			if err == nil {
				var ID = GetIDReplies(db) // GET SUITABLE ID FOR Replies
				json.Unmarshal(reqBody1, &newReplies)
				if newReplies.SenderID == 0 {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply course " +
							"information " + "in JSON format"))
					return
				}
				InsertReplies(db, ID, newReplies.ConversationID, newReplies.SenderID, newReplies.ReceiverID, time.Now(), newReplies.Header, newReplies.Content)
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

			var newReplies Replies
			reqBody1, err := ioutil.ReadAll(r.Body)
			if err == nil {
				json.Unmarshal(reqBody1, &newReplies)
				//UPDATE DRIVER
				EditReplies(db, newReplies.ID, newReplies.SenderID, newReplies.ReceiverID, newReplies.Header, newReplies.Content)
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

func replies(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("Selected Conversation ID: ", x)

	if r.Method == "GET" {
		fmt.Println("Get Request Called")
		p := GetRepliesbyConversation(db, int(x))
		json.NewEncoder(w).Encode(
			p)
		defer db.Close()
	}
	//DELETE
	if r.Method == "DELETE" {
		fmt.Println("Delete Request Called")
		fmt.Println("Delete Replies in ConversationID: ", ID)
		//DeleteReplies(db, int(x)) //DELETE DRIVER WITH ID OF VAR X
		defer db.Close()
	}
}

func allReplies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the REST API!")
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
		p := GetAllReplies(db) //RETRIEVE PASSENGER AND RETURN TO UI
		json.NewEncoder(w).Encode(
			p)
		defer db.Close()
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/replies/", allReplies)
	router.HandleFunc("/api/v1/replies/{conversation_id}", replies)
	router.HandleFunc("/api/v1/reply/{reply_id}", reply).Methods(
		"GET", "PUT", "POST", "DELETE")

	fmt.Println("Listening at port 9156")
	log.Fatal(http.ListenAndServe(":9156", router))

}
