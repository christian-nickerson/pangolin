package models

import (
	"encoding/base64"
	"encoding/binary"
)

type PaginationRequest struct {
	ContinuationToken string `query:"continuationToken" validate:"omitempty,base64"`
	PageSize          int    `query:"pageSize" validate:"required,min=5,max=100"`
}

func (p *PaginationRequest) DecodeToken() (uint64, error) {
	decodedByte, err := base64.StdEncoding.DecodeString(p.ContinuationToken)
	if err != nil {
		return 0, err
	}
	decodedToken := binary.LittleEndian.Uint64(decodedByte)
	return decodedToken, nil
}

type PaginationResponse struct {
	ContinuationToken string `json:"continuationToken"`
	TotalRecords      int64  `json:"totalRecords"`
	TotalPages        int64  `json:"totalPages"`
}
