// https://blog.logrocket.com/making-http-requests-in-go/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	People  []struct {
		Craft string `json:"craft"`
		Name  string `json:"name"`
	} `json:"people"`
	Number int `json:"number"`
}

func main() {

	// Get request
	resp, err := http.Get("http://api.open-notify.org/astros.json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// fmt.Println(PrettyPrint(result))
	//fmt.Println(result)
	// Loop through the data node for the FirstName
	for _, rec := range result.People {
		fmt.Println(rec.Name)
		fmt.Println(rec.Craft)
	}

	fmt.Println(result.Number)
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
