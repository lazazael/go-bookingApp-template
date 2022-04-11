package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

/*func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("pagehit")
		next.ServeHTTP(w, r)
	})
}
*/

//NoSurf adds CSRF protection to every POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

//SessionStateLoad loads and saves session on every request
func SessionStateLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
