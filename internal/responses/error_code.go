package responses

const (
	//[group][c 1 r 2 u 3 d 4]
	ERROR         = -1
	ERROR_MESSAGE = "fail"

	//member 1
	ACCOUNT_NOT_EXISTS         = 100000
	USERNAME_OR_PASSWORD_ERROR = 100001

	//organization 2
	ORGANIZATION_NOT_FOUND = 220001

	//course 3
	COURSE_NOT_FOUND          = 320000
	WORDS_NOT_FOUND_ON_COURSE = 320001
	//unit 4
	UNIT_NOT_FOUND          = 420000
	WORDS_NOT_FOUND_ON_UNIT = 420001

	//word 5
	SHOULD_HAVE_AT_LEST_2_WORDS = 500002

	//quiz 6
	FAILED_TO_CREATE_QUIZ   = 610000
	QUIZ_NOT_FOUND          = 620000
	FAILED_TO_DEL_QUIZ      = 630000
	FAILED_TO_UPDATE_QUIZ   = 640000
	INVALID_QUESTION_TYPE   = 610001
	NO_CONTENT_GENERATED    = 610002
	FAILED_TO_GET_QUIZ_LIST = 610003
)