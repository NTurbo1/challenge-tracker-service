package session

import (
	"time"
)

const sessionCSVHeader = "id,userId,createdAt,expiresAt"
const numSessionCSVCols = 5 // Depends on the sessionCSVHeader variable value. Keep it up to date with it!
const timeLayout = time.UnixDate
