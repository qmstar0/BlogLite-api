package articles

import (
	"github.com/qmstar0/BlogLite-api/internal/common/e"
	"slices"
	"time"
)

type Version struct {
	Title,
	Description,
	Text,
	Hash,
	Source,
	Note string
}

func NewVersion(title, description, text, hash, source, note string) (v Version, err error) {
	if hash == "" {
		return Version{}, e.InvalidActionError("文章来源解析错误")
	}
	if source == "" {
		return Version{}, e.InvalidActionError("文章来源不可为空")
	}
	if text == "" {
		return Version{}, e.InvalidActionError("文章内容不可为空")
	}
	if title == "" {
		return Version{}, e.InvalidActionError("文章`title`不可为空")
	}
	if description == "" {
		return Version{}, e.InvalidActionError("文章`description`不可为空")
	}
	if note == "" {
		return Version{}, e.InvalidActionError("文章`note`不可为空")
	}

	return Version{
		Title:       title,
		Description: description,
		Text:        text,
		Hash:        hash,
		Source:      source,
		Note:        note,
	}, nil
}

func (a *Article) AddNewVresion(version Version) error {
	if slices.Contains(a.versionList, version.Hash) {
		return e.InvalidActionError("该版本已上传，不可重复上传")
	}
	a.versionList = append(a.versionList, version.Hash)
	now := time.Now()
	a.Emit(ArticleNewVersionCreatedEvent{
		URI:         a.uri.String(),
		Version:     version.Hash,
		Title:       version.Title,
		Content:     version.Text,
		Description: version.Description,
		Source:      version.Source,
		Note:        version.Note,
		CreatedAt:   now,
	})

	if !a.HasCurrentVersion() {
		a.currentVersion = version.Hash
		a.Emit(ArticleFirstVersionCreatedEvent{
			URI:       a.uri.String(),
			Version:   version.Hash,
			CreatedAt: now,
		})
	}

	return nil
}

func (a *Article) RemoveVersion(versionHash string) error {

	if a.CurrentVersion() == versionHash {
		return e.InvalidActionError("无法删除当前正在使用的版本")
	}

	index := slices.Index(a.versionList, versionHash)
	if index < 0 {
		return versionNotExist
	}

	a.versionList = slices.Delete(a.versionList, index, index+1)
	a.Emit(ArticleVersionContentDeletedEvent{
		URI:     a.uri.String(),
		Version: versionHash,
	})

	return nil
}

func (a *Article) CurrentVersion() string {
	return a.currentVersion
}

func (a *Article) HasCurrentVersion() bool {
	return a.currentVersion != ""
}

func (a *Article) SetCurrentVersion(versionHash string) error {
	if a.CurrentVersion() == versionHash {
		return nil
	}

	if !slices.Contains(a.versionList, versionHash) {
		return versionNotExist
	}

	a.setCurrentVersion(versionHash)

	a.Emit(ArticleContentSetSuccessfullyEvent{
		URI:     a.uri.String(),
		Version: versionHash,
	})

	return nil
}

func (a *Article) setCurrentVersion(versionHash string) {
	a.currentVersion = versionHash
}

var versionNotExist = e.InvalidActionError("该文章版本不存在或已删除")
