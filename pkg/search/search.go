package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SearchResponse struct {
	Kind              string            `json:"kind"`
	URL               URL               `json:"url"`
	Queries           Queries           `json:"queries"`
	Context           Context           `json:"context"`
	SearchInformation SearchInformation `json:"searchInformation"`
	Items             []Item            `json:"items"`
}

type URL struct {
	Type     string `json:"type"`
	Template string `json:"template"`
}

type Queries struct {
	Request  []Request `json:"request"`
	NextPage []Request `json:"nextPage"`
}

type Request struct {
	Title          string `json:"title"`
	TotalResults   string `json:"totalResults"`
	SearchTerms    string `json:"searchTerms"`
	Count          int    `json:"count"`
	StartIndex     int    `json:"startIndex"`
	InputEncoding  string `json:"inputEncoding"`
	OutputEncoding string `json:"outputEncoding"`
	Safe           string `json:"safe"`
	Cx             string `json:"cx"`
}

type Context struct {
	Title string `json:"title"`
}

type SearchInformation struct {
	SearchTime            float64 `json:"searchTime"`
	FormattedSearchTime   string  `json:"formattedSearchTime"`
	TotalResults          string  `json:"totalResults"`
	FormattedTotalResults string  `json:"formattedTotalResults"`
}

type Item struct {
	Kind        string `json:"kind"`
	Title       string `json:"title"`
	HTMLTitle   string `json:"htmlTitle"`
	Link        string `json:"link"`
	DisplayLink string `json:"displayLink"`
	Snippet     string `json:"snippet"`
	HTMLSnippet string `json:"htmlSnippet"`
	Mime        string `json:"mime"`
	FileFormat  string `json:"fileFormat"`
	Image       Image  `json:"image"`
}

type Image struct {
	ContextLink     string `json:"contextLink"`
	Height          int    `json:"height"`
	Width           int    `json:"width"`
	ByteSize        int    `json:"byteSize"`
	ThumbnailLink   string `json:"thumbnailLink"`
	ThumbnailHeight int    `json:"thumbnailHeight"`
	ThumbnailWidth  int    `json:"thumbnailWidth"`
}

type Searcher struct {
	Key      string
	EngineID string
	BaseURL  string
}

const baseURL = "https://www.googleapis.com/customsearch/v1"

func NewSearcher(key, engineID string) *Searcher {
	return &Searcher{
		Key:      key,
		EngineID: engineID,
		BaseURL:  baseURL,
	}
}

func (s *Searcher) SearchImages(query string, start int, transparent bool) (*SearchResponse, error) {
	u, err := url.Parse(s.BaseURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("key", s.Key)
	q.Set("cx", s.EngineID)
	q.Set("q", query)
	q.Set("filter", "1")
	q.Set("searchType", "image")
	q.Set("start", fmt.Sprintf("%d", start))
	q.Set("excludeTerms", "buy kaufen shop store bestellen product produkt logo sale discount angebot preis shopping ecommerce katalog kita kindergarten organisation organization öl oil medicine medizin heilung heilmittel verpackung schule rätsel")
	q.Set("num", "10")
	if transparent {
		q.Set("imgColorType", "trans")
		q.Set("fileType", "png")
		q.Set("q", query)
	}

	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var sr SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return nil, err
	}

	return &sr, nil
}
