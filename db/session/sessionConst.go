package session

import (
	"time"
)

const (
	timeLayout = time.DateTime

	valueSizeId = 48
	valueSizeUserId = 8
	valueSizeCreatedAt = 19
	valueSizeExpiresAt = 19
	valueSizeValid = 1

	sessionCSVRowSize = valueSizeId + 1 + 
						valueSizeUserId + 1 + 
						valueSizeCreatedAt + 1 +
						valueSizeExpiresAt + 1 +
						valueSizeValid + 1

	columnOffsetId = 0
	columnOffsetUserId = valueSizeId + 1
	columnOffsetCreatedAt = valueSizeId + 1 + valueSizeUserId + 1
	columnOffsetExpiresAt = valueSizeId + 1 + valueSizeUserId + 1 + valueSizeCreatedAt + 1
	columnOffsetValid = valueSizeId + 1 + valueSizeUserId + 1 + valueSizeCreatedAt + 1 + 
						valueSizeExpiresAt + 1
)
