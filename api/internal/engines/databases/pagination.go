package databases

import (
	"encoding/base64"
	"strconv"

	"github.com/christian-nickerson/pangolin/api/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Cursor pagination Gorm Scope. Will return pageSize + 1 records
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

// Get next continuation token from pagination results
func GetContinuationToken[T models.IDGetter](records []T) string {
	// fetch last record from paginated limit + 1 records
	lastRecord := records[len(records)-1]
	idByteString := []byte(strconv.FormatUint(uint64(lastRecord.GetID()), 10))
	return base64.StdEncoding.EncodeToString(idByteString)
}
