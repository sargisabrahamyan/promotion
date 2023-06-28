package main

import (
	"fmt"
	"log"
	"database/sql"
	 _ "github.com/lib/pq"
	 "strconv"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "demo_pass"
  dbname   = "go_db"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
  		panic(err)
  		log.Fatal("Cannot insert Promotion: %v", err)
	}
	// _, err2 := db.Exec("CREATE TABLE IF NOT EXISTS promoton (id UUID NOT NULL PRIMARY KEY, price NUMBER NOT NULL, )")
    // if err2 != nil {
    //     panic(err2)
    // }
	return db
}

func InsertPromotion(p Promotion, db *sql.DB) {

	result, err := db.Exec(`INSERT INTO promotion(id, price, expiration_date) VALUES ($1, $2, $3);`, p.Id,p.Price, p.Exp_date)
    if err != nil {
    	log.Fatal("Cannot insert Promotion: ", p, " Error: ", err)
        return
    }
    log.Println("Inserted Promotion: %v", result)
}


func InsertPromotions(promotions [] Promotion, db *sql.DB) {
	sqlStr := "INSERT INTO promotion(id, price, expiration_date) VALUES "

	for _, promotion := range promotions {
		if promotion.Id == "" { // TODO we have empty elements in buffer
			continue
		}
	    sqlStr += "( '" + promotion.Id + "', " + strconv.FormatFloat(promotion.Price,'g', 10, 64) + ", '" + promotion.Exp_date + "' ),"
	}
	sqlStr = sqlStr[0:len(sqlStr)-1]
	stmt, err := db.Prepare(sqlStr)

	if err != nil {
		log.Fatal("Error inserting :", err)
		return
	}
	log.Println("Execute build insert")
	_, err2 := stmt.Exec()
		if err2 != nil {
		log.Fatal("Error inserting :", err2)
		return
	}
}

func GetPromotion(id string, db *sql.DB) *Promotion {
	var id_ string
	var price float64
	var date string

	row := db.QueryRow(`SELECT * FROM promotion WHERE id = $1 ;`, id)
	if row == nil {
		log.Fatal("DB ERROR")
		return nil
	}
	err := row.Scan(&id_, &price, &date)
    switch err {
		case sql.ErrNoRows:
	  		log.Println("No rows were returned!")
	  		return nil
		case nil:
	  		log.Println("Found row by id: ", id)
		default:
	  		panic(err)
  	}
  	p := &Promotion{id_, price, date}
	return p
}

