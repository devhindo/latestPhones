package main

import (
	"fmt"
	"net/http"
	"io"
	
	"os"
	"github.com/joho/godotenv"
)

func main() {

	url := "https://mobile-phone-specs-database.p.rapidapi.com/gsm/all-brands"

	req, _ := http.NewRequest("GET", url, nil)

	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	key := os.Getenv("X_RapidAPI_Key")

	req.Header.Add("X-RapidAPI-Key", key)
	req.Header.Add("X-RapidAPI-Host", "mobile-phone-specs-database.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
