package dto

import (
	"github.com/RomeroGabriel/brandmonitor-challange/backend/pkg/search"
)

type SearchDTO struct {
	SearchString   string `json:"search_string" query:"search_string"`
	LanguageSearch string `json:"language_search,omitempty" query:"language_search"`
}

func (s *SearchDTO) ParseToSearchParameters() search.SearchParameters {
	return search.SearchParameters{SearchString: s.SearchString}
}

// http://localhost:8080/?search_string=Corinthians&language_search=pt
