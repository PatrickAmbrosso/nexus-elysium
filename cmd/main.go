package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/patrickambrosso/nexus-elysium/templates"
)

func main() {
	mux := http.NewServeMux()

	// Page Handling
	mux.HandleFunc("GET /", HomePageHandler)
	mux.HandleFunc("GET /references", AboutPageHandler)
	mux.HandleFunc("GET /references/company-policy", AboutPageHandler)
	mux.HandleFunc("GET /references/regulatory-guidelines", AboutPageHandler)
	mux.HandleFunc("GET /references/operational-guidelines", AboutPageHandler)
	mux.HandleFunc("GET /references/glossary", AboutPageHandler)
	mux.HandleFunc("GET /trainings", AboutPageHandler)
	mux.HandleFunc("GET /trainings/{course_name}", AboutPageHandler)
	mux.HandleFunc("GET /extras", AboutPageHandler)

	// Route Handling

	fmt.Println("Server running on http://localhost:4544")
	log.Fatal(http.ListenAndServe(":4544", mux))
}

// Page Handlers
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Base(templates.Home()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ReferencesPageHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Base(templates.Home()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CompanyPolicyPageHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Base(templates.Home()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegulatoryGuidelinesPageHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Base(templates.Home()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func OperationalGuidelinesPageHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Base(templates.Home()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GlossaryPageHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Base(templates.Home()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Base(templates.Page()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Route Handlers
