package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"log"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

type profile struct {
	id int
	eid string
	ename string
	eemail string
	econtact string
}

func init(){
	connStr := "user=postgres password=mypassword host=localhost port=5432 dbname=crud_api sslmode=disable"
	db,err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Now we are connected to POSTGRESQL DATABASE")
}

func main(){
	http.HandleFunc("/data", dataRecord)
	http.ListenAndServe(":8000", nil)
}

func dataRecord(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Great!!! We are connected to the browser\n")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(404), http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query("SELECT * FROM employee")

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	emps := make([]profile,0)

	for rows.Next(){
		emp := profile{}
		err := rows.Scan(&emp.id,&emp.eid,&emp.ename,&emp.eemail,&emp.econtact)
		if err != nil{
			log.Println(err)
			http.Error(w, http.StatusText(500),500)
			return
		}
		emps = append(emps, emp)
		for _,emp := range emps {
			fmt.Fprintf(w,"%d, %s,%s,%s,%s", emp.id,emp.eid,emp.ename,emp.eemail,emp.econtact)
		}
	}
}