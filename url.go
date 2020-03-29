package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getData() Country {
	data := Country{}
	resp, err := http.Get("https://www.bing.com/covid/data?IG=1")
	if err != nil {
		fmt.Println("Error", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
	}

	erro := json.Unmarshal([]byte(body), &data)
	if erro != nil {
		log.Fatal(erro)
	}

	return data
}
