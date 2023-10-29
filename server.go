package serverfolderselect

import (
	"fmt"
	"net/http"
)

func CreateFSApi(ApiRoute string, startLocation string) *FileSelection {
	fileSystem := &FileSelection{
		ApiRoute:      ApiRoute,
		StartLocation: startLocation,
		BlackList:     make([]string, 0),
		PreCheck:      make([]func(http.ResponseWriter, *http.Request) bool, 0),
	}

	http.HandleFunc(ApiRoute, func(w http.ResponseWriter, r *http.Request) {
		CreateHandler(fileSystem)(w, r)
	})

	return fileSystem
}

func CreateHandler(FS *FileSelection) func(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Creating Handler")
	ShowDirectory := func(w http.ResponseWriter, r *http.Request) {
		// Run the registered prechecks
		preChecks := FS.RunPreChecks(w, r)
		// If the prechecks have ran successfully continue
		if preChecks {
			// Pass request to the API switch to filter the request to the desired function
			FS.ApiSwitch(w, r)
		} else {
			// Done the prechecks until one filed, so not doing anything else.
			return
		}

	}
	return ShowDirectory

}

// Testing Systems
