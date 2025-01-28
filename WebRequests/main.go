package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myUrl = "https://api.restful-api.dev/objects?id=3"

func main() {
	fmt.Println("Welcome to Web Request Tutorial")

	// webRequests()
	// urls()
	performPostJsonRequest()
	// performPostFormRequest()

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func webRequests() {
	response, err := http.Get(myUrl)
	checkNilErr(err)

	defer response.Body.Close()

	var responseString strings.Builder
	dataBytes, err := io.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Print(string(dataBytes)) // One way

	byteCount, _ := responseString.Write(dataBytes)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String()) // Second way
}

func urls() {
	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["id"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}

func performPostJsonRequest() {
	const myUrl = "https://api.restful-api.dev/objects"

	requestBody := strings.NewReader(`
	 	{
   			"name": "Apple MacBook Pro 16",
   			"data": {
     		 	"year": 2019,
      			"price": 1849.99,
      			"CPU model": "Intel Core i9",
      			"Hard disk size": "1 TB"
   			}
		}
	 `)

	response, err := http.Post(myUrl, "application/json", requestBody)
	checkNilErr(err)

	defer response.Body.Close()

	dataBytes, _ := io.ReadAll(response.Body)

	println(string(dataBytes))

}

func performPostFormRequest() {
	const myUrl = "https://api.restful-api.dev/objects"

	data := url.Values{}
	data.Add("name", "Sajib")
	data.Add("data[year]", "2025")
	data.Add("data[price]", "100")
	data.Add("data[CPU model]", "Intel Core i9")
	data.Add("data[Hard disk size]", "1 TB")

	response, err := http.PostForm(myUrl, data)
	checkNilErr(err)

	defer response.Body.Close()

	dataBytes, _ := io.ReadAll(response.Body)

	println(string(dataBytes))

}
