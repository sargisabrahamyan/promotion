package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "encoding/json"
    "log"
)

func main() {
    router := gin.Default()
    router.GET("/promotions/:id", getPromotionById)
    router.GET("/server/readcsv/:name", readCsv)
    router.Run("localhost:1321")
}

var dbreader = OpenConnection()
func readCsv(c *gin.Context) {
    name := c.Param("name")
    log.Println("Given csv file by name :", name)
    result := RunProcesscsv(name)
    c.IndentedJSON(http.StatusOK, result)
}

func getPromotionById(c *gin.Context) {
    id := c.Param("id")
    p := findPromotionById(id)
    p_str := "Promotion not found"
    if p != nil {
        p_str = promotionToString(*p)
    }
    c.IndentedJSON(http.StatusOK, p_str)
}

func findPromotionById(id string) *Promotion {
    fmt.Println("Finding Promition by ID ", id)
    p := GetPromotion(id, dbreader)
    return p
}

func promotionToString(p Promotion) string {
    p_bytes, err := json.Marshal(p)
    if err != nil {
        log.Fatal(err)
        return ""
    }
    p_str := fmt.Sprintf(string(p_bytes))
    log.Println(p_str)
    return p_str
}
