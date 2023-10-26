package serverfolderselect

import (
	"fmt"
	"net/http"
)

// type PreCheckFunc

type FileSelection struct {
	StartLocation    string
	ApiRoute         string
	BlackList        []string
	BlackListTrigger func(http.ResponseWriter, *http.Request, string)

	PreCheck []func(w http.ResponseWriter, r *http.Request) bool

	// postCheck func(w http.ResponseWriter, r *http.Request) (shouldContinue bool)
}

func (fl *FileSelection) AppendPreCheck(prcheck func(w http.ResponseWriter, r *http.Request) bool) {
	fl.PreCheck = append(fl.PreCheck, prcheck)
	// a := append(fl.PreCheck, prcheck)
	// fl.PreCheck = a
}

// Run Every precheck function
func (fl *FileSelection) RunPreChecks(w http.ResponseWriter, r *http.Request) bool {
	fmt.Println("Running PreChecks")

	// Loop through registered prechecks
	responce := true
	for _, value := range fl.PreCheck {
		responce = value(w, r)
		if !responce {
			fmt.Println("Running PreChecks! Got False")
			return false
		} else {
			fmt.Println("Running PreChecks! Got True")
		}
	}
	return true
}
