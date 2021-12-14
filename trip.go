package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Trips struct {
	CurrLocation int
	DestLocation int
	DriverID     int
}

func GetTrip(db *sql.DB) {
	results, err := db.Query("Select * FROM trip_db.Trip")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var trip Trips
		err = results.Scan(&trip.CurrLocation, &trip.DestLocation,
			&trip.DriverID)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(trip.CurrLocation, trip.DestLocation,
			trip.DriverID)
	}
}
func NewTrip(db *sql.DB, CL int, DL int, D string) {
	query := fmt.Sprintf("INSERT INTO Trip VALUES (%d, %d,'%s')",
		CL, DL, D)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}
