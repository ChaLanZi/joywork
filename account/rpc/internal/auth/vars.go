package auth

import "time"

const (
	cookieName = "joywork-faraday"

	currentUserMetadata = "faraday-current-user-uuid"

	currentUserHeader = "Grpc-Metadata-Faraday-Current-User-Uuid"

	AuthorizationMetadata = "authorization"

	AuthorizationHeader = "Authorization"

	AuthorizationAnonymousWeb = "faraday-anonymous"

	AuthorizationSupportUser = "faraday-support"

	AuthorizationSuperpowersService = "super-powers-support"

	AuthorizationWWWService = "www-service"

	AuthorizationCompanyService = "company-service"

	AuthorizationAccountService = "account-service"

	AuthorizationWhoamiService = "whoami-service"

	AuthorizationBotService = "bot-service"

	AuthorizationAuthenticatedUser = "faraday-authenticated"
)

var (
	signingSecret string
	longSession   = 30 * 24 * time.Hour
	shortSession  = time.Hour * 12
)
