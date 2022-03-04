package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("name.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var requestBody bytes.Buffer

	multipartWriter := multipart.NewWriter(&requestBody)
	fileWriter, err := multipartWriter.CreateFormFile("file_field", "name.txt")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		log.Fatalln(err)
	}

	fieldWriter, err := multipartWriter.CreateFormField("file_field")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = fieldWriter.Write([]byte("Value"))
	if err != nil {
		log.Fatalln(err)
	}

	multipartWriter.Close()

	request, err := http.NewRequest(
		"POST", "https://httpbin.org/post", &requestBody,
	)
	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	/*
		requestBody, err := json.Marshal(map[string]string{
			"name":  "Mark Lark Dark",
			"email": "markiz@gmail.com",
		})

		if err != nil {
			log.Fatalln(err)
		}
	*/
	//timeout := time.Duration(5 * time.Second)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
		return
	}
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	/*
		request.Header.Set("Content-type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}



		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
	*/
	log.Println(result)
}
