Package httptools tries to augment the basic net/http package with
functionality found in webframeworks without breaking the original API.

For details and examples, please see the [documentation](http://godoc.org/github.com/surma/httptools).

## Contrived example

```Go
r := httptools.NewRegexpSwitch(map[string]http.Handler{
	"/people/(.+)": httptools.L{
		httptools.SilentHandler(AuthenticationHandler),
		httptools.MethodSwitch{
			"GET": ListPeople,
			"PUT": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				vars := w.(httptools.VarsResponseWriter).Vars()
				AddNewPerson(vars["1"])
			})
		},
		SaveSessionHandler,
	},
	"/.+": http.FileServer(http.Dir("./static")),
})
http.ListenAndServe("localhost:8080", r)
```

## Tools
httptools provides the following tools:
### Handler list
Define a sequence of `http.Handler`. One will be executed after another. A
customized `http.ResponseWriter` allows the passing of data in between handlers.
### Silent handler
If a silent handler produces output, it is assumed to be an error. If the
silent handler is in a handler list, the execution of the list will be aborted.
### Method switch
Dispatch requests to different handlers according the the HTTP verb used
in the request.
### RegexpSwitch
Dispatch requests to different handlers according to regexps being matched
agains the request path.


---
Version 1.0.0
