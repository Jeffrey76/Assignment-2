package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	//"reflect"
	//"strconv"
)

//TO USE IN DOCKER
//const conversationURL = "http://conversationapi15:9155/api/v1/conversation"
//const replyURL = "http://repliesapi15:9156/api/v1/reply"

//TO USE IN LOCAL COMPUTER
const conversationURL = "http://localhost:9155/api/v1/conversation"
const replyURL = "http://localhost:9156/api/v1/reply"

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

//GET Conversation of Code ID
func GetConversation(code string) Conversation {
	url := replyURL
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
			fmt.Println(err)
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
func GetConversation1(code string) aConversation {
	url := replyURL
	if code != "" {
		url = conversationURL + "/" + code
	}
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
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
		return aconvo
	}
	return aConversation{}
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
			fmt.Println(err)
		}
		fmt.Println("ID            :", driver.ID)
		fmt.Println("ConversationID    :", driver.ConversationID)
		fmt.Println("SenderID     :", driver.SenderID)
		fmt.Println("ReceiverID :", driver.ReceiverID)
		fmt.Println("ReceiverID :", driver.TimeSent)
		fmt.Println("Header     :", driver.Header)
		fmt.Println("Content     :", driver.Content)
		response.Body.Close()
		return driver
	}
	return Replies{}
}

//POST
func addReply(code string, p Replies) {

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

func main() {
	loop := true
	for loop == true {
		fmt.Println("Conversation Console")
		fmt.Println("---------------------")
		//DONE
		fmt.Println("1) Create Conversation")
		fmt.Println("2) View Conversation")
		fmt.Println("3) Edit Conversation")
		fmt.Println("4) Delete Conversation")

		fmt.Println("5) Get All Conversations")
		fmt.Println("6) Get All Conversation of User")
		//DONE
		fmt.Println("11) Create Reply")
		fmt.Println("12) View Reply")
		fmt.Println("13) Edit Reply")
		fmt.Println("14) Delete Reply")

		//NOT DONE
		fmt.Println("15) Get All Replies")
		fmt.Println("16) Get All Replies of User")

		fmt.Println("0) Quit")
		fmt.Print("Option: ")
		var option string
		fmt.Scanln(&option)
		//loop2 := true
		switch option {
		case "1":
			fmt.Println("Create Conversation ----------------")
			//fmt.Print("ID: ")
			//var ID string
			//fmt.Scanln(&ID)
			fmt.Print("InitiatorID: ")
			var IID int
			fmt.Scanln(&IID)
			fmt.Print("RecipientID: ")
			var RID int
			fmt.Scanln(&RID)
			fmt.Print("NoofMessages: ")
			var NooM int
			fmt.Scanln(&NooM)
			fmt.Print("ConversationName: ")
			var ConvoName string
			fmt.Scanln(&ConvoName)
			p := Conversation{
				//ID:               ID,
				InitiatorID:      IID,
				RecipientID:      RID,
				NoofMessages:     NooM,
				ConversationName: ConvoName,
			}
			addConversation("1", p)
		case "2":
			fmt.Print("ConversationID: ")
			var ID string
			fmt.Scanln(&ID)
			GetConversation(ID)
			//GetConversation1(ID)
		case "3":
			fmt.Println("Update Conversation ----------------")
			/*fmt.Print("ID: ")
			var ID int
			fmt.Scanln(&ID)*/
			fmt.Print("InitiatorID: ")
			var IID int
			fmt.Scanln(&IID)
			fmt.Print("RecipientID: ")
			var RID int
			fmt.Scanln(&RID)
			fmt.Print("NoofMessages: ")
			var NooM int
			fmt.Scanln(&NooM)
			fmt.Print("ConversationName: ")
			var ConvoName string
			fmt.Scanln(&ConvoName)
			p := Conversation{
				//ID:               ID,
				InitiatorID:      IID,
				RecipientID:      RID,
				NoofMessages:     NooM,
				ConversationName: ConvoName,
			}
			updateConversation("1", p)
		case "4":
			fmt.Print("ConversationID: ")
			var ID string
			fmt.Scanln(&ID)
			DeleteConversation(ID)
		case "5":
			GetConversations()
		case "6":

		case "11":
			fmt.Println("Create Reply ----------------")
			//fmt.Print("ID: ")
			//var ID string
			//fmt.Scanln(&ID)
			fmt.Print("ConversationID: ")
			var CID int
			fmt.Scanln(&CID)
			fmt.Print("SenderID: ")
			var SID int
			fmt.Scanln(&SID)
			fmt.Print("ReceiverID: ")
			var RID int
			fmt.Scanln(&RID)
			fmt.Print("Header: ")
			var Header string
			fmt.Scanln(&Header)
			fmt.Print("Content: ")
			var Content string
			fmt.Scanln(&Content)
			p := Replies{
				//ID:               ID,
				ConversationID: CID,
				SenderID:       SID,
				ReceiverID:     RID,
				Header:         Header,
				Content:        Content,
			}
			addReply("1", p)
		case "12":
			fmt.Print("ReplyID: ")
			var ID string
			fmt.Scanln(&ID)
			GetReply(ID)
		case "13":
			fmt.Println("Update Reply ----------------")
			fmt.Print("ID: ")
			var ID int
			fmt.Scanln(&ID)
			fmt.Print("ConversationID: ")
			var CID int
			fmt.Scanln(&CID)
			fmt.Print("SenderID: ")
			var SID int
			fmt.Scanln(&SID)
			fmt.Print("ReceiverID: ")
			var RID int
			fmt.Scanln(&RID)
			fmt.Print("Header: ")
			var Header string
			fmt.Scanln(&Header)
			fmt.Print("Content: ")
			var Content string
			fmt.Scanln(&Content)
			p := Replies{
				ID:             ID,
				ConversationID: CID,
				SenderID:       SID,
				ReceiverID:     RID,
				Header:         Header,
				Content:        Content,
			}
			updateReply("1", p)
		case "14":
			fmt.Print("ReplyID: ")
			var ID string
			fmt.Scanln(&ID)
			DeleteReply(ID)

			//addPassenger("0", p)
			//fmt.Println("Your Account Details. Please Remember Your ID and Password.")
			//GetPassenger("0")

			/*fmt.Print("ID: ")
				var ID int
				fmt.Scanln(&ID)
				val := ID
				string_id := fmt.Sprintf("%d", val)

				temp := CheckPassengerPassword(string_id) //RETRIEVE PASSENGER
				fmt.Print("Password: ")
				var Password string
				fmt.Scanln(&Password)
				if Password != temp { //CHECK PASSWORD
					fmt.Println("Wrong Password / ID")
					loop2 = false
				}
				for loop2 == true {
					fmt.Println("Passenger Management")
					fmt.Println("---------------------")
					fmt.Println("1) Edit")
					fmt.Println("2) View Account")
					fmt.Println("3) Request Trip")
					fmt.Println("4) View Trip")
					fmt.Println("0) Quit")
					fmt.Print("Option: ")
					var option1 string
					fmt.Scanln(&option1)
					switch option1 {
					case "1":
						//Edit Passenger
						b := true
						for b {
							fmt.Println("Current Details : ")
							passenger := GetPassenger(string_id) //Print Current Information
							fmt.Println("Choose Which Detail to Edit: ")
							fmt.Println("1) First Name 2) Last Name 3) Mobile Number 4) Email Address 5) Password 0) Exit Edit")
							fmt.Print("Option: ")
							var option2 string
							fmt.Scanln(&option2)
							switch option2 {
							case "1":
								fmt.Print("First Name: ")
								var FN string
								fmt.Scanln(&FN)
								passenger.FirstName = FN
							case "2":
								fmt.Print("Last Name: ")
								var LN string
								fmt.Scanln(&LN)
								passenger.LastName = LN
							case "3":
								fmt.Print("Phone: ")
								var PhoneNo string
								fmt.Scanln(&PhoneNo)
								passenger.MobileNo = PhoneNo
							case "4":
								fmt.Print("Email: ")
								var Email string
								fmt.Scanln(&Email)
								passenger.EmailAddr = Email
							case "5":
								fmt.Print("Password: ")
								var Password string
								fmt.Scanln(&Password)
								passenger.Password = Password
							case "0":
								b = false
							}
							val := ID
							i := fmt.Sprintf("%d", val)
							updatePassenger(i, passenger)
						}
					case "2":
						GetPassenger(string_id)

					case "0":
						loop2 = false
					}
				}
			case "2":
				fmt.Print("ID: ")
				var ID int
				fmt.Scanln(&ID)
				val := ID
				string_id := fmt.Sprintf("%d", val)
				temp := CheckDriverPassword(string_id) //RETRIEVE DRIVER
				fmt.Print("Password: ")
				var Password string
				fmt.Scanln(&Password)
				if Password != temp { //CHECK PASSWORD
					fmt.Println("Wrong Password / ID")
					loop2 = false
				}

				for loop2 == true {
					fmt.Println("Driver Management")
					fmt.Println("---------------------")
					fmt.Println("1) Edit")
					fmt.Println("2) View Account")
					fmt.Println("3) Update Trip")
					fmt.Println("0) Quit")
					fmt.Print("Option: ")
					var option1 string
					fmt.Scanln(&option1)
					switch option1 {
					case "1":
						//Edit Driver
						b := true
						for b {
							fmt.Println("Current Details : ")
							driver := GetDriver(string_id) //Print Current Information
							fmt.Println("Choose Which Detail to Edit: ")
							fmt.Println("1) First Name 2) Last Name 3) Mobile Number 4) Email Address 5) Licence Plate 6) Status 7) Password 0) Exit Edit")
							fmt.Print("Option: ")
							var option2 string
							fmt.Scanln(&option2)
							switch option2 {
							case "1":
								fmt.Print("First Name: ")
								var FN string
								fmt.Scanln(&FN)
								driver.FirstName = FN
							case "2":
								fmt.Print("Last Name: ")
								var LN string
								fmt.Scanln(&LN)
								driver.LastName = LN
							case "3":
								fmt.Print("Phone: ")
								var PhoneNo string
								fmt.Scanln(&PhoneNo)
								driver.MobileNo = PhoneNo
							case "4":
								fmt.Print("Email: ")
								var Email string
								fmt.Scanln(&Email)
								driver.EmailAddr = Email
							case "5":
								fmt.Print("Licence Plate: ")
								var licence string
								fmt.Scanln(&licence)
								driver.LicencePlate = licence
							case "6":
								fmt.Print("Status: ")
								var Status string
								fmt.Scanln(&Status)
								driver.DriverStatus = Status
							case "7":
								fmt.Print("Password: ")
								var Password string
								fmt.Scanln(&Password)
								driver.Password = Password
							case "0":
								b = false
							}
							val := ID
							i := fmt.Sprintf("%d", val)
							updateDriver(i, driver)
						}

					case "2":
						GetDriver(string_id)

					case "0":
						loop2 = false
					}
				}
			case "3":
				fmt.Println("Sign Up----------------")
				fmt.Println("1) Passenger Sign Up")
				fmt.Println("2) Driver Sign Up")
				fmt.Println("0) Exit")
				fmt.Print("Option: ")
				var option1 string
				fmt.Scanln(&option1)
				switch option1 {
				case "1":
					//ADD Passenger
					fmt.Println("Passenger Sign Up----------------")
					fmt.Print("First Name: ")
					var FN string
					fmt.Scanln(&FN)
					fmt.Print("Last Name: ")
					var LN string
					fmt.Scanln(&LN)
					fmt.Print("Phone: ")
					var PhoneNo string
					fmt.Scanln(&PhoneNo)
					fmt.Print("Email: ")
					var Email string
					fmt.Scanln(&Email)
					fmt.Print("Password: ")
					var Password string
					fmt.Scanln(&Password)
					p := Passenger{
						FirstName: FN,
						LastName:  LN,
						MobileNo:  PhoneNo,
						EmailAddr: Email,
						Password:  Password,
					}
					addPassenger("0", p)
					fmt.Println("Your Account Details. Please Remember Your ID and Password.")
					GetPassenger("0")
				case "2":
					//ADD Driver
					fmt.Println("Driver Sign Up-------------")
					//ADD Driver
					fmt.Print("First Name: ")
					var FN string
					fmt.Scanln(&FN)
					fmt.Print("Last Name: ")
					var LN string
					fmt.Scanln(&LN)
					fmt.Print("Phone: ")
					var PhoneNo string
					fmt.Scanln(&PhoneNo)
					fmt.Print("Email: ")
					var Email string
					fmt.Scanln(&Email)
					fmt.Print("Identification Number: ")
					var ID_No string
					fmt.Scanln(&ID_No)
					fmt.Print("Licence Plate: ")
					var licence string
					fmt.Scanln(&licence)
					fmt.Print("Password: ")
					var Password string
					fmt.Scanln(&Password)
					d := Driver{
						FirstName:    FN,
						LastName:     LN,
						MobileNo:     PhoneNo,
						EmailAddr:    Email,
						ID_No:        ID_No,
						LicencePlate: licence,
						Password:     Password,
					}
					addDriver("0", d)
					GetDriver("0")
				case "0":
					fmt.Println("Exiting...")
					//default:
					//	fmt.Println("No Matching Option")
				}
			case "0":
				fmt.Println("Quitting...")
				loop = false*/
			//default:
			//fmt.Println("No matching option")

		}

	}
}
