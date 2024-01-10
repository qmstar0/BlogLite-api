package dto

type CategoryCreateD struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	SeoDesc     string `json:"seoDesc"`
}
