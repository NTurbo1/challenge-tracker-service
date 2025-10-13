package session

import (
	"time"
)

const (
	sessionCSVHeader = "id,userId,createdAt,expiresAt,valid,offset"
	numSessionCSVCols = 6 // Depends on the sessionCSVHeader variable value. Keep it up to date with it!
	timeLayout = time.DateTime

	valueSizeId = 48
	valueSizeUserId = 8
	valueSizeCreatedAt = 19
	valueSizeExpiresAt = 19
	valueSizeValid = 1
	valueSizeOffset = 8

	sessionCSVRowSize = valueSizeId + 1 + 
						valueSizeUserId + 1 + 
						valueSizeCreatedAt + 1 +
						valueSizeExpiresAt + 1 +
						valueSizeValid + 1 +
						valueSizeOffset + 1

	columnOffsetId = 0
	columnOffsetUserId = valueSizeId + 1
	columnOffsetCreatedAt = valueSizeId + 1 + valueSizeUserId + 1
	columnOffsetExpiresAt = valueSizeId + 1 + valueSizeUserId + 1 + valueSizeCreatedAt + 1
	columnOffsetValid = valueSizeId + 1 + valueSizeUserId + 1 + valueSizeCreatedAt + 1 + 
						valueSizeExpiresAt + 1
	columnOffsetOffset = valueSizeId + 1 + valueSizeUserId + 1 + valueSizeCreatedAt + 1 + 
					     valueSizeExpiresAt + 1 + valueSizeValid + 1
)
