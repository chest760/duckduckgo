package main

import (
	"fmt"

	"github.com/chest760/duckduckgo/duckduckgo_search"
	"github.com/chest760/duckduckgo/types"
)

func main() {
	req := types.GetUrls{
		Keyword: "機械学習",
		Region:  "jp-jp",
		Limit:   5,
	}
	res, _ := duckduckgo_search.GetUrls(
		req,
	)

	fmt.Print(res)
}
