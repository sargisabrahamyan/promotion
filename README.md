# promotion

## This project designed for Parsing Saving and Serving the given CSV file with predefined structure content

## Requirements
	1. User must have database with the following credentials :
	  	host     = "localhost"
  		port     = 5432
  		user     = "postgres"
  		password = "demo_pass"
  		dbname   = "go_db"
  	2. User must create table 'promotion' in the 'go_db'

## To run the project enter the project folder and type `go run .`
### User has two main endpoint to acces :
#### http://localhost:1321/server/readcsv/<fileName\> - this will read given csv file and store in DB
	Note : the fileName is the csv file which should be located under 'data' folder of the project
#### http://localhost:1321/promotions/<uuid\> - this will find promotion object in the DB.
