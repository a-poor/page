package page

import "github.com/go-git/go-billy/v5"

const DefaultPageSize uint = 1 << 12

const (
	isActiveData uint = 1 << iota
)

type Pager struct {
	file     billy.File
	PageSize uint
}

func NewPager(file billy.File, size uint) *Pager {
	if size == 0 {
		size = DefaultPageSize
	}
	return &Pager{
		file,
		size,
	}
}

type Page struct {
	Size uint
	Data []byte
}
