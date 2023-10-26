
```go
type FileSelection struct {
	StartLocation string
	BlackList []string
	PreCheck []func(http.ResponseWriter, *http.Request) ( bool )
	BackListTrigger func(http.ResponceWriter, *http.Request, string)
}

```

`ApiRoute` - the route that will be used on the web for the client to connect to the api. This will have a wildcard at the end of this route so no routes should contain this path.
`StartLocation` - the initial folder the page should load ( if no page is passed )
`BlackList` - An Array of File paths that will trigger the blacklist function and result in no response  returned to the user. A response can be returned to the user but this must be set in the `BlackListTrigger`
`BlackListTrigger` this is the function that is triggered if a blacklist item is found. For Example a black list item may be 'C:/system32'. If the system finds this in the path the system will trigger this function and not automatically send a response.
`PreCheck` This is a function that the developer can insert. This can allow the checking of state and any function to be run before allowing the system to continue. For example a key or password could be checked or a log function. After the check is run a bool should be returned, if false is returned the system will stop handling the request assuming that the precheck has completed it. The order of the prechecks is that they appear in the array. All the prechecks will be run before a response is given from the server. As a pointer to the request is passed to the precheck the request may be modified. All prechecks will be run before responding to the user. Once a precheck returns false the system will not continue to run anymore of the checks. This will allow the user to add custom security measures to the system without the need to edit the code directly.

The route will be resisted with `http.Handlefunc` but the server will not be started by this code.
## Getting Started



