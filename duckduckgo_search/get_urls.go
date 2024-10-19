package duckduckgo_search

import (
	"github.com/anaskhan96/soup"
	"github.com/chest760/duckduckgo/types"
)

func GetUrls(search types.GetUrls) ([]types.ResponseUrls, error) {
	url := "https://html.duckduckgo.com/html"
	params := map[string]interface{}{
		"q":  search.Keyword,
		"kl": search.Region,
		"df": string(search.TimeRange),
	}

	res, err := Request("GET", url, params, nil)
	if err != nil {
		return []types.ResponseUrls{}, err
	}

	result, err := parse_html(string(*res))
	if err != nil {
		return []types.ResponseUrls{}, err
	}
	return result[:search.Limit], nil
}

func parse_html(html string) ([]types.ResponseUrls, error) {

	response := []types.ResponseUrls{}

	doc := soup.HTMLParse(html)
	results := doc.FindAll("div", "class", "result__body")
	for _, result := range results {
		title := result.Find("a", "class", "result__a").FullText()
		url := result.Find("a", "class", "result__a").Attrs()["href"][2:]
		response = append(response, types.ResponseUrls{
			Title: title,
			URL:   url,
		})
	}
	return response, nil
}
