package main

import ( //passenger
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Passenger struct {
	PassengerID int
	UserName    string
	FirstName   string
	LastName    string
	MobileNo    int
	Email       string
	Password    string
}

//var mysqlErr *mysql.MySQLError
func FindPassengerID(db *sql.DB, UN string, PUser []Passenger) int {
	UserList := ScanPassengerDB(db, PUser)
	for _, v := range UserList {
		if v.UserName == UN {
			//fmt.Print(v.PassengerID) was used to check if the passenger id was correctly formatted
			return v.PassengerID
		}
	}
	return 0
}

func ScanPassengerDB(db *sql.DB, PUser []Passenger) []Passenger {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/user_db")
	results, err := db.Query("SELECT * FROM user_db.Passenger")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	for results.Next() {
		// map this type to the record in the table
		var passenger Passenger
		err = results.Scan(&passenger.PassengerID, &passenger.UserName, &passenger.FirstName, &passenger.LastName, &passenger.MobileNo, &passenger.Email, &passenger.Password)
		if err != nil {
			panic(err.Error())
		}

		PUser = append(PUser, passenger)
	}
	return PUser
}

func CheckLogin(db *sql.DB, PUser []Passenger, UN string, Pw string) int { //for user verification
	var UserList []Passenger
	UserList = ScanPassengerDB(db, PUser)
	for _, v := range UserList {
		if v.UserName == UN {
			if v.Password == Pw {
				fmt.Println("you are logged in.")
				return 1
			} else {
				fmt.Println("user found but password is wrong")
				return 2
			}
		}
	}
	fmt.Println("you have not signed up before, or have entered the wrong username")
	return 3
}

func NewPassenger(db *sql.DB, UN string, FN string, LN string, MNo int, Email string, Pw string) {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/user_db")

	sqlStatement2 := fmt.Sprintf("INSERT INTO Passenger (UserName, FirstName, LastName, MobileNo, Email, Password) VALUES ('%s','%s','%s',%d,'%s','%s')", UN, FN, LN, MNo, Email, Pw)

	_, err = db.Exec(sqlStatement2)
	/*if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		fmt.Println("duplicate error, try again")
	}*/
	if err != nil {
		panic(err)
	}

}

func EditPassenger(db *sql.DB, UN string, FN string, LN string, MNo int, Email string, Pw string, ID int) { //need to edit
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/user_db")
	sqlStatement2 := fmt.Sprintf("UPDATE Passenger SET UserName='%s', FirstName='%s', LastName='%s', MobileNo=%d, Email='%s', Password='%s' WHERE PassengerID = %d", UN, FN, LN, MNo, Email, Pw, ID)
	_, err = db.Exec(sqlStatement2)
	if err != nil {
		panic(err)
	}
	Mno := strconv.Itoa(MNo)
	fmt.Println("New details:" + UN + " " + FN + " " + LN + "\n" + Mno + " " + Email + " " + Pw)
}

func BookRide(db *sql.DB, CL int, DL int, PID int, DID int, UN string) { //book a ride with username, open the database for trip, create new trip with the first driver in the list that is available
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db")
	if err != nil {
		panic(err)
	}
	NewTrip(db, CL, DL, PID, DID, UN)

}

//sqlStatement1 := UN+ ","+FN+","+ LN+","+ MNo+"," +Email+"," +Pw
/*sqlStatement := `
INSERT INTO Passenger (UserName, FirstName, LastName, MobileNo, Email, Password)
VALUES ("Jw","jiwon","jung", 12341234, "sand@gmail.com","pass11")`*/
//MNo2 := strconv.Itoa(MNo)
