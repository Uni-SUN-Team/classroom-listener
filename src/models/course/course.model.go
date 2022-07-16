package course

import "unisun/api/classroom-listener/src/models/advisor"

type courses struct {
	Data []CourseData `json:"data"`
	Meta pagination   `json:"meta"`
}

type course struct {
	Data CourseData `json:"data"`
	Meta pagination `json:"meta"`
}

type CourseData struct {
	Id              int64                 `json:"id"`
	Title           string                `json:"title"`
	Description     string                `json:"description"`
	Content         string                `json:"content"`
	Short_content   string                `json:"short_content"`
	CreatedAt       string                `json:"createdAt"`
	UpdatedAt       string                `json:"updatedAt"`
	PublishedAt     string                `json:"publishedAt"`
	Locale          string                `json:"locale"`
	Slug            string                `json:"slug"`
	Thumbnail       thumbnailLarge        `json:"thumbnail"`
	CourseHighLight video                 `json:"course_high_light"`
	Advisors        []advisor.AdvisorData `json:"advisors"`
	Categories      []categories          `json:"categories"`
	SEO             seo                   `json:"SEO"`
	VideoContent    video                 `json:"video_content"`
	VideoPreview    video                 `json:"video_preview"`
	Attachment      []thumbnailLarge      `json:"file"`
}

type pagination struct {
	Pagination paginationContent `json:"pagination"`
}

type paginationContent struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"pageSize"`
	PageCount int64 `json:"pageCount"`
	Total     int64 `json:"total"`
}
