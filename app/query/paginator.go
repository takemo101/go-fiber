package query

import (
	"math"
	"net/url"
	"strconv"

	"gorm.io/gorm"
)

// PagingParameter
type PagingParameter struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
}

const PaginationLimit = 10
const PaginationKey = "page"

type PaginatorElement struct {
	Page int    `json:"page"`
	URL  string `json:"url"`
}

// Paginator
type Paginator struct {
	TotalCount  int                `json:"total_count"`
	TotalPage   int                `json:"total_page"`
	Offset      int                `json:"offset"`
	Limit       int                `json:"limit"`
	CurrentPage int                `json:"current_page"`
	PrevPage    int                `json:"prev_page"`
	NextPage    int                `json:"next_page"`
	LastPage    int                `json:"last_page"`
	FirstCount  int                `json:"first_count"`
	LastCount   int                `json:"last_count"`
	PrevURL     string             `json:"prev_url"`
	NextURL     string             `json:"next_url"`
	Elements    []PaginatorElement `json:"elements"`
}

func (p *Paginator) SetURL(original string) {
	u, _ := url.Parse(original)
	query := u.Query()

	// PrevURL
	if p.CurrentPage > p.PrevPage {
		p.PrevURL = p.createURL(p.PrevPage, u, query)
	} else {
		p.PrevURL = ""
	}

	if p.CurrentPage < p.NextPage {
		p.NextURL = p.createURL(p.NextPage, u, query)
	} else {
		p.NextURL = ""
	}

	firstPage := p.CurrentPage - PaginationLimit
	if 1 > firstPage {
		firstPage = 1
	}

	lastPage := p.CurrentPage + PaginationLimit
	if lastPage > p.TotalPage {
		lastPage = p.TotalPage
	}

	elementLength := lastPage - (firstPage - 1)
	elements := make([]PaginatorElement, elementLength)

	count := 0
	for i := firstPage; i <= lastPage; i++ {
		elements[count] = PaginatorElement{
			Page: i,
			URL:  p.createURL(i, u, query),
		}
		count++
	}
	p.Elements = elements
}

func (p *Paginator) createURL(
	page int,
	u *url.URL,
	q url.Values,
) string {
	strconv.Itoa(p.TotalPage)
	q.Set(PaginationKey, strconv.Itoa(page))
	u.RawQuery = q.Encode()
	return u.String()
}

// Paging
func Paging(p *PagingParameter, result interface{}, paginator *Paginator) error {
	db := p.DB

	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	var offset int

	// record count
	count, countErr := countPaginationRecord(*db, result)
	if countErr != nil {
		return countErr
	}

	if p.Page <= 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	// find records
	if err := db.Limit(p.Limit).Offset(offset).Find(result).Error; err != nil {
		return err
	}

	// paginator data process
	paginator.TotalCount = count
	paginator.CurrentPage = p.Page

	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}

	paginator.FirstCount = offset
	nextCount := p.Page * p.Limit
	if count < nextCount {
		paginator.LastCount = count
	} else {
		paginator.LastCount = nextCount
	}

	return nil
}

func countPaginationRecord(db gorm.DB, anyType interface{}) (int, error) {
	var count int64
	err := db.Model(anyType).Count(&count).Error
	return int(count), err
}
