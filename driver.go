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
	Availability bool
	DriverID     int
}

func ScanDriverDB(db *sql.DB, DUser []Driver) []Driver {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
	results, err := db.Query("SELECT * FROM driver_db.Driver")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		// map this type to the record in the table
		var driver Driver
		err = results.Scan(&driver.DriverID, &driver.FirstName, &driver.LastName,
			&driver.MobileNo, &driver.Email, &driver.Password, &driver.NRIC, &driver.Carlicense, &driver.Availability)
		if err != nil {
			panic(err.Error())
		}
		DUser = append(DUser, driver)
	}
	return DUser
}
func CheckDriver(db *sql.DB, DUser []Driver, nric string, Pw string) int { //for driver verification
	DriverList := ScanDriverDB(db, DUser)
	//fmt.Println(DUser) was used to check if data was correctly appended to list
	for _, v := range DriverList {
		if v.NRIC == nric {
			if v.Password == Pw {
				fmt.Println("you are logged in.")
				return 1
			} else {
				fmt.Println("driver NRIC found but password is wrong")
				return 2
			}
		}
	}
	fmt.Println("you have not signed up before, or have entered the wrong username")
	return 3
}

func GetAvailDriver(db *sql.DB, DUser []Driver) int {
	DriverList := ScanDriverDB(db, DUser)
	var availdriver []Driver
	for _, v := range DriverList {
		if v.Availability == false {
			fmt.Println("your driver is " + v.FirstName + v.LastName)
			availdriver = append(availdriver, v)
			avail := true
			EditDriver(db, v.DriverID, v.FirstName, v.LastName, v.MobileNo, v.Email, v.Password, v.NRIC, v.Carlicense, avail)
			return v.DriverID
		}
	}
	if availdriver == nil {
		fmt.Println("there are no available drivers at the moment, please wait")
		return 0
	}
	return 0
}

func NewDriver(db *sql.DB, FN string, LN string, MNo int, Email string, Pw string, NRIC string, CL string, Avail bool) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
	/*var numb int
	if Avail == true {
		numb = 1
	} else {
		numb = 0
	}*/
	sqlStatement := fmt.Sprintf("INSERT INTO Driver (FirstName, LastName, MobileNo, Email, Password, NRIC, Carlicense, Availability) VALUES ('%s', '%s', %d , '%s', '%s', '%s', '%s', %t)", FN, LN, MNo, Email, Pw, NRIC, CL, Avail)

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func EditDriver(db *sql.DB, ID int, FN string, LN string, MNo int, Email string, Pw string, NRIC string, CL string, Avail bool) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
	sqlStatement2 := fmt.Sprintf("UPDATE Driver SET FirstName='%s', LastName='%s', MobileNo=%d, Email='%s', Password='%s', NRIC='%s', Carlicense='%s', Availability=%t WHERE DriverID = %d", FN, LN, MNo, Email, Pw, NRIC, CL, Avail, ID)
	_, err = db.Exec(sqlStatement2)
	if err != nil {
		panic(err)
	}

}

func FindDriverID(db *sql.DB, NRIC string, DUser []Driver) int {
	DriverList := ScanDriverDB(db, DUser)
	for _, v := range DriverList {
		if v.NRIC == NRIC {
			return v.DriverID
		}
	}
	return 0
}
