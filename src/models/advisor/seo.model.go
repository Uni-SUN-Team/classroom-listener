package advisor

type seo struct {
	Id              int64         `json:"id"`
	MetaTitle       string        `json:"metaTitle"`
	MetaDescription string        `json:"metaDescription"`
	Keywords        string        `json:"keywords"`
	PreventIndexing bool          `json:"preventIndexing"`
	ShareImage      seoShareImage `json:"shareImage"`
}

type seoShareImage struct {
	Id    int64          `json:"id"`
	Alt   string         `json:"alt"`
	Media thumbnailLarge `json:"media"`
}
