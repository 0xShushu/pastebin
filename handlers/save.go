package handlers

import (
	"github.com/0xshushu/pastebin/utils"
	"net/http"
	"os"
)

//save code in the server
func SaveHandler(w http.ResponseWriter, r *http.Request) {
	//generate random name
	name := utils.RandString()
	
	//get post parameters
	r.ParseForm()
	code := r.Form.Get("code")

	//write file
	err := os.WriteFile("code/"+name, []byte(code), 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//redirect to the code
	http.Redirect(w, r, "/code/"+name, http.StatusMovedPermanently)
}