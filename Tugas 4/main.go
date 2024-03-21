package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.GET("/", handleRequest)
	router.Run()
}

func handleRequest(ctx *gin.Context) {
	status := generateStatus()
	saveStatusToFile(status)

	data := getStatusFromFile()

	waterStatus := getStatusLabel(data.Water, 8, 6)
	windStatus := getStatusLabel(data.Wind, 15, 7)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"Water":       data.Water,
		"WaterStatus": waterStatus,
		"Wind":        data.Wind,
		"WindStatus":  windStatus,
	})
}

func generateStatus() Status {
	rand.Seed(time.Now().UnixNano())
	return Status{
		Water: rand.Intn(100),
		Wind:  rand.Intn(100),
	}
}

func saveStatusToFile(status Status) {
	file, err := os.Create("status.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonData, err := json.Marshal(status)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(string(jsonData))
	if err != nil {
		panic(err)
	}
}

func getStatusFromFile() Status {
	dataFile, err := os.ReadFile("status.json")
	if err != nil {
		panic(err)
	}

	var status Status
	err = json.Unmarshal(dataFile, &status)
	if err != nil {
		panic(err)
	}

	return status
}

func getStatusLabel(value, highThreshold, mediumThreshold int) string {
	if value > highThreshold {
		return "bahaya"
	} else if value >= mediumThreshold {
		return "siaga"
	}
	return "aman"
}
