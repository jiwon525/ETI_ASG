package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Trips struct {
	CurrLocation int
	DestLocation int
	DriverID     string //driver NRIC
	PUserName    string //passenger username
}

func GetTrip(db *sql.DB, Username string) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db")
	results, err := db.Query("Select * FROM trip_db.Trip WHILE USERNAME : 's%' ", Username)

	if err != nil {
		panic(err.Error())
	}
	counter := 0
	for results.Next() {
		var trip Trips
		err = results.Scan(&trip.CurrLocation, &trip.DestLocation,
			&trip.DriverID, &trip.PUserName)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(trip.CurrLocation, trip.DestLocation, //to print out the trips that user of the certain username.
			trip.DriverID, trip.PUserName)
		counter++
	}
}
func NewTrip(db *sql.DB, CL int, DL int, UN string, DI string) { //current location, destination location, username, driver ID
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db")
	sqlStatement := fmt.Sprintf("INSERT INTO Trip (CurrLocation, DestLocation, UserName, DriverID) VALUES (%d,%d,'%s','%s')", CL, DL, UN, DI)
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	cl := strconv.Itoa(CL)
	dl := strconv.Itoa(DL)
	fmt.Println("Current Location: " + cl + "\nDestination Location : " + dl + "\nYour Username: " + UN + "\nDriver details: " + DI + "\nPlease wait for the driver.\n")
}

func AllTrips(db *sql.DB, ATrips []Trips, UN string) { //for printing of all past trips
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db")
	results, err := db.Query("SELECT * FROM trip_db.Trip")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		// map this type to the record in the table
		var trip Trips
		err = results.Scan(&trip.CurrLocation, &trip.DestLocation, &trip.PUserName, &trip.DriverID)
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(passenger.UserName, passenger.FirstName, passenger.LastName, passenger.MobileNo, passenger.Email, passenger.Password) was used to check if the data is correct

		ATrips = append(ATrips, trip)
	}
	var aTrips []Trips
	for _, v := range ATrips {
		int := 0
		if v.PUserName == UN {
			aTrips = append(aTrips, v)
		}
		cl := strconv.Itoa(v.CurrLocation)
		dl := strconv.Itoa(v.DestLocation)
		fmt.Println("Start Location: " + cl + "\nEnd Location : " + dl + "\nYour Username: " + v.PUserName + "\nDriver details: " + v.DriverID + "\n")
		int++

	} /*
		for _, v := range aTrips {
			int := 0
			cl := strconv.Itoa(v.CurrLocation)
			dl := strconv.Itoa(v.DestLocation)
			fmt.Println("Start Location: " + cl + "\nEnd Location : " + dl + "\nYour Username: " + v.PUserName + "\nDriver details: " + v.DriverID + "\n")
			int++

		}*/

}
