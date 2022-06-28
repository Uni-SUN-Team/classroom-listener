package advisor

type categories struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	PublishedAt string `json:"publishedAt"`
	Locale      string `json:"locale"`
	Slug        string `json:"slug"`
}
