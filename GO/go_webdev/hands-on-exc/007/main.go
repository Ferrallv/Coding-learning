package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type codes []struct {
	Descrip string `json:"Descrip"`
	Code int `json:"Code"`	
}

func main() {
	rcvd := `[{"Descrip":"StatusOK","Code":200},{"Descrip":"StatusMovedPermanently","Code":301},{"Descrip":"StatusFound","Code":302},{"Descrip":"StatusSeeOther","Code":303},{"Descrip":"StatusTemporaryRedirect","Code":307},{"Descrip":"StatusBadRequest","Code":400},{"Descrip":"StatusUnauthorized","Code":401},{"Descrip":"StatusPaymentRequired","Code":402},{"Descrip":"StatusForbidden","Code":403},{"Descrip":"StatusNotFound","Code":404},{"Descrip":"StatusMethodNotAllowed","Code":405},{"Descrip":"StatusTeapot","Code":418},{"Descrip":"StatusInternalServerError","Code":500}]`

	var data codes

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln("Donked the Unmarshalling. Try again Mike Dawson", err)
	}

	for _, v := range data {
		fmt.Println("Description:", v.Descrip)
		fmt.Println("Code:", v.Code)
	}

	// fmt.Print(data)
}