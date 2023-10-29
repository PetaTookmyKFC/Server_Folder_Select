package serverfolderselect

import (
	"fmt"
	"net/http"
	"text/template"
)

func SendFile(fileLoc string, w http.ResponseWriter, r *http.Request) {

	// wd, err := os.Getwd()
	// if err != nil {
	// 	log.Panic(err)
	// }

	fileLoc = fmt.Sprintf("./Components/%s", fileLoc)
	// fileLoc, err := filepath.Abs(fileLoc)
	// if err != nil {
	// 	w.WriteHeader(500)
	// 	w.Write([]byte("NOT WORKING :SAD:"))
	// }
	temp, err := template.ParseFiles(fileLoc)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("NOT WORKING :SAD:"))
	} else {
		temp.Execute(w, nil)
	}
}
