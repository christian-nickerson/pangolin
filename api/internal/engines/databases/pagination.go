package databases

import (
	"encoding/base64"
	"encoding/binary"
	"math"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/christian-nickerson/pangolin/api/internal/models"
)

// Cursor pagination Gorm Scope. Will return pageSize + 1 records
func Paginate(
	model interface{},
	rq *models.PaginationRequest,
	rp *models.PaginationResponse,
	db *gorm.DB,
) func(db *gorm.DB) *gorm.DB {

	// get total rows & pages
	var totalRows int64
	db.Model(model).Count(&totalRows)
	rp.TotalRecords = totalRows
	rp.TotalPages = int64(math.Ceil(float64(totalRows) / float64(rq.PageSize)))

	return func(db *gorm.DB) *gorm.DB {

		// set up base query without cursor token
		query := db.Limit(rq.PageSize + 1).Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   true,
		})

		if rq.ContinuationToken != "" {
			// base64 validation handled earlier by validator
			cursor, _ := rq.DecodeToken()
			query = query.Where("id <= ?", cursor)
		}

		return query
	}
}

// Process results into paginated response, page + continuation token
func PaginatedResponse[T models.IDGetter](records []T, pageSize int) ([]T, string) {
	var continuationToken string
	var page []T

	// create continuation token from last record
	if len(records) == pageSize+1 {
		page = records[:len(records)-1]
		lastRecord := records[len(records)-1]
		idByteString := make([]byte, 8)
		binary.LittleEndian.PutUint64(idByteString, lastRecord.GetID())
		continuationToken = base64.StdEncoding.EncodeToString(idByteString)
	} else {
		page = records
		continuationToken = ""
	}

	return page, continuationToken
}
