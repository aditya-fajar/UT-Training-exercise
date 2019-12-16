package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Umur   string `json:"umur"`
}

func main() {
	//Ganti ^ menjadi main jika mau mencoba

	slice := []Profile{{"Adit", "Bogor", "22"}}
	profile, err := json.Marshal(slice)

	if err != nil {
		fmt.Println("Error Marshalling Json " + err.Error())
		return
	}

	fmt.Println(string(profile))

}
