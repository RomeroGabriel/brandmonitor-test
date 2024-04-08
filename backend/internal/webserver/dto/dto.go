package dto

import (
	"strings"

	"github.com/RomeroGabriel/brandmonitor-challange/backend/pkg/search"
)

type SearchDTO struct {
	SearchString   string `json:"search_string" query:"search_string"`
	LanguageSearch string `json:"language_search,omitempty" query:"language_search"`
}

func (s *SearchDTO) ParseToSearchParameters() search.SearchParameters {
	queryParsed := strings.Replace(s.SearchString, " ", "-", -1)
	return search.SearchParameters{SearchString: queryParsed}
}
