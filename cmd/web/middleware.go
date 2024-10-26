package main

import "net/http"

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// style can be fetched from self and fonts.googleapis.com.
		// fonts can be fetched from self and fonts.gstatic.com
		// everything else must be fetched from self.
		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")

		// If referred to same origin, entire URL is included in the referred header.
		// If referred to a different origin, URL path and query string parameters are stripped off.
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")

		// Instruct the browser not to sniff the content
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Disable webpage from being loaded as a fram/iframe/object
		w.Header().Set("X-Frame-Options", "deny")

		// Disable XSS protection. XSS protection is taken cary by the CSP header set above.
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}
