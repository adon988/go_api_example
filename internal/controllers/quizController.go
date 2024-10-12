package controllers

import (
	"strings"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/services"
	"github.com/adon988/go_api_example/internal/utils"
	"github.com/adon988/go_api_example/internal/utils/requests"
	"github.com/adon988/go_api_example/internal/utils/responses"
	"github.com/gin-gonic/gin"
)

type QuizController struct {
	InfoDb utils.InfoDb
}

func (c QuizController) CreateQuiz(ctx *gin.Context) {
	var req requests.QuizCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	err := utils.CheckQuestionTypes(req.QuestionTypes)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
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
		responses.FailWithMessage("organiztion not found", ctx)
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
	// check UnitInfo is empty
	wordService := services.NewWordService(Db)
	if UnitInfo.Id != "" {
		words, err := wordService.GetWordByMemberIDAndUnitID(memberId.(string), UnitInfo.Id)
		if err != nil {
			responses.FailWithMessage(err.Error(), ctx)
			return
		}
		if len(words) == 0 {
			responses.FailWithMessage("no words found for the unit", ctx)
			return
		}
	}

	if UnitInfo.Id == "" && CourseInfo.Id != "" {
		words, err := wordService.GetWordByMemberIDAndCourseID(memberId.(string), CourseInfo.Id)
		if err != nil {
			responses.FailWithMessage(err.Error(), ctx)
			return
		}
		if len(words) == 0 {
			responses.FailWithMessage("no words found for the course", ctx)
			return
		}
	}

	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
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

	// enstablish quiz content

	contentItems := models.ContentItems{}
	// an utils function to generate quiz content

	ContentItem := models.ContentItem{
		QuestionType: "multiple_choice",
		Question: []models.Title{
			{Title: "What is the capital of France?"},
			{Title: "What is the capital of France2?"},
			{Title: "What is the capital of France3?"},
		},
		Answer:        "0",
		WordId:        "1",
		Word:          "capital",
		Definition:    "The city or town that functions as the seat of government and administrative centre of a country or region.",
		Pronunciation: "ˈkapɪt(ə)l",
	}
	contentItems.ContentItems = append(contentItems.ContentItems, ContentItem)

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
