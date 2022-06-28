package advisor

type thumbnailLarge struct {
	Id                int64                `json:"id"`
	Name              string               `json:"name"`
	AlternativeText   string               `json:"alternativeText"`
	Caption           string               `json:"caption"`
	Width             int64                `json:"width"`
	Height            int64                `json:"height"`
	Formats           thumbnailLargeFormat `json:"formats"`
	Hash              string               `json:"hash"`
	Ext               string               `json:"ext"`
	Mime              string               `json:"mime"`
	Size              float64              `json:"size"`
	Url               string               `json:"url"`
	PreviewUrl        string               `json:"previewUrl"`
	Provider          string               `json:"provider"`
	Provider_metadata string               `json:"provider_metadata"`
	CreatedAt         string               `json:"createdAt"`
	UpdatedAt         string               `json:"updatedAt"`
}

type thumbnailLargeFormat struct {
	Large     formatAttribute `json:"large"`
	Small     formatAttribute `json:"small"`
	Medium    formatAttribute `json:"medium"`
	Xsmall    formatAttribute `json:"xsmall"`
	Thumbnail formatAttribute `json:"thumbnail"`
}

type formatAttribute struct {
	Ext    string  `json:"ext"`
	Url    string  `json:"url"`
	Hash   string  `json:"hash"`
	Mime   string  `json:"mime"`
	Name   string  `json:"name"`
	Path   string  `json:"path"`
	Size   float64 `json:"size"`
	Width  int64   `json:"width"`
	Height int64   `json:"height"`
}
