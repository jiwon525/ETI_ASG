package main //driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Drivers struct {
	FirstName    string
	LastName     string
	MobileNo     int
	Email        string
	Password     string
	NRIC         string
	Carlicense   string
	Availability int
}

func GetAvailDriver(db *sql.DB) {

}
func GetDriver(db *sql.DB) {
	results, err := db.Query("Select * FROM driver_db.Driver")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var driver Drivers
		err = results.Scan(&driver.FirstName, &driver.LastName,
			&driver.MobileNo, &driver.Email, &driver.Password, &driver.NRIC, &driver.Carlicense)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(driver.FirstName, driver.LastName,
			driver.MobileNo, driver.Email, driver.Password, driver.NRIC, driver.Carlicense)
	}
}

func NewDriver(db *sql.DB, FN string, LN string, MNo int, Email string, Pw string, NRIC string, CL string, Avail int) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
	sqlStatement := fmt.Sprintf("INSERT INTO Driver (FirstName, LastName, MobileNo, Email, Password, NRIC, CarLicense, Availability) VALUES ('%s', '%s', %d , '%s', '%s', '%s', '%s', %d)", FN, LN, MNo, Email, Pw, NRIC, CL, Avail)

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func EditDriver(db *sql.DB, FN string, LN string, MNo int, Email string, Pw string, NRIC string, CL string, Avail int) {
	query := fmt.Sprintf(
		"UPDATE Driver SET FirstName='%s', LastName='%s', MobileNumber=%d, Email='%s', Password='%s', CarLicense='%s', WHERE NRIC='%s'",
		FN, LN, MNo, Email, Pw, CL, NRIC)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
} //EditRecord(db, 2, "Taylor", "Swift", 23)
