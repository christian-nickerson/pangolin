package databases

import (
	"encoding/base64"
	"math"
	"strconv"

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
			Desc:   rq.OrderDesc,
		})

		if rq.ContinuationToken != "" {
			// base64 validation handled earlier by validator
			cursor, _ := rq.DecodeToken()
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
