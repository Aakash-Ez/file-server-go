package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
    "os"
	"strings"
)

func HandlePut(rw http.ResponseWriter, req *http.Request) {

    file, handler, err := req.FormFile("File")
	fileName := req.FormValue("FileName")

	extensionList := strings.Split(handler.Filename, ".")
	extension := extensionList[len(extensionList) - 1]

    if fileName == "" {
        panic("Input not Recieved")
    }

    if err != nil {
        fmt.Println("File not Recieved")
        fmt.Println(err)
        return
    }

    defer file.Close()

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    
    err = ioutil.WriteFile("temp/"+ fileName + "." + extension, []byte(fileBytes), 0755)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Fprint(rw, "File uploaded successfully\n")
}

func HandleGet(rw http.ResponseWriter, req *http.Request) {

    fileName := req.FormValue("FileName")

    if fileName == "" {
        fmt.Fprint(rw, "Input not Recieved")
        fmt.Println("Input not Recieved")
        return
    }

	fileBytes, err := ioutil.ReadFile("temp/" + fileName)
	if err != nil {
        fmt.Fprint(rw, err.Error())
        fmt.Println(err.Error())
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/octet-stream")
	rw.Write(fileBytes)
}

func HandleDelete(rw http.ResponseWriter, req *http.Request) {
    
    fileName := req.FormValue("FileName")

    if fileName == "" {
        fmt.Fprint(rw, "Input not Recieved")
        fmt.Println("Input not Recieved")
        return
    }

    err := os.Remove("temp/" + fileName)
    if err != nil {
        fmt.Fprint(rw, err.Error())
        fmt.Println(err.Error())
		return
	}

    fmt.Fprint(rw, "File Deleted")

}

func HandleList(rw http.ResponseWriter, req *http.Request) {

    output := ""

    files, err :=  os.ReadDir("temp/")

    if err != nil {
        fmt.Fprint(rw, err.Error())
        fmt.Println(err.Error())
		return
	}


    for i := range files {
        if !files[i].IsDir() {
            output = output + files[i].Name() + "\n"
        }
    }

    fmt.Fprint(rw, output)
}