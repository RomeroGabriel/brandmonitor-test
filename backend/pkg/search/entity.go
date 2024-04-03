package search

type SearchParameters struct {
	SearchString string
}

type SearchResult struct {
	Index       int
	URL         string
	Title       string
	Description string
}
