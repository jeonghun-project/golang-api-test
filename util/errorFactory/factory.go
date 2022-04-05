package errorFactory

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	UnknownServerError                   = 5000
	Unauthorized                         = 5001
	EmailAlreadyRegistered               = 5002
	PasswordLessThanRequired             = 5003
	SomethingMissing                     = 5004
	WrongConfirmationCode                = 5005
	WrongAccessToken                     = 5006
	EmailNotRegistered                   = 5007
	EmailAlreadyConfirmed                = 5008
	WrongPassword                        = 5009
	DuplicateAPIKey                      = 5010
	WrongEmailOrPassword                 = 5011
	NameAlreadyRegistered                = 5012
	UserDeleted                          = 5013
	ContentMissing                       = 5014
	NoLeaguePresent                      = 5015
	CantUnfollowLOT                      = 5016
	ReferralCodeMissing                  = 5017
	EmailBelongsToAnotherAccount         = 5018
	PhoneNumberBelongsToAnotherAccount   = 5019
	ActionRequiresSubscription           = 5020
	FollowNumberFullUpgradeSubscription  = 5021
	UnfollowedManyTimesBlocked           = 5022
	NotFollowingTheUser                  = 5023
	PrivateDataNotAllowed                = 5024
	OverdrawnPoints                      = 5025
	CopyBotNotAllowed                    = 5026
	GuildDeleted                         = 5027
	LeagueConditionNotMet                = 5028
	CantUpdateAPIKeyDifferentAccount     = 5029
	CantUpdateAPIKeyExchangeNotSupported = 5030

	BotBalanceLessThanMinimum = 5031
	BotKeyCantDelete          = 5032
	BotCompanyError           = 5033

	FunctionNotSupportedYetForGivenExchange = 5035
	InvoicePaymentVerifyUnknownError        = 5036
	KeyHasNoPermissionForGivenOperation     = 5037

	LeagueRegisterByPassed      = 5040
	LeagueMinimumBalanceNotMet  = 5041
	LeagueRemovedDueToKeyDelete = 5042

	ErrorBotRatioDifferenceHigh      = 6001
	ErrorBotCopyingKeyInvalidated    = 6002
	ErrorBotOriginalKeyInvalidated   = 6003
	ErrorBotPositionMismatch         = 6004
	ErrorBotUnknown                  = 6005
	ErrorBotOpenOrderFound           = 6006
	ErrorKeyIsRegisteredInAnotherBot = 6007
	ErrorKeyIsNotCopyLeader          = 6008
	ErrorCopyLeaderConditionsNotMet  = 6009
	ErrorBotBalanceNotSufficient     = 6010
	ErrorBotDoesntExist              = 6011
	ErrorLeaderCopierCountReached    = 6012
	ErrorBannedFromBeingLeader       = 6013
	ErrorLeaderKeyCantBeABot         = 6014
)

var (
	OperationErrorBalanceNotFound  = errors.New("balance not found")
	OperationErrorPositionNotFound = errors.New("position not found")
	OperationErrorSideUknown       = errors.New("side unknown")
	OperationErrorUnknown          = errors.New("unknown")
	OperationNoPermission          = errors.New("no permission")
)

//Error returned by APIHandler
type Error struct {
	Error   error  `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func New(message string, code int) *Error {
	return &Error{nil, message, code}
}

func InternalServerError(err error) *Error {
	log.Println("Internal server error", err)
	return &Error{err, "Unknown error! Try again", UnknownServerError}
}

func NotAuthorizedError(message string) *Error {
	return &Error{nil, message, http.StatusUnauthorized}
}

func SomethingMissingError(message string) *Error {
	return &Error{nil, message, SomethingMissing}
}

func NotFollowingError() *Error {
	return &Error{nil, "not following", NotFollowingTheUser}
}
