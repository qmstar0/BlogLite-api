package post

import "common/domainevent"

type Post struct {
	ID     uint32
	UserID uint32

	Uri   Uri
	Title string

	Content string

	SeoDesc string

	State State

	TagGroup   TagGroup
	CategoryID uint32

	events []domainevent.DomainEvent
}

func NewPost(userID uint32, uri Uri) *Post {
	hashID := uri.ToID()
	return &Post{
		ID:      hashID,
		UserID:  userID,
		Uri:     uri,
		Title:   "",
		Content: "",
		SeoDesc: "",
		State:   DraftState,
		TagGroup: TagGroup{
			Tags:     make([]uint32, 0),
			CheckBit: 0,
		},
		CategoryID: 0,
		events: []domainevent.DomainEvent{domainevent.NewDomainEvent(hashID, Created, PostCreated{
			UserID: hashID,
			Uri:    uri.String(),
		})},
	}
}

func (p *Post) SetTags(group TagGroup) {
	if group.CheckBit == p.TagGroup.CheckBit {
		return
	}

	p.events = append(p.events, domainevent.NewDomainEvent(p.ID, TagsChanged, PostTagsChanged{
		ID:  p.ID,
		Old: p.TagGroup.Tags,
		New: group.Tags,
	}))
	p.TagGroup = group
}

func (p *Post) SetCategory(newCate uint32) {
	if newCate <= 0 || newCate == p.CategoryID {
		return
	}
	p.events = append(p.events, domainevent.NewDomainEvent(p.ID, CategoryChanged, PostCategoryChanged{
		ID:  p.ID,
		Old: p.CategoryID,
		New: newCate,
	}))
	p.CategoryID = newCate
}

func (p *Post) Publish() {
	if !p.State.IsPublished() {
		p.events = append(p.events, domainevent.NewDomainEvent(p.ID, StateChanged, PostStateChanged{
			ID:  p.ID,
			Old: uint16(p.State),
			New: uint16(PublishedState),
		}))
		p.State = PublishedState
	}
}

func (p *Post) Trash() {
	if !p.State.IsTrash() {
		p.events = append(p.events, domainevent.NewDomainEvent(p.ID, StateChanged, PostStateChanged{
			ID:  p.ID,
			Old: uint16(p.State),
			New: uint16(TrashState),
		}))
		p.State = TrashState
	}
}

func (p *Post) Delete() {
	p.events = append(p.events, domainevent.NewDomainEvent(p.ID, Deleted, PostDeleted{ID: p.ID}))
}

func (p *Post) Modify(title, content, SeoDesc string) {
	if title != "" {
		p.events = append(p.events, domainevent.NewDomainEvent(p.ID, TitleChanged, PostChanged{
			ID:  p.ID,
			Old: p.Title,
			New: title,
		}))
		p.Title = title
	}

	if content != "" {
		p.events = append(p.events, domainevent.NewDomainEvent(p.ID, ContentChanged, PostChanged{
			ID:  p.ID,
			Old: p.Content,
			New: content,
		}))
		p.Content = content
	}

	if SeoDesc != "" {
		p.events = append(p.events, domainevent.NewDomainEvent(p.ID, SeoDescChanged, PostChanged{
			ID:  p.ID,
			Old: p.SeoDesc,
			New: SeoDesc,
		}))
		p.SeoDesc = SeoDesc
	}
}

func EventsFromEntity(entity *Post) []domainevent.DomainEvent {
	return entity.events
}
