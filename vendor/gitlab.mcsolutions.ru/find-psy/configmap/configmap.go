package configmap

import (
	"strconv"

	"gitlab.mcsolutions.ru/lib/common/config"
	"gitlab.mcsolutions.ru/lib/common/consts"
)

const (
	APPLICATION  = "find-psy"
	APPLICATION_ = "find-psy-"

	USERSROUTER      = "usersrouter"
	ACCOUNTS         = "accounts"
	EMAIL            = "email"
	SMS              = "sms"
	USERS            = "users"
	PSYS             = "psys"
	LOCATIONS        = "locations"
	REGIONS          = "regions"
	SCHEDULE         = "schedule"
	SCHEDULESETTINGS = "schedulesettings"
	NOTWORKING       = "notworking"
	EDUCATIONS       = "educations"
	BROADCAST        = "broadcast"
	PAYMENTS         = "payments"
	FEEDBACK         = "feedback"
	FAQ              = "faq"
	ADMINSROUTER     = "adminsrouter"
	SUPERADMINROUTER = "superadminrouter"
	ADMINACCOUNTS    = "adminaccounts"
	LISTS            = "lists"
	USERSADMIN       = "usersadmin"
	PSYSADMIN        = "psysadmin"
	NEWS             = "news"
	EVENTS           = "events"
	ORGANIZATIONS    = "organizations"
	PUBLICATIONS     = "publications"
	TAGS             = "tags"
	PAYMENTSADMIN    = "paymentsadmin"
	IMAGES           = "images"
	PRICES           = "prices"

	API = "/api/"

	USERSROUTER_BASE_PATH      = API + USERSROUTER
	ACCOUNTS_BASE_PATH         = API + ACCOUNTS
	EMAIL_BASE_PATH            = API + EMAIL
	SMS_BASE_PATH              = API + SMS
	USERS_BASE_PATH            = API + USERS
	PSYS_BASE_PATH             = API + PSYS
	LOCATIONS_BASE_PATH        = API + LOCATIONS
	REGIONS_BASE_PATH          = API + REGIONS
	SCHEDULE_BASE_PATH         = API + SCHEDULE
	SCHEDULESETTINGS_BASE_PATH = API + SCHEDULESETTINGS
	NOTWORKING_BASE_PATH       = API + NOTWORKING
	EDUCATIONS_BASE_PATH       = API + EDUCATIONS
	BROADCAST_BASE_PATH        = API + BROADCAST
	PAYMENTS_BASE_PATH         = API + PAYMENTS
	FEEDBACK_BASE_PATH         = API + FEEDBACK
	FAQ_BASE_PATH              = API + FAQ
	ADMINSROUTER_BASE_PATH     = API + ADMINSROUTER
	SUPERADMINROUTER_BASE_PATH = API + SUPERADMINROUTER
	ADMINACCOUNTS_BASE_PATH    = API + ADMINACCOUNTS
	LISTS_BASE_PATH            = API + LISTS
	USERSADMIN_BASE_PATH       = API + USERSADMIN
	PSYSADMIN_BASE_PATH        = API + PSYSADMIN
	NEWS_BASE_PATH             = API + NEWS
	EVENTS_BASE_PATH           = API + EVENTS
	ORGANIZATIONS_BASE_PATH    = API + ORGANIZATIONS
	PUBLICATIONS_BASE_PATH     = API + PUBLICATIONS
	TAGS_BASE_PATH             = API + TAGS
	PAYMENTSADMIN_BASE_PATH    = API + PAYMENTSADMIN
	IMAGES_BASE_PATH           = API + IMAGES
	PRICES_BASE_PATH           = API + PRICES

	RESOURCE_GROUP           = "https://gitlab.mcsolutions.ru/find-psy/"
	BACK_USERS_GITLAB_GROUP  = RESOURCE_GROUP + "back/users/"
	BACK_ADMINS_GITLAB_GROUP = RESOURCE_GROUP + "back/admins/"

	USERSROUTER_GITLAB_SUBGROUP      = BACK_USERS_GITLAB_GROUP
	ACCOUNTS_GITLAB_SUBGROUP         = BACK_USERS_GITLAB_GROUP
	EMAIL_GITLAB_SUBGROUP            = BACK_USERS_GITLAB_GROUP
	SMS_GITLAB_SUBGROUP              = BACK_USERS_GITLAB_GROUP
	USERS_GITLAB_SUBGROUP            = BACK_USERS_GITLAB_GROUP
	PSYS_GITLAB_SUBGROUP             = BACK_USERS_GITLAB_GROUP
	LOCATIONS_GITLAB_SUBGROUP        = BACK_USERS_GITLAB_GROUP
	REGIONS_GITLAB_SUBGROUP          = BACK_USERS_GITLAB_GROUP
	SCHEDULE_GITLAB_SUBGROUP         = BACK_USERS_GITLAB_GROUP
	SCHEDULESETTINGS_GITLAB_SUBGROUP = BACK_USERS_GITLAB_GROUP
	NOTWORKING_GITLAB_SUBGROUP       = BACK_USERS_GITLAB_GROUP
	EDUCATIONS_GITLAB_SUBGROUP       = BACK_USERS_GITLAB_GROUP
	BROADCAST_GITLAB_SUBGROUP        = BACK_USERS_GITLAB_GROUP
	PAYMENTS_GITLAB_SUBGROUP         = BACK_USERS_GITLAB_GROUP
	FEEDBACK_GITLAB_SUBGROUP         = BACK_USERS_GITLAB_GROUP
	FAQ_GITLAB_SUBGROUP              = BACK_USERS_GITLAB_GROUP
	ADMINSROUTER_GITLAB_SUBGROUP     = BACK_ADMINS_GITLAB_GROUP
	SUPERADMINROUTER_GITLAB_SUBGROUP = BACK_ADMINS_GITLAB_GROUP
	ADMINACCOUNTS_GITLAB_SUBGROUP    = BACK_ADMINS_GITLAB_GROUP
	LISTS_GITLAB_SUBGROUP            = BACK_ADMINS_GITLAB_GROUP
	USERSADMIN_GITLAB_SUBGROUP       = BACK_ADMINS_GITLAB_GROUP
	PSYSADMIN_GITLAB_SUBGROUP        = BACK_ADMINS_GITLAB_GROUP
	NEWS_GITLAB_SUBGROUP             = BACK_ADMINS_GITLAB_GROUP
	EVENTS_GITLAB_SUBGROUP           = BACK_ADMINS_GITLAB_GROUP
	ORGANIZATIONS_GITLAB_SUBGROUP    = BACK_ADMINS_GITLAB_GROUP
	PUBLICATIONS_GITLAB_SUBGROUP     = BACK_ADMINS_GITLAB_GROUP
	TAGS_GITLAB_SUBGROUP             = BACK_ADMINS_GITLAB_GROUP
	PAYMENTSADMIN_GITLAB_SUBGROUP    = BACK_ADMINS_GITLAB_GROUP
	IMAGES_GITLAB_SUBGROUP           = BACK_ADMINS_GITLAB_GROUP
	PRICES_GITLAB_SUBGROUP           = BACK_ADMINS_GITLAB_GROUP
)

var (
	USERSROUTER_PORT      = config.GetEnvInt("_USERSROUTER_PORT", 3000)
	ACCOUNTS_PORT         = config.GetEnvInt("_ACCOUNTS_PORT", 3001)
	EMAIL_PORT            = config.GetEnvInt("_EMAIL_PORT", 3013)
	SMS_PORT              = config.GetEnvInt("_SMS_PORT", 3014)
	USERS_PORT            = config.GetEnvInt("_USERS_PORT", 3015)
	PSYS_PORT             = config.GetEnvInt("_PSYS_PORT", 3016)
	LOCATIONS_PORT        = config.GetEnvInt("_LOCATIONS_PORT", 3022)
	REGIONS_PORT          = config.GetEnvInt("_REGIONS_PORT", 3026)
	SCHEDULE_PORT         = config.GetEnvInt("_SCHEDULE_PORT", 3017)
	SCHEDULESETTINGS_PORT = config.GetEnvInt("_SCHEDULESETTINGS_PORT", 3023)
	NOTWORKING_PORT       = config.GetEnvInt("_NOTWORKING_PORT", 3024)
	EDUCATIONS_PORT       = config.GetEnvInt("_EDUCATIONS_PORT", 3025)
	BROADCAST_PORT        = config.GetEnvInt("_BROADCAST_PORT", 3018)
	PAYMENTS_PORT         = config.GetEnvInt("_PAYMENTS_PORT", 3019)
	FEEDBACK_PORT         = config.GetEnvInt("_FEEDBACK_PORT", 3020)
	FAQ_PORT              = config.GetEnvInt("_FAQ_PORT", 3021)
	ADMINSROUTER_PORT     = config.GetEnvInt("_ADMINSROUTER_PORT", 3004)
	SUPERADMINROUTER_PORT = config.GetEnvInt("_SUPERADMINROUTER_PORT", 3002)
	ADMINACCOUNTS_PORT    = config.GetEnvInt("_ADMINACCOUNTS_PORT", 3003)
	LISTS_PORT            = config.GetEnvInt("_LISTS_PORT", 3005)
	USERSADMIN_PORT       = config.GetEnvInt("_USERSADMIN_PORT", 3006)
	PSYSADMIN_PORT        = config.GetEnvInt("_PSYSADMIN_PORT", 3007)
	NEWS_PORT             = config.GetEnvInt("_NEWS_PORT", 3008)
	EVENTS_PORT           = config.GetEnvInt("_EVENTS_PORT", 3009)
	ORGANIZATIONS_PORT    = config.GetEnvInt("_ORGANIZATIONS_PORT", 3010)
	PUBLICATIONS_PORT     = config.GetEnvInt("_PUBLICATIONS_PORT", 3011)
	TAGS_PORT             = config.GetEnvInt("_TAGS_PORT", 3029)
	PAYMENTSADMIN_PORT    = config.GetEnvInt("_PAYMENTSADMIN_PORT", 3012)
	IMAGES_PORT           = config.GetEnvInt("_IMAGES_PORT", 3027)
	PRICES_PORT           = config.GetEnvInt("_PRICES_PORT", 3028)

	USERSROUTER_URL      = config.GetEnv("_USERSROUTER_URL", consts.LOCALHOST+strconv.Itoa(USERSROUTER_PORT))
	ACCOUNTS_URL         = config.GetEnv("_ACCOUNTS_URL", consts.LOCALHOST+strconv.Itoa(ACCOUNTS_PORT))
	EMAIL_URL            = config.GetEnv("_EMAIL_URL", consts.LOCALHOST+strconv.Itoa(EMAIL_PORT))
	SMS_URL              = config.GetEnv("_SMS_URL", consts.LOCALHOST+strconv.Itoa(SMS_PORT))
	USERS_URL            = config.GetEnv("_USERS_URL", consts.LOCALHOST+strconv.Itoa(USERS_PORT))
	PSYS_URL             = config.GetEnv("_PSYS_URL", consts.LOCALHOST+strconv.Itoa(PSYS_PORT))
	LOCATIONS_URL        = config.GetEnv("_LOCATIONS_URL", consts.LOCALHOST+strconv.Itoa(LOCATIONS_PORT))
	REGIONS_URL          = config.GetEnv("_REGIONS_URL", consts.LOCALHOST+strconv.Itoa(REGIONS_PORT))
	SCHEDULE_URL         = config.GetEnv("_SCHEDULE_URL", consts.LOCALHOST+strconv.Itoa(SCHEDULE_PORT))
	SCHEDULESETTINGS_URL = config.GetEnv("_SCHEDULESETTINGS_URL", consts.LOCALHOST+strconv.Itoa(SCHEDULESETTINGS_PORT))
	NOTWORKING_URL       = config.GetEnv("_NOTWORKING_URL", consts.LOCALHOST+strconv.Itoa(NOTWORKING_PORT))
	EDUCATIONS_URL       = config.GetEnv("_EDUCATIONS_URL", consts.LOCALHOST+strconv.Itoa(EDUCATIONS_PORT))
	BROADCAST_URL        = config.GetEnv("_BROADCAST_URL", consts.LOCALHOST+strconv.Itoa(BROADCAST_PORT))
	PAYMENTS_URL         = config.GetEnv("_PAYMENTS_URL", consts.LOCALHOST+strconv.Itoa(PAYMENTS_PORT))
	FEEDBACK_URL         = config.GetEnv("_FEEDBACK_URL", consts.LOCALHOST+strconv.Itoa(FEEDBACK_PORT))
	FAQ_URL              = config.GetEnv("_FAQ_URL", consts.LOCALHOST+strconv.Itoa(FAQ_PORT))
	ADMINSROUTER_URL     = config.GetEnv("_ADMINSROUTER_URL", consts.LOCALHOST+strconv.Itoa(ADMINSROUTER_PORT))
	SUPERADMINROUTER_URL = config.GetEnv("_SUPERADMINROUTER_URL", consts.LOCALHOST+strconv.Itoa(SUPERADMINROUTER_PORT))
	ADMINACCOUNTS_URL    = config.GetEnv("_ADMINACCOUNTS_URL", consts.LOCALHOST+strconv.Itoa(ADMINACCOUNTS_PORT))
	LISTS_URL            = config.GetEnv("_LISTS_URL", consts.LOCALHOST+strconv.Itoa(LISTS_PORT))
	USERSADMIN_URL       = config.GetEnv("_USERSADMIN_URL", consts.LOCALHOST+strconv.Itoa(USERSADMIN_PORT))
	PSYSADMIN_URL        = config.GetEnv("_PSYSADMIN_URL", consts.LOCALHOST+strconv.Itoa(PSYSADMIN_PORT))
	NEWS_URL             = config.GetEnv("_NEWS_URL", consts.LOCALHOST+strconv.Itoa(NEWS_PORT))
	EVENTS_URL           = config.GetEnv("_EVENTS_URL", consts.LOCALHOST+strconv.Itoa(EVENTS_PORT))
	ORGANIZATIONS_URL    = config.GetEnv("_ORGANIZATIONS_URL", consts.LOCALHOST+strconv.Itoa(ORGANIZATIONS_PORT))
	PUBLICATIONS_URL     = config.GetEnv("_PUBLICATIONS_URL", consts.LOCALHOST+strconv.Itoa(PUBLICATIONS_PORT))
	TAGS_URL             = config.GetEnv("_TAGS_URL", consts.LOCALHOST+strconv.Itoa(TAGS_PORT))
	PAYMENTSADMIN_URL    = config.GetEnv("_PAYMENTSADMIN_URL", consts.LOCALHOST+strconv.Itoa(PAYMENTSADMIN_PORT))
	IMAGES_URL           = config.GetEnv("_IMAGES_URL", consts.LOCALHOST+strconv.Itoa(IMAGES_PORT))
	PRICES_URL           = config.GetEnv("_PRICES_URL", consts.LOCALHOST+strconv.Itoa(PRICES_PORT))
)
