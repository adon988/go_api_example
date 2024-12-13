package controllers

import (
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

	memberService := services.NewMemberService(Db)

	// add creater id to member ids
	memberIDs := append(req.MembersId, memberId.(string))
	// check members id is valid
	_, err = memberService.GetValidMembers(memberIDs)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
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
			responses.FailWithMessage(err.Error(), ctx)
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
			responses.FailWithMessage(err.Error(), ctx)
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
			responses.FailWithMessage(err.Error(), ctx)
			return
		}
		if len(words) == 0 {
			responses.FailWithErrorCode(responses.NO_WORDS_FOUND_FOR_THE_COURSE, ctx)
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
		responses.FailWithErrorCode(responses.NO_WORDS_FOUND_FOR_THE_UNIT, ctx)
		return
	}

	if len(words) <= 2 {
		responses.FailWithErrorCode(responses.SHOULD_HAVE_AT_LEST_2_WORDS, ctx)
		return
	}

	//generate quiz content
	contentItems := []models.ContentItem{}

	mutipleChoiceContent := utils.GenerateMutipleChoiceContent(words)
	contentItems = append(contentItems, mutipleChoiceContent.ContentItems...)

	trueFalseContent := utils.GenerateTrueFalseContent(words)
	contentItems = append(contentItems, trueFalseContent.ContentItems...)

	fullInBlankContent := utils.GenerateFullInBlankContent(words)
	contentItems = append(contentItems, fullInBlankContent.ContentItems...)

	questionTypes := strings.Join(req.QuestionTypes, ",")
	quizId, _ := utils.GenId()
	// create quiz
	quiz := models.Quiz{
		Id:           quizId,
		CreaterId:    memberId.(string),
		QuestionType: questionTypes,
		Topic:        req.Topic,
		Type:         1,
		Info:         utils.MarshalJSONToRaw(quizInfo),
		Content:      utils.MarshalJSONToRaw(contentItems),
	}

	quizService := services.NewQuizService(Db)
	err = quizService.CreateQuiz(quiz)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	// create quiz answer record
	responses.Ok(ctx)
}
