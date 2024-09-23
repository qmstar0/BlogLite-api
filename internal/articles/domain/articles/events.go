package articles

import "time"

type ArticleInitializedSuccessfullyEvent struct {
	URI        string
	CategoryID string
}

type ArticleVisibilityChangedEvent struct {
	URI        string
	Visibility bool
}

type ArticleFirstVersionCreatedEvent struct {
	URI       string
	Version   string
	CreatedAt time.Time
}

type ArticleCategoryChangedEvent struct {
	URI           string
	OldCategoryID string
	NewCategoryID string
}

type ArticleTagsModifiedEvent struct {
	URI     string
	OldTags []string
	NewTags []string
}

type ArticleDeletedEvent struct {
	URI string
}

type ArticleContentSetSuccessfullyEvent struct {
	URI     string
	Version string
}

type ArticleNewVersionCreatedEvent struct {
	URI         string
	Version     string
	Title       string
	Content     string
	Description string
	Source      string
	Note        string
	CreatedAt   time.Time
}

type ArticleVersionContentDeletedEvent struct {
	URI     string
	Version string
}

type CategoryCreatedEvent struct {
	Slug        string
	Name        string
	Description string
}

type CategoryDescriptionModifiedEvent struct {
	Slug        string
	Description string
}

type CategoryDeletedEvent struct {
	Slug string
}
