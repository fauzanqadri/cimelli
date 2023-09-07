package models

import (
	"math"
)

type Photo struct {
	Id           uint64 `gorm:"primaryKey"`
	Name         string `gorm:"index,unique"`
	FileLocation string
	UrlPath      string
	CreatedAt    string `gorm:"->"`
	UpdatedAt    string `gorm:"->"`
	FileSize     int64
}

type PagedPhoto struct {
	Count                  int64
	CurrentPage, TotalPage int
	Photos                 []Photo
	Pages                  []int
}

func NewPhoto(name, floc, upath string, fsize int64) (p *Photo, err error) {
	id, err := Sf.NextID()

	if err != nil {
		return nil, err
	}

	return &Photo{
		Id:           id,
		Name:         name,
		FileLocation: floc,
		UrlPath:      upath,
		FileSize:     fsize,
	}, nil
}

func GetPagedPhoto(currentPage int) (p *PagedPhoto, err error) {
	var photoCount int64
	var offset int
	var totalPage float64
	photos := []Photo{}

	countRes := Db.Table("photos").Count(&photoCount)

	if countRes.Error != nil {
		return nil, countRes.Error
	}

	// photoCount 8
	// perPage 4
	// totalPage => photoCount / perPage => 8 / 4 => 2
	// offset => (page - 1) * perPage

	// page 1 => (1 - 1) * 4 => offset 0
	// page 2 => (2 - 1) * 4 => offset 4

	ttlPage := float64(photoCount) / 4

	totalPage = math.Ceil(ttlPage)

	offset = (currentPage - 1) * 4

	qRes := Db.Table("photos").Limit(4).Offset(offset).Order("created_at desc").Find(&photos)

	if qRes.Error != nil {
		return nil, qRes.Error
	}

	var pages []int
	for i := 1; i <= int(totalPage); i++ {
		pages = append(pages, i)
	}

	return &PagedPhoto{
		Count:       photoCount,
		CurrentPage: currentPage,
		TotalPage:   int(totalPage),
		Photos:      photos,
		Pages:       pages,
	}, nil

}

func (p *Photo) Insert() error {

	id, err := Sf.NextID()
	if err != nil {
		return err
	}

	p.Id = id

	qRes := Db.Table("photos").Create(p)

	if qRes.Error != nil {
		p.Id = 0
		return qRes.Error
	}

	return nil
}
