package common

const (
	PostCreatedEvent = "Event.Post.Created"
	PostUpdatedEvent = "Event.Post.Updated"
	PostDeletedEvent = "Event.Post.Deleted"
)

const (
	PostPublishedEvent = "Event.PostState.Published"
	PostTrashedEvent   = "Event.PostState.Trashed"
	PostRestoredEvent  = "Event.PostState.Restored"
)

const (
	PostCategoryUpdatedEvent = "Event.PostCategory.Updated"
	PostCategoryRemovedEvent = "Event.PostCategory.Removed"
)

const PostTagsUpdatedEvent = "Event.PostTags.Updated"

const (
	TagCreatedEvent        = "Event.Tag.Created"
	TagUpdatedEvent        = "Event.Tag.Updated"
	TagDeletedEvent        = "Event.Tag.Deleted"
	TagIncreasedUsageEvent = "Event.Tag.IncreasedUsage"
	TagReducedUsageEvent   = "Event.Tag.ReducedUsage"
)

const (
	CategoryCreatedEvent        = "Event.Category.Created"
	CategoryUpdatedEvent        = "Event.Cateogry.Updated"
	CategoryDeletedEvent        = "Event.Category.Deleted"
	CategoryIncreasedUsageEvent = "Event.Category.IncreasedUsage"
	CategoryReducedUsageEvent   = "Event.Category.ReducedUsage"
)

const (
	UserCreatedEvent       = "Event.User.Created"
	UserUpdatedEvent       = "Event.User.Updated"
	UserDeletedEvent       = "Event.User.Deleted"
	UserResetPasswordEvent = "Event.User.PassowrdReseted"
)
