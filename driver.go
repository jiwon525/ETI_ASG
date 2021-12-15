package main //driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Driver struct {
	FirstName    string
	LastName     string
	MobileNo     int
	Email        string
	Password     string
	NRIC         string
	Carlicense   string
	Availability int
}

func CheckDriver(db *sql.DB, DUser []Driver, nric string, Pw string) int { //for driver verification
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
	results, err := db.Query("SELECT * FROM driver_db.Driver")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		// map this type to the record in the table
		var driver Driver
		err = results.Scan(&driver.FirstName, &driver.LastName,
			&driver.MobileNo, &driver.Email, &driver.Password, &driver.NRIC, &driver.Carlicense, &driver.Availability)
		if err != nil {
			panic(err.Error())
		}
		DUser = append(DUser, driver)
	}
	//fmt.Println(DUser) was used to check if data was correctly appended to list
	for _, v := range DUser {
		int := 0
		if v.NRIC == nric {
			if v.Password == Pw {
				fmt.Println("you are logged in.")
				return 1
			} else {
				fmt.Println("driver NRIC found but password is wrong")
				return 2
			}
		} else {
			int++
		}
	}
	fmt.Println("you have not signed up before, or have entered the wrong username")
	return 3
}

func GetAvailDriver(db *sql.DB, DUser []Driver) string {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
	results, err := db.Query("SELECT * FROM driver_db.Driver")
	if err != nil {
		panic(err.Error())
	}
	var driver Driver
	var availdriver []Driver
	defer results.Close()
	for results.Next() {
		// map this type to the record in the table

		err = results.Scan(&driver.FirstName, &driver.LastName,
			&driver.MobileNo, &driver.Email, &driver.Password, &driver.NRIC, &driver.Carlicense, &driver.Availability)
		if err != nil {
			panic(err.Error())
		}
		DUser = append(DUser, driver)

		for _, v := range DUser {
			int := 0
			if v.Availability == 0 {
				fmt.Println("your driver is " + v.FirstName + v.LastName)
				availdriver = append(availdriver, v)
				EditDriver(db, v.FirstName, v.LastName, v.MobileNo, v.Email, v.Password, v.NRIC, v.Carlicense, 1)
				return v.NRIC
			} else {
				int++
				continue
			}
		}
		if availdriver == nil {
			fmt.Println("there are no available drivers at the moment, please wait")
			return ""
		}
	}
	return ""
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
		"UPDATE Driver SET FirstName='%s', LastName='%s', MobileNo=%d, Email='%s', Password='%s', CarLicense='%s', Availability = %d WHERE NRIC='%s'",
		FN, LN, MNo, Email, Pw, CL, Avail, NRIC)
	_, err := db.Exec(query)
	if err != nil {
		panic(err.Error())

	}
} //EditRecord(db, 2, "Taylor", "Swift", 23)

/*

func GetDriver(db *sql.DB) {
	results, err := db.Query("Select * FROM driver_db.Driver")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var driver Driver
		err = results.Scan(&driver.FirstName, &driver.LastName,
			&driver.MobileNo, &driver.Email, &driver.Password, &driver.NRIC, &driver.Carlicense)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(driver.FirstName, driver.LastName,
			driver.MobileNo, driver.Email, driver.Password, driver.NRIC, driver.Carlicense)
	}
}*/
