package databases

import (
	"github.com/christian-nickerson/pangolin/api/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Paginate(p *models.PaginationRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// set up base query without cursor token
		query := db.Limit(p.PageSize + 1).Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   p.OrderDesc,
		})

		if p.ContinuationToken != "" {
			// base64 validation handled earlier by validator
			cursor, _ := p.DecodeToken()
			query = query.Where("id <= ?", cursor)
		}

		return query
	}
}
