package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/patrickambrosso/nexus-elysium/models"
	"github.com/patrickambrosso/nexus-elysium/pages"
	"github.com/patrickambrosso/nexus-elysium/routes"
)

func main() {
	mux := http.NewServeMux()

	// Page-Handlers - Base Pages
	mux.HandleFunc("GET /", HomePageHandler)
	mux.HandleFunc("GET /references", ReferencesPageHandler)
	mux.HandleFunc("GET /trainings", TrainingsPageHandler)
	mux.HandleFunc("GET /extras", ExtrasPageHandler)

	// Pahe-Handlers - References Section
	mux.HandleFunc("GET /references/company-policy", AboutPageHandler)

	mux.HandleFunc("GET /references/sop", SOPCollectionsPageHandler)
	mux.HandleFunc("GET /references/sop/{id}", SOPReaderPageHandler)

	mux.HandleFunc("GET /references/spec", SOPCollectionsPageHandler)
	mux.HandleFunc("GET /references/spec/{id}", SOPReaderPageHandler)

	mux.HandleFunc("GET /references/stp", SOPCollectionsPageHandler)
	mux.HandleFunc("GET /references/stp/{id}", SOPReaderPageHandler)

	mux.HandleFunc("GET /references/gtp", SOPCollectionsPageHandler)
	mux.HandleFunc("GET /references/gtp/{id}", SOPReaderPageHandler)

	mux.HandleFunc("GET /references/glossary", AboutPageHandler)

	// Pahe-Handlers - Trainings Section
	mux.HandleFunc("GET /trainings/{course_name}", AboutPageHandler)

	// Route Handling
	mux.HandleFunc("GET /actions/sop-documents", SOPRouteHandler)
	mux.HandleFunc("POST /actions/sop-documents/search", SOPSearchHandler)
	mux.HandleFunc("GET /does-it-work", EndpointTester)

	fmt.Println("Server running on http://localhost:4544")
	log.Fatal(http.ListenAndServe(":4544", mux))
}

// Page Handlers - Base Pages
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	err := pages.Base(pages.Home()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ReferencesPageHandler(w http.ResponseWriter, r *http.Request) {
	err := pages.Base(pages.References()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TrainingsPageHandler(w http.ResponseWriter, r *http.Request) {
	err := pages.Base(pages.Trainings()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ExtrasPageHandler(w http.ResponseWriter, r *http.Request) {
	err := pages.Base(pages.Extras()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Page Handlers - References Section
func CompanyPoliciesCollectionsPageHandler(w http.ResponseWriter, r *http.Request) {
	err := pages.Base(pages.SOPCollections()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SOPCollectionsPageHandler(w http.ResponseWriter, r *http.Request) {
	err := pages.Base(pages.SOPCollections()).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SOPReaderPageHandler(w http.ResponseWriter, r *http.Request) {
	// Path to the JSON file
	filePath := "data/sop-docs.json"

	// Read the JSON file
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Create a map to hold the unmarshaled JSON data
	var data map[string]models.SOPDocument

	// Unmarshal the JSON data into the map
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	id := strings.ToLower(r.PathValue("id"))

	if _, exists := data[id]; exists {
		fmt.Printf("%s is in the map\n", id)
		// message := fmt.Sprintf("Found the key %s", id)
		err = pages.SOPReader(data[id]).Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		err = pages.Base(pages.NotFoundError()).Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Page Handlers - Trainings Section

// Page Handlers - Extras Section

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You have visited the About page")
}

// Route Handlers
func SOPRouteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hitting the Path")

	// err := templates.FilteredSOPDocuments().Render(r.Context(), w)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}

func SOPSearchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SOP Search Handler Triggered")

	// Path to the JSON file
	filePath := "data/sop-docs.json"

	// Read the JSON file
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Create a map to hold the unmarshaled JSON data
	var data map[string]models.SOPDocument

	// Unmarshal the JSON data into the map
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get the search query from the form data
	searchDepartment := r.FormValue("department-selection")
	searchQuery := strings.ToLower(r.FormValue("sop-search-query"))

	// Search for matches
	results := make(map[string]models.SOPDocument)

	for id, doc := range data {
		if (searchDepartment == "" || doc.SOPDept == searchDepartment) && strings.Contains(strings.ToLower(doc.SearchIndex), searchQuery) {
			results[id] = doc
		}
	}

	// Print the search query
	fmt.Printf("Client has requested for the query %s from the department %s\n", searchQuery, searchDepartment)

	// Set headers for HTMX response
	w.Header().Set("Content-Type", "text/html")

	// Render the filtered documents
	err = routes.FilteredSOPDocuments(results).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

}

func EndpointTester(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The test endpoint is being reached by a GET Request")
}
