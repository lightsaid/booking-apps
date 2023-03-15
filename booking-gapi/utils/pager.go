package utils

import "github.com/lightsaid/booking-gapi/pb"

func GetOffset(pager *pb.Pagation) (limit int32, offset int32) {
	if pager.Size <= 0 || pager.Size > 100 {
		limit = 10
	} else {
		limit = pager.Size
	}

	if pager.Page <= 0 {
		pager.Page = 1
	}

	offset = (pager.Page - 1) * limit

	return
}
