package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Trips struct {
	TripID       int
	CurrLocation int
	DestLocation int
	DriverID     int
	PassengerID  int
	StartTrip    bool
	EndTrip      bool
}

func ScanTripDB(db *sql.DB, T []Trips) []Trips {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db")
	results, err := db.Query("SELECT * FROM trip_db.Trip")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		// map this type to the record in the table
		var trip Trips
		err = results.Scan(&trip.TripID, &trip.CurrLocation, &trip.DestLocation, &trip.DriverID, &trip.PassengerID, &trip.StartTrip, &trip.EndTrip)
		if err != nil {
			panic(err.Error())
		}
		T = append(T, trip)
		//fmt.Println(T)
	}
	return T
}
func GetTrip(db *sql.DB, Username string) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db")
	results, err := db.Query("Select * FROM trip_db.Trip WHILE USERNAME : 's%' ", Username)

	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var trip Trips
		err = results.Scan(&trip.TripID, &trip.CurrLocation, &trip.DestLocation,
			&trip.DriverID, &trip.PassengerID, &trip.StartTrip, &trip.EndTrip)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(trip.TripID, trip.CurrLocation, trip.DestLocation, //to print out the trips that user of the certain username.
			trip.DriverID, trip.PassengerID, trip.StartTrip, trip.EndTrip)
	}
}
func NewTrip(db *sql.DB, CL int, DL int, PassengerID int, DriverID int, UN string) { //current location, destination location, username, driver ID
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db")
	sqlStatement := fmt.Sprintf("INSERT INTO Trip (CurrLocation, DestLocation, DriverID, PassengerID, StartTrip, EndTrip) VALUES (%d,%d,%d,%d,%t, %t)", CL, DL, DriverID, PassengerID, false, false)
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	cl, dl, did := strconv.Itoa(CL), strconv.Itoa(DL), strconv.Itoa(DriverID)
	fmt.Println("Current Location: " + cl + "\nDestination Location : " + dl + "\nYour Username: " + UN + "\nDriver details: " + did + "\nPlease wait for the driver.")
}

func AllTrips(db *sql.DB, ATrips []Trips, PID int) { //for printing of all past trips
	aTrips := ScanTripDB(db, ATrips)
	for _, v := range aTrips {
		if v.PassengerID == PID {
			aTrips = append(aTrips, v)
		}
		cl := strconv.Itoa(v.CurrLocation)
		dl := strconv.Itoa(v.DestLocation)
		did := strconv.Itoa(v.DriverID)
		pid := strconv.Itoa(PID)
		fmt.Println("Start Location: " + cl + "\nEnd Location : " + dl + "\nPassenger details: " + pid + "\nDriver details: " + did + "\n")

	}
}

func FindTripID(db *sql.DB, DID int, T []Trips) int {
	TripList := ScanTripDB(db, T)
	for _, v := range TripList {
		if v.DriverID == DID {
			return v.TripID
		}
	}
	return 0
}

func InitiateStartTrip(db *sql.DB, DID int, T []Trips) {
	aTrips := ScanTripDB(db, T)
	for _, v := range aTrips {
		if v.DriverID == DID {
			if v.StartTrip == false {
				v.StartTrip = true
				EditTrip(db, v.TripID, v.CurrLocation, v.DestLocation, v.DriverID, v.PassengerID, v.StartTrip, v.EndTrip)
				tid, cl, dl, did, pid := strconv.Itoa(v.TripID), strconv.Itoa(v.CurrLocation), strconv.Itoa(v.DestLocation), strconv.Itoa(v.DriverID), strconv.Itoa(v.PassengerID)
				fmt.Println("Your Trip has started.\nTrip ID: " + tid + "\nPassenger going from " + cl + " to " + dl + "\nDriver ID: " + did + " Passenger ID: " + pid)
				return
			}
		}

	}
	fmt.Println("you have no assigned trips")
	return
}

func InitiateEndTrip(db *sql.DB, DID int, T []Trips) {
	aTrips := ScanTripDB(db, T)
	for _, v := range aTrips {
		if v.DriverID == DID {
			if v.StartTrip == true {
				if v.EndTrip == false {
					v.EndTrip = true
					EditTrip(db, v.TripID, v.CurrLocation, v.DestLocation, v.DriverID, v.PassengerID, v.StartTrip, v.EndTrip)
					tid := strconv.Itoa(v.TripID)
					fmt.Println("Trip ID: " + tid + "\nYour Trip has Ended.")
					return
				}
			}
		}
	}
	fmt.Println("trip has not started yet.")
	return
}

func EditTrip(db *sql.DB, TID int, CL int, DL int, DID int, PID int, ST bool, ET bool) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db")
	sqlStatement2 := fmt.Sprintf("UPDATE Trip SET CurrLocation=%d, DestLocation=%d, PassengerID=%d, StartTrip=%t, EndTrip=%t WHERE DriverID=%d", CL, DL, PID, ST, ET, DID)
	_, err = db.Exec(sqlStatement2)
	if err != nil {
		panic(err)
	}
}

func TripsAssigned(db *sql.DB, DID int, T []Trips) {
	aTrips := ScanTripDB(db, T)
	for _, v := range aTrips {
		if v.DriverID == DID {
			if v.StartTrip == false {
				cl, dl := strconv.Itoa(v.CurrLocation), strconv.Itoa(v.DestLocation)
				fmt.Println("One trip assigned to you.\nPassenger going from " + cl + " to " + dl)
			}
		}
	}
}

func DeleteTrip(db *sql.DB, DID int, T []Trips) {
	//this is for when drivers do not want to take the customer
	aTrips := ScanTripDB(db, T)
	for _, v := range aTrips {
		if v.DriverID == DID {
			if v.StartTrip == false {
				v.EndTrip = true
				EditTrip(db, v.TripID, v.CurrLocation, v.DestLocation, v.DriverID, v.PassengerID, v.StartTrip, v.EndTrip)
				tid := strconv.Itoa(v.TripID)
				fmt.Println("Trip ID: " + tid + "\nYour Trip has been deleted.")
				return
			}
		}
	}
	fmt.Println("trip does not exist.")
	return
}

func TripStatus(db *sql.DB, PID int, T []Trips) {
	aTrips := ScanTripDB(db, T)
	for _, v := range aTrips {
		if v.PassengerID == PID {
			if v.StartTrip == false {
				if v.EndTrip == true {
					fmt.Println("Your driver has cancelled. Please try booking again")
				} else if v.EndTrip == false {
					fmt.Println("Your driver has not picked you up")
				}
			} else {
				if v.EndTrip == false {
					fmt.Println("Driver has initiated the ride, you are on your way to your destination")
				} else {
					fmt.Println("Your trip has ended.") //issue of printing here
				}

			}
		}
	}
	return
}

/*
	for _, v := range aTrips {
		int := 0
		cl := strconv.Itoa(v.CurrLocation)
		dl := strconv.Itoa(v.DestLocation)
		fmt.Println("Start Location: " + cl + "\nEnd Location : " + dl + "\nYour Username: " + v.PUserName + "\nDriver details: " + v.DriverID + "\n")
		int++

	}
*/
