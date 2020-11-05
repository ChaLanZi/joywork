package auth

import "time"

const (
	cookieName = "joywork-faraday"
	cookie
	uuidKey               = "uuid"
	supportKey            = "support"
	expirationKey         = "exp"
	currentUserMetadata   = "faraday-current-user-uuid"
	currentUserHeader     = "Grpc-Metadata-Faraday-Current-User-Uuid"
	AuthorizationHeader   = "Authorization"
	AuthorizationMetadata = "authorization"
)

var (
	signingSecret string
	shortSession  = time.Duration(12 * time.Hour)
	longSession   = time.Duration(30 * 24 * time.Hour)
)
