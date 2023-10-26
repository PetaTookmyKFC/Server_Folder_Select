package serverfolderselect

import (
	"net/http"
	"strings"
)

func (fl *FileSelection) ApiSwitch(w http.ResponseWriter, r *http.Request) {
	// Get File Array
	pathURL := strings.Split(r.URL.Path, "/")

	// pstring, _ := json.Marshal(pathURL)
	// w.Write([]byte(pstring))
	// Remove first element if empty
	if pathURL[0] == "" && len(pathURL) > 1 {
		pathURL = pathURL[1:]
	}

	rApi := pathURL[0]
	// Main Switch Statment
	rApi = strings.ToUpper(rApi)

	switch rApi {

	default:
		{
			SendFile("browser.html", w, r)
		}
	}

}
