package types

type TimeRange string

const (
	TimeRangeD TimeRange = "d"
	TimeRangeW TimeRange = "w"
	TimeRangeS TimeRange = "m"
	TimeRangeY TimeRange = "y"
)

type GetUrlsParams struct {
	Query     string    `json:"query"`
	Region    string    `json:"region,omitempty"`
	Limit     int       `json:"limit,omitempty"`
	TimeRange TimeRange `json:"time_range,omitempty"`
}

type GetUrls struct {
	Keyword   string    `json:"keyword,omitempty"`
	Region    string    `json:"region,omitempty"`
	Limit     int       `json:"limit,omitempty"`
	TimeRange TimeRange `json:"time_range,omitempty"`
}

type ResponseUrls struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}
