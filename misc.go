package httptools

import (
	"net/http"
	"strings"
)

// DiscardPathElements discards n elements from the request path.
// It's most useful in a handler list.
func DiscardPathElements(n int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		elems := strings.Split(r.URL.Path[1:], "/")
		if n >= len(elems) {
			r.URL.Path = "/"
			return
		}
		r.URL.Path = "/" + strings.Join(elems[n:], "/")
	})
}
