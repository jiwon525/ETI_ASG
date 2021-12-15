package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Trips struct {
	CurrLocation int
	DestLocation int
	DriverID     int //driver NRIC
	PUserName    int //passenger username
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
func NewTrip(db *sql.DB, CL int, DL int, UN string) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db")
	sqlStatement := fmt.Sprintf("INSERT INTO Trip (CurrLocation, DestLocation, UserName) VALUES (%d,%d,'%s')", CL, DL, UN)

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}
