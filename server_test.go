package serverfolderselect_test

import (
	"net/http"
	"strings"
	"sync"

	serverfolderselect "github.com/PetaTookmyKFC/Server_Folder_Select"
)

var wg sync.WaitGroup
var fileSystem *serverfolderselect.FileSelection

// Closes the server if the url is /close. This is also used as an example of how the prechecks can work
// A return statment of false should be returned if the request is not to continue, or if the precheck has already returned a responce to the user.
func CloseWithURL(w http.ResponseWriter, r *http.Request) bool {
	// Split the path into an array
	pathURL := strings.Split(r.URL.Path, "/")
	for _, pA := range pathURL {
		// Check if the world close is a directory in the url
		if strings.ToUpper(pA) == "CLOSE" {
			// Send responce to user
			w.WriteHeader(412)
			w.Write([]byte("Goodbye World!"))
			wg.Done()
			// Return false to system * Not required in this example as the waitgroup is completed*
			return false
			// False prevents the filesystem from sending the responce
		}
	}
	// True informs the api that no isues where found with THIS precheck - Other prechecks are not affected by this.
	// If another precheck returns false the next checks are not checked.
	return true
}

// Simple Function to run the server
func main() {
	// Create a waitgroup to prevent blocking of main thread
	wg.Add(1)
	// Inform the API what the base route is
	fileSystem = serverfolderselect.CreateFSApi("/", "/")
	// Add the preCheck example
	fileSystem.AppendPreCheck(CloseWithURL)
	// Start the http server
	go http.ListenAndServe(":9090", nil)
	// Wait for the waitgroup to finnish
	wg.Wait()
}
