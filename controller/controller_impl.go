package controller

import (
	"assignment3/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	h8HelperRand "github.com/novalagung/gubrak/v2"
	"gorm.io/gorm"
)

func Update(db *gorm.DB) {
	for {

		water := h8HelperRand.RandomInt(1, 100)
		wind := h8HelperRand.RandomInt(1, 100)

		windStatus := "belum"
		waterStatus := "belum"

		switch {
		case wind <= 6:
			windStatus = "aman"
		case (wind > 6) && (wind <= 15):
			windStatus = "siaga"
		case wind > 15:
			windStatus = "bahaya"
		default:
			{
				windStatus = "not defined"

			}
		}

		switch {
		case water <= 5:
			waterStatus = "aman"
		case (water > 5) && (water <= 15):
			waterStatus = "siaga"
		case water > 15:
			waterStatus = "bahaya"
		default:
			{
				waterStatus = "not defined"
			}
		}

		input := model.Weather{
			Wind:  wind,
			Water: water,
		}

		requestJson, err := json.Marshal(input)
		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest("PUT", "http://localhost:8080/weather",
			bytes.NewBuffer(requestJson))
		req.Header.Set("Content-type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(body))
		fmt.Printf("\nwater : %d\nwind : %d\n", water, wind)
		fmt.Printf("water status : %s\nwind status: %s\n", waterStatus, windStatus)

		time.Sleep(15 * time.Second)
	}

}
