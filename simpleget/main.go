package main

import (
	"log"
	"net/http"

	//   "net/url"
	// "io/ioutil"
	"bytes"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
)

func main() {
	//GET
	// 	values:=url.Values{
	// 		"query":{"hello world"},
	// 	}
	// 	resp,_:=http.Get("http://localhost:18888"+"?"+values.Encode())
	// // if err != nil{
	// // 	panic(err)
	// // }

	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// // if err != nil{
	// // 	panic(err)
	// // }

	//   log.Println(string(body))

	// //HEAD method
	//   values:=url.Values{
	// 	"query":{"hello world"},
	// }
	// resp,err:=http.Head("http://localhost:18888"+"?"+values.Encode())
	// if err != nil{
	// 	panic(err)
	// }
	// log.Println("Status: ",resp.Status)
	// log.Println("Headers: ",resp.Header)

	//POST(multiipart/form-data)
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	// fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}

	// defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp.Status)

}
