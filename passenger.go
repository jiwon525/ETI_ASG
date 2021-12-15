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
	var Pssng Passenger
	defer results.Close()
	for results.Next() {
		// map this type to the record in the table
		var passenger Passenger
		err = results.Scan(&passenger.UserName, &passenger.FirstName, &passenger.LastName, &passenger.MobileNo, &passenger.Email, &passenger.Password)
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(passenger.UserName, passenger.FirstName, passenger.LastName, passenger.MobileNo, passenger.Email, passenger.Password)
		PUser = append(PUser, Pssng)
	}
	for _, v := range PUser {
		int := 0
		if v.UserName == UN {
			if PUser[int].Password == Pw {
				fmt.Println("you are logged in.")
				return 1
			} else {
				fmt.Println("user found but password is wrong")
				return 2
			}
		} else {
			int++
			fmt.Println("you have not signed in before, or have entered the wrong username")
			return 3
		}

	}
	return 0
}

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
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

func EditPassenger(db *sql.DB, UN string, FN string, LN string, MNo int, Email string, Pw string) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/user_db")
	query := fmt.Sprintf(
		"UPDATE Persons SET FirstName='%s', LastName='%s', MobileNumber=%d, Email='%s', Password='%s', WHERE UserName='%s'",
		FN, LN, MNo, Email, Pw, UN)
	_, err = db.Exec(query)
	if err != nil {
		panic(err.Error())
	}
} //EditRecord(db, 2, "Taylor", "Swift", 23)
