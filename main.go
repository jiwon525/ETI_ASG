package main //main.go

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	//P "./passenger"
)

var db *sql.DB

func passengers() { //opening database for passenger
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/user_db")
	// handle error
	if err != nil {
		fmt.Println("Error with opening database")
		//panic(err.Error())
	}

	fmt.Println("Database opened")
	// defer the close till after the function has finished executing
	defer db.Close()
}
func drivers() { //opening database for driver
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/user_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Database opened")
	// defer the close till after the main function has finished executing
	defer db.Close()
}
func trips() { //opening database for trips
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db")
	// handle error
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Database opened")
	// defer the close till after the main function has finished executing
	defer db.Close()
}
func PLogin() {
	fmt.Println("log in page, input all the details.")

}
func passengeroptions(char int) {

	switch char {
	case 1:
		var CL, DL int
		fmt.Println("Current location postal code: ")
		fmt.Scanf("%s\n", &CL)
		fmt.Println("Destination location postal code: ")
		fmt.Scanf("%s\n", &DL)
	case 2:
		fmt.Println("retrieving past trips")

	case 3:
		fmt.Println("editing account details")
	case 4:
		return
	default:
		fmt.Println("wrong input, please try again.")
	}

}
func main() {
	var end bool
	char := 0
	for end == false {

		fmt.Println("1.Passenger Login \n2.Passenger SignUp \n3.Driver Login \n4.Driver SignUp\n5.Exit")
		fmt.Scanf("%d\n", &char)
		switch char {
		case 1:
			PLogin() //log in
			fmt.Println("1.Book a ride\n2.Retrieve past trips\n3.Edit account details\n4.Back")
			fmt.Scanf("%d\n", &char)
			passengeroptions(char)
		case 2:
			var UN, FN, LN, Email, Pw string
			var MNo int

			fmt.Println("please enter a username to use")
			fmt.Scanf("%s\n", &UN)
			fmt.Println("please enter your first name")
			fmt.Scanf("%s\n", &FN)
			fmt.Println("please enter your last name")
			fmt.Scanf("%s\n", &LN)
			fmt.Println("please enter your mobile number")
			fmt.Scanf("%d\n", &MNo)
			Mno := strconv.Itoa(MNo)
			fmt.Println("please enter your email")
			fmt.Scanf("%s\n", &Email)
			fmt.Println("please enter a password to use")
			fmt.Scanf("%s\n", &Pw)
			NewPassenger(db, UN, FN, LN, MNo, Email, Pw)
			fmt.Println("successfully signed up, repeating details:\n" + "username: " + UN + "\nfull name: " + FN + LN + "\nmobile number: " + Mno + "\nEmail: " + Email + "\nPassword: " + Pw + "\n")

		case 3:
			fmt.Println("driver log in")
		case 4:
			var FN, LN, Email, Pw, NRIC, CarL string
			var MNo, Avail int
			fmt.Println("please enter your first name")
			fmt.Scanf("%s\n", &FN)
			fmt.Println("please enter your last name")
			fmt.Scanf("%s\n", &LN)
			fmt.Println("please enter your mobile number")
			fmt.Scanf("%d\n", &MNo)
			Mno := strconv.Itoa(MNo)
			fmt.Println("please enter your email")
			fmt.Scanf("%s\n", &Email)
			fmt.Println("please enter a password to use")
			fmt.Scanf("%s\n", &Pw)
			fmt.Println("please enter your NRIC")
			fmt.Scanf("%s\n", &NRIC)
			fmt.Println("please enter Car License")
			fmt.Scanf("%s\n", &CarL)
			Avail = 0
			NewDriver(db, FN, LN, MNo, Email, Pw, NRIC, CarL, Avail)
			fmt.Println("successfully signed up, repeating details:\n" + "full name: " + FN + LN + "\nmobile number: " + Mno + "\nEmail: " + Email + "\nPassword: " + Pw + "\nNRIC: " + NRIC + "\nCar License: " + CarL + "\n")

		case 5:
			fmt.Println("exit")
			end = true
		default:
			fmt.Println("wrong input")

		}
	}

}

/*
type Driver struct {
	Cartype string `json:"car type"` //?
}

type Location struct { //start location & end location
	StartLocation string
	EndLocation   string
}

func main() {
	var passenger []Passenger
	jsonString :=
		`[
        {
            "firstname":"Wei-Meng",
            "lastname":"Lee"
        },
        {
            "firstname":"Mickey",
            "lastname":"Mouse"
        }

    ]`

	json.Unmarshal([]byte(jsonString), &passenger)
	for _, v := range user {
		fmt.Println(v.Firstname)
		fmt.Println(v.Lastname)
	}
	fmt.Println(passenger)

	jsonString2 :=
		`{
        "car type":"Hyndai"

     }`

	var driver Driver
	json.Unmarshal([]byte(jsonString2), &driver)
	fmt.Println(driver.Cartype) // Hyndai

	func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Ride App!")
}
func passenger(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of all courses") //should change to passenger details
	// returns the key/value pairs in the query string as a map object
	kv := r.URL.Query()

	for k, v := range kv {
		fmt.Println(k, v) // print out the key/value pair
	}

	router := mux.NewRouter()
	router.HandleFunc("/ride/", home).Methods("GET", "POST", "PUT")
	router.HandleFunc("/ride/registerPassenger", P.InsertRecord()).Methods("PUT")
	router.HandleFunc("/ride/passengerInterface", passengerInterface)
	router.HandleFunc("/ride/passenger/{Dpostalcode}", destination)
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

	func destination(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintf(w, "Destination Postal Code "+params["Dpostalcode"])

	func passengerInterface(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "passenger ID, first and last name, email")
}
}
}*/
