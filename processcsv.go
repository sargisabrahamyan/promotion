package main

import (
    "encoding/csv"
    "os"
    "io"
    "log"
    "strconv"
    "time"
)

func RunProcesscsv(csvName string) string {

    //processCSV("./data/promotions (2).csv")
    return processCSV("./data/" + csvName)
}

var db = OpenConnection()

//var readbuffersize = 100  // took 3.49 sec
var readbuffersize = 1000 // took 2.1 sec
//var readbuffersize = 10000 // took 6.1 sec
//var readbuffersize = 10 // took 17 sec
//var readbuffersize = 1 // took 4 mins
//var readbuffersize = 200000 // took 4 mins
func processCSV(filePath string) string {
    ts := time.Now()
    log.Println("Start processing csv file in path: ", filePath)
    log.Println("Time: ", ts)
    csvFile, err := os.Open(filePath)
    if err != nil {
        log.Println("ERROR: Unable to read input file " + filePath, err)
        return "Error"
    }
    defer csvFile.Close()

    csvReader := csv.NewReader(csvFile)
    if err != nil {
        log.Fatal("Unable to parse file in path " + filePath, err)
    }
    buffer := make([]Promotion, readbuffersize)
    i := 0
    for {
        record, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        p := createPromotion(record)
        buffer[i] = p
        i = i+1
        if i == readbuffersize {
            WritePromotions(buffer)
            buffer = make([]Promotion, readbuffersize) // TODO clean buffer more optinal
            i = 0
        }

    }
    if len(buffer) > 0 {
        WritePromotions(buffer)
    }
    te := time.Now()
    log.Println("End processing csv file at ", te)
    duration := te.Sub(ts)
    log.Println("Insert bulk size : ?, Took: ?", readbuffersize, duration)
    return "Done"
}

func createPromotion(record []string) Promotion {
    price, _ := strconv.ParseFloat(record[1], 64)
    p := Promotion{record[0], price, record[2]}
    return p;
}

func WritePromotions(promotions []Promotion) {
    // TODO nil elements in array
    //log.Println("Write Promotions of size:", len(promotions))
    InsertPromotions(promotions, db)
}

