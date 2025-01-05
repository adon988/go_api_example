package controllers

import (
	"encoding/json"
	"strconv"
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
	memberId := ctx.GetString("account")

	reqJson, _ := json.Marshal(req)
	hashInfos := []string{memberId, string(reqJson)}
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
	memberIDs := append(req.MembersId, memberId)
	// check members id is valid
	_, err = memberService.GetValidMembers(memberIDs)
	if err != nil {
		responses.FailWithErrorCode(responses.ACCOUNT_NOT_EXISTS, ctx)
		return
	}

	//get organization name
	organizationService := services.NewOrganizationService(Db)
	organization, err := organizationService.GetOrganizationByMemberIDAndOrganizationID(memberId, req.OrganizationId)
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
		course, err := courseService.GetCourseByMemberIDAndCourseID(memberId, req.CourseId)
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
		unit, err := unitService.GetUnitByMemberIDAndUnitID(memberId, req.UnitId)
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
		words, _ = wordService.GetWordByMemberIDAndUnitID(memberId, UnitInfo.Id)

	}

	if UnitInfo.Id == "" && CourseInfo.Id != "" {
		words, err := wordService.GetWordByMemberIDAndCourseID(memberId, CourseInfo.Id)
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
		CreaterId:      memberId,
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
		MemberId: memberId,
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

// @Summary Get Quiz List With Answers By Member
// @Description Get Quiz List With Answers By Member
// @Tags Quiz
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param req body requests.QuizListRequest true "Quiz object page list"
// @Success 200 {object} responses.QuizListResponse
// @Router /my/quiz_list [get]
func (c QuizController) GetQuizsListWithAnswersByMember(ctx *gin.Context) {
	var req requests.QuizListRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	memberId := ctx.GetString("account")
	var reqPage string = ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(reqPage)
	page := int32(pageInt)

	Db, _ := c.InfoDb.InitDB()

	quizServices := services.NewQuizService(Db)
	quizList, total, err := quizServices.GetQuizsListWithAnswersByMember(memberId, page)
	if err != nil {
		responses.FailWithErrorCode(responses.FAILED_TO_GET_QUIZ_LIST, ctx)
		return
	}

	data := responses.QuizListResponse{
		QuizList: quizList,
		Total:    total,
	}

	responses.OkWithData(data, ctx)
}

// @Summary Get Quiz By Member
// @Description Get Quiz By Member
// @Tags Quiz
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param quiz_id path string true "Quiz ID"
// @Success 200 {object} responses.QuizWithAnswer
// @Failure 400 {string} string '{"code":620000,"data":{},"msg":"fail"}'
// @Router /my/quiz_answer_record [patch]
func (c QuizController) GetQuizByMember(ctx *gin.Context) {
	quizId := ctx.Param("quiz_id")
	memberId := ctx.GetString("account")
	Db, _ := c.InfoDb.InitDB()
	quizService := services.NewQuizService(Db)
	quiz, err := quizService.GetQuizByMember(quizId, memberId)
	if err != nil {
		responses.FailWithErrorCode(responses.QUIZ_NOT_FOUND, ctx)
	}

	responses.OkWithData(quiz, ctx)
}

// @Summary Update Quiz Answer Record
// @Description Update Quiz Answer Record
// @Tags Quiz
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param quiz_id path string true "Quiz ID"
// @Param req body requests.QuizUpdateQuizAnswerRecordRequest true "Quiz object that needs to be updated"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string '{"code":620003,"data":{},"msg":"fail"}'
// @Router /my/quiz/{quiz_id} [patch]
func (c QuizController) UpdateQuizAnswerRecord(ctx *gin.Context) {
	var req requests.QuizUpdateQuizAnswerRecordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	memberId := ctx.GetString("account")
	quizId := req.QuizId

	var answerQuestions models.AnswerQuestions
	if err := json.Unmarshal([]byte(req.AnswerQuestion), &answerQuestions); err != nil {
		responses.FailWithErrorCode(responses.INVALID_ANSWER_QUESTION_FORMAT, ctx)
		return
	}

	Db, _ := c.InfoDb.InitDB()
	quizService := services.NewQuizService(Db)

	quiz, err := quizService.GetQuizByMember(quizId, memberId)
	if err != nil {
		responses.FailWithErrorCode(responses.QUIZ_NOT_FOUND, ctx)
		return
	}
	ContentItems := &models.ContentItems{}
	json.Unmarshal([]byte(quiz.Content), &ContentItems)

	//簡單匹配 quiz content 和 answer question
	if len(ContentItems.ContentItems) != len(answerQuestions.AnswerQuestion) {
		responses.FailWithErrorCode(responses.INVALID_ANSWER_QUESTION, ctx)
		return
	}

	// 紀錄錯題本
	var FailLogs []models.FailedLog
	for i, v := range answerQuestions.AnswerQuestion {
		if v.Answer != ContentItems.ContentItems[i].Answer {
			//FailLog include the anser and original quiz content
			FailLog := models.FailedLog{
				UserAnswer:  v.Answer,
				ContentItem: ContentItems.ContentItems[i],
			}
			FailLogs = append(FailLogs, FailLog)
		}
	}

	// update quiz answer record
	FailedAnswerCount := int32(len(FailLogs))
	TotalQuestionCount := int32(len(ContentItems.ContentItems))
	Scope := int32((1 - (float32(FailedAnswerCount) / float32(TotalQuestionCount))) * 100)
	AnswerQuestions := json.RawMessage(req.AnswerQuestion)

	quizAnswerRecord := models.QuizAnswerRecord{
		Id:                 quiz.QuizAnswerRecordId,
		AnswerQuestion:     &AnswerQuestions,
		Status:             int32(3), //finished
		FailedAnswerCount:  &FailedAnswerCount,
		TotalQuestionCount: &TotalQuestionCount,
		FailedLogs:         utils.MarshalJSONToRaw(FailLogs),
		Scope:              &Scope,
	}
	if err := quizService.UpdateQuizAnswerRecord(quizAnswerRecord); err != nil {
		responses.FailWithErrorCode(responses.FAILED_TO_UPDATE_QUIZ_ANSWER, ctx)
		return
	}

	responses.Ok(ctx)
}
