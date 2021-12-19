package main //main.go

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func driveroptions(NRIC string) {
	char := 0
	var ATrips []Trips
	var Duser []Driver
	DID := FindDriverID(db, NRIC, Duser)
	//TID := FindTripID(db, DID, ATrips) to find trip id, but not needed.
	var end bool
	for end == false {
		fmt.Println("1.Initiate trip\n2.Edit account details\n3.Back")
		fmt.Scanf("%d\n", &char)
		switch char {
		case 1:
			option := 0
			fmt.Println("1.Start trip\n2.End Trip\n3.Back")
			fmt.Scanf("%d\n", &option)
			if option == 1 {
				fmt.Println("initiating trip")
				TripsAssigned(db, DID, ATrips)
				fmt.Println("Would you like to accept? y for yes n for no")
				var trip string
				fmt.Scanf("%s\n", &trip)
				if trip == "y" {
					InitiateStartTrip(db, DID, ATrips)
				} else if trip == "n" {
					DeleteTrip(db, DID, ATrips)
				}

			} else if option == 2 {
				InitiateEndTrip(db, DID, ATrips)
			} else if option == 3 {
				continue
			} else {
				fmt.Println("wrong input, try again")
				continue
			}
		case 2:
			var FN, LN, Email, Pw, Cl string
			var MNo int
			var Avail bool
			fmt.Println("Editing account details.\nPlease enter your new first name: ")
			fmt.Scanf("%s\n", &FN)
			fmt.Println("please enter your new last name: ")
			fmt.Scanf("%s\n", &LN)
			fmt.Println("please enter your new mobile number: ") //to have check if their nmbers are the correct number
			fmt.Scanf("%d\n", &MNo)
			fmt.Println("please enter new your email: ")
			fmt.Scanf("%s\n", &Email)
			fmt.Println("please enter new your car license: ")
			fmt.Scanf("%s\n", &Cl)
			fmt.Println("please enter a new password to use: ")
			fmt.Scanf("%s\n", &Pw)
			Avail = false
			EditDriver(db, DID, FN, LN, MNo, Email, Pw, NRIC, Cl, Avail)
			fmt.Println("Details have been updated.")
		case 3:
			end = true
			return
		default:
			fmt.Println("wrong input, please try again.")
		}
	}

}
func passengeroptions(Username string) {
	char := 0
	var ATrips []Trips
	var Duser []Driver
	var Puser []Passenger
	ID := FindPassengerID(db, Username, Puser)
	var end bool
	for end == false {
		fmt.Println("\n1.Book a ride\n2.See status of current Trip\n3.Retrieve past trips\n4.Edit account details\n5.Back")
		fmt.Scanf("%d\n", &char)
		switch char {
		case 1:
			var CL, DL int

			fmt.Println("Current location postal code: ")
			fmt.Scanf("%d\n", &CL)
			fmt.Println("Destination location postal code: ")
			fmt.Scanf("%d\n", &DL)
			DI := GetAvailDriver(db, Duser)
			if DI == 0 {
				continue
			} else {
				NewTrip(db, CL, DL, ID, DI, Username)
			}

		case 2:
			AllTrips(db, ATrips, ID, 0) //if there are no trips, to print something like sorry, no past trips
		case 3:
			var FN, LN, Email, Pw string
			var MNo int
			fmt.Println("Editing account details.\nPlease enter your new first name: ")
			fmt.Scanf("%s\n", &FN)
			fmt.Println("please enter your new last name: ")
			fmt.Scanf("%s\n", &LN)
			fmt.Println("please enter your new mobile number: ") //to have check if their nmbers are the correct number
			fmt.Scanf("%d\n", &MNo)
			fmt.Println("please enter new your email: ")
			fmt.Scanf("%s\n", &Email)
			fmt.Println("please enter a new password to use: ")
			fmt.Scanf("%s\n", &Pw)
			EditPassenger(db, Username, FN, LN, MNo, Email, Pw, ID)
		case 4:
			end = true
			return
		default:
			fmt.Println("wrong input, please try again.")
		}
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
			var end bool
			var UN, Pw string
			for end == false {
				fmt.Println("Please enter your UserName: ")
				fmt.Scanf("%s\n", &UN)
				fmt.Println("Please enter your Password: ")
				fmt.Scanf("%s\n", &Pw)
				var PUser []Passenger
				loginstatus := CheckLogin(db, PUser, UN, Pw)
				if loginstatus == 1 {
					passengeroptions(UN)
					end = true
				} else if loginstatus == 2 {
					continue
				} else if loginstatus == 3 {
					end = true
				} else {
					fmt.Println("error")
					end = true
				}

			}

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
			fmt.Println("\nsuccessfully signed up, repeating details:\n" + "username: " + UN + "\nfull name: " + FN + LN + "\nmobile number: " + Mno + "\nEmail: " + Email + "\nPassword: " + Pw + "\n")

		case 3:
			var end bool
			var NRIC, Pw string
			for end == false {
				fmt.Println("Please enter your NRIC: ")
				fmt.Scanf("%s\n", &NRIC)
				fmt.Println("Please enter your Password: ")
				fmt.Scanf("%s\n", &Pw)
				var Duser []Driver
				loginstatus := CheckDriver(db, Duser, NRIC, Pw)
				if loginstatus == 1 {
					driveroptions(NRIC)
					end = true
				} else if loginstatus == 2 {
					continue
				} else if loginstatus == 3 {
					end = true
				} else {
					fmt.Println("error")
					end = true
				}

			}

		case 4:
			var FN, LN, Email, Pw, NRIC, CarL string
			var MNo int
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
			Avail := false
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
