package routes

import (
	"net/http"
	"fmt"
)

func FormHandler(res http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodPost {
		http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := req.ParseForm(); err != nil{
		fmt.Fprintf(res, "Parseform() err: %v", err)
		return
	}

	fmt.Fprintf(res, "POST request successful")
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(res, "Name = %s\n", name)
	fmt.Fprintf(res, "Address = %s\n", address)
}

