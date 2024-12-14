package controllers

import (
	"encoding/json"
	"strings"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/requests"
	"github.com/adon988/go_api_example/internal/responses"
	"github.com/adon988/go_api_example/internal/services"
	"github.com/adon988/go_api_example/internal/utils"
	"github.com/gin-gonic/gin"
)

type QuizController struct {
	InfoDb utils.InfoDb
}

// @Summary Create Quiz
// @Description Create a quiz
// @Tags Quiz
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param title body requests.QuizCreateRequest true "Quiz object that needs to be created"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string '{"code":-1,"data":{},"msg":""}'
// @Router /my/quiz [post]
func (c QuizController) CreateQuiz(ctx *gin.Context) {
	var req requests.QuizCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	err := utils.CheckQuestionTypes(req.QuestionTypes)
	if err != nil {
		responses.FailWithErrorCode(responses.INVALID_QUESTION_TYPE, ctx)
		return
	}

	Db, _ := c.InfoDb.InitDB()
	memberId, _ := ctx.Get("account")

	reqJson, _ := json.Marshal(req)
	hashInfos := []string{memberId.(string), string(reqJson)}
	quizId := utils.GenSha256IdempotentId(hashInfos)

	//check quiz exist
	quizService := services.NewQuizService(Db)
	quizExist := quizService.CheckQuizExist(quizId)
	if quizExist {
		responses.OkWithMessage("quiz exist", ctx)
		return
	}

	memberService := services.NewMemberService(Db)

	// add creater id to member ids
	memberIDs := append(req.MembersId, memberId.(string))
	// check members id is valid
	_, err = memberService.GetValidMembers(memberIDs)
	if err != nil {
		responses.FailWithErrorCode(responses.ACCOUNT_NOT_EXISTS, ctx)
		return
	}

	//get organization name
	organizationService := services.NewOrganizationService(Db)
	organization, err := organizationService.GetOrganizationByMemberIDAndOrganizationID(memberId.(string), req.OrganizationId)
	if err != nil {
		responses.FailWithErrorCode(responses.ORGANIZATION_NOT_FOUND, ctx)
		return
	}
	OrganizationInfo := models.ClassInfo{
		Id:    organization.Id,
		Title: organization.Title,
	}

	//get course name
	var CourseInfo models.ClassInfo
	if req.CourseId != "" {
		courseService := services.NewCourseSerive(Db)
		course, err := courseService.GetCourseByMemberIDAndCourseID(memberId.(string), req.CourseId)
		if err != nil {
			responses.FailWithErrorCode(responses.COURSE_NOT_FOUND, ctx)
			return
		}
		CourseInfo.Id = course.Id
		CourseInfo.Title = course.Title
	}

	//get unit name
	var UnitInfo models.ClassInfo
	if req.UnitId != "" {
		unitService := services.NewUnitService(Db)
		unit, err := unitService.GetUnitByMemberIDAndUnitID(memberId.(string), req.UnitId)
		if err != nil {
			responses.FailWithErrorCode(responses.UNIT_NOT_FOUND, ctx)
			return
		}
		UnitInfo.Id = unit.Id
		UnitInfo.Title = unit.Title
	}

	// get words
	var words []models.Word
	words = []models.Word{}

	// check UnitInfo is empty
	wordService := services.NewWordService(Db)
	if UnitInfo.Id != "" {
		words, _ = wordService.GetWordByMemberIDAndUnitID(memberId.(string), UnitInfo.Id)

	}

	if UnitInfo.Id == "" && CourseInfo.Id != "" {
		words, err := wordService.GetWordByMemberIDAndCourseID(memberId.(string), CourseInfo.Id)
		if err != nil {
			responses.FailWithErrorCode(responses.COURSE_NOT_FOUND, ctx)
			return
		}
		if len(words) == 0 {
			responses.FailWithErrorCode(responses.WORDS_NOT_FOUND_ON_COURSE, ctx)
			return
		}
	}

	// enstablish info
	quizInfo := models.QuizInfo{
		QuizCount:    req.QuizCount,
		QuizDays:     0, //default (just for challenge)
		RetryTime:    0, //default
		Organization: OrganizationInfo,
		Course:       CourseInfo,
		Unit:         UnitInfo,
	}

	// verify the words length

	if len(words) == 0 {
		responses.FailWithErrorCode(responses.WORDS_NOT_FOUND_ON_UNIT, ctx)
		return
	}

	if len(words) <= 2 {
		responses.FailWithErrorCode(responses.SHOULD_HAVE_AT_LEST_2_WORDS, ctx)
		return
	}

	//generate quiz content by factory
	contentItems, errCode := utils.GetQuizContent(req.QuestionTypes, words)
	if errCode != responses.SUCCESS {
		responses.FailWithErrorCode(errCode, ctx)
		return
	}

	questionTypes := strings.Join(req.QuestionTypes, ",")

	// create quiz
	quiz := models.Quiz{
		Id:             quizId,
		CreaterId:      memberId.(string),
		QuestionType:   questionTypes,
		OrganizationId: &OrganizationInfo.Id,
		CourseId:       &CourseInfo.Id,
		UnitId:         &UnitInfo.Id,
		Topic:          req.Topic,
		Type:           1,
		Info:           utils.MarshalJSONToRaw(quizInfo),
		Content:        utils.MarshalJSONToRaw(contentItems),
	}

	var QuizAnswerRecordUnstart int32 = 1
	quizAnswerRecordId, _ := utils.GenId()
	initQuizAnswerRecord := models.QuizAnswerRecord{
		Id:       quizAnswerRecordId,
		QuizId:   quizId,
		MemberId: memberId.(string),
		Status:   QuizAnswerRecordUnstart,
	}

	err = quizService.InitQuizWithTx(quiz, initQuizAnswerRecord)
	if err != nil {
		responses.FailWithErrorCode(responses.FAILED_TO_CREATE_QUIZ, ctx)
		return
	}
	// create quiz answer record
	responses.Ok(ctx)
}
