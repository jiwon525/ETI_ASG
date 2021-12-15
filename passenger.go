package main

import ( //passenger
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Passenger struct {
	UserName  string
	FirstName string
	LastName  string
	MobileNo  int
	Email     string
	Password  string
}

func CheckLogin(db *sql.DB, PUser []Passenger, UN string, Pw string) int { //for user verification
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/user_db")
	results, err := db.Query("SELECT * FROM user_db.Passenger")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		// map this type to the record in the table
		var passenger Passenger
		err = results.Scan(&passenger.UserName, &passenger.FirstName, &passenger.LastName, &passenger.MobileNo, &passenger.Email, &passenger.Password)
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(passenger.UserName, passenger.FirstName, passenger.LastName, passenger.MobileNo, passenger.Email, passenger.Password) was used to check if the data is correct

		PUser = append(PUser, passenger)
	}
	//fmt.Println(PUser) was used to check if data was correctly appended to list
	for _, v := range PUser {
		int := 0
		if v.UserName == UN {
			if v.Password == Pw {
				fmt.Println("you are logged in.")
				return 1
			} else {
				fmt.Println("user found but password is wrong")
				return 2
			}
		} else {
			int++
		}

	}
	fmt.Println("you have not signed up before, or have entered the wrong username")
	return 3
}

func NewPassenger(db *sql.DB, UN string, FN string, LN string, MNo int, Email string, Pw string) {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/user_db")

	sqlStatement2 := fmt.Sprintf("INSERT INTO Passenger (UserName, FirstName, LastName, MobileNo, Email, Password) VALUES ('%s','%s','%s',%d,'%s','%s')", UN, FN, LN, MNo, Email, Pw)
	//sqlStatement1 := UN+ ","+FN+","+ LN+","+ MNo+"," +Email+"," +Pw
	/*sqlStatement := `
	INSERT INTO Passenger (UserName, FirstName, LastName, MobileNo, Email, Password)
	VALUES ("Jw","jiwon","jung", 12341234, "sand@gmail.com","pass11")`*/
	//MNo2 := strconv.Itoa(MNo)
	_, err = db.Exec(sqlStatement2)
	if err != nil {
		panic(err)
	}

}

func EditPassenger(db *sql.DB, UN string, FN string, LN string, MNo int, Email string, Pw string) { //need to edit
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/user_db")
	sqlStatement2 := fmt.Sprintf("UPDATE Persons SET FirstName='%s', LastName='%s', MobileNo=%d, Email='%s', Password='%s', WHERE UserName='%s'", FN, LN, MNo, Email, Pw, UN)
	_, err = db.Exec(sqlStatement2)
	if err != nil {
		panic(err)
	}

} //EditRecord(db, 2, "Taylor", "Swift", 23)

func BookRide(db *sql.DB, CL int, DL int, UN string, DI string) { //book a ride with username, open the database for trip, create new trip with the first driver in the list that is available
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db")
	if err != nil {
		panic(err)
	}
	NewTrip(db, CL, DL, UN, DI)

}
