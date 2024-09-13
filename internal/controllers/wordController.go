package controllers

import (
	models "github.com/adon988/go_api_example/internal/models"
	"github.com/adon988/go_api_example/internal/services"
	"github.com/adon988/go_api_example/internal/utils"
	"github.com/adon988/go_api_example/internal/utils/requests"
	"github.com/adon988/go_api_example/internal/utils/responses"
	"github.com/gin-gonic/gin"
)

type WordController struct {
	InfoDb utils.InfoDb
}

// @Summary Create Word
// @Description Create a word
// @Tags Word
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param title body requests.WordCreateRequest true "Word object that needs to be created"
// @Success 200 {string} string "ok"
// @Router /my/word [post]
func (c WordController) CreateWord(ctx *gin.Context) {
	Db, _ := c.InfoDb.InitDB()
	var req requests.WordCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	wordService := services.NewWordService(Db)
	var word models.Word
	_ = ctx.ShouldBindJSON(&word)
	wordId, _ := utils.GenId()
	wordData := models.Word{
		Id:            wordId,
		UnitId:        req.UnitId,
		Word:          req.Word,
		Definition:    req.Definition,
		Image:         req.Image,
		Pronunciation: req.Pronunciation,
		Description:   req.Description,
		Comment:       req.Comment,
		Order:         req.Order,
	}

	err := wordService.CreateWord(wordData)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}

// @Summary Update Word
// @Description Update a word
// @Tags Word
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param title body requests.WordUpdateRequest true "Word object that needs to be updated"
// @Success 200 {string} string "ok"
// @Router /my/word [put]
func (c WordController) UpdateWord(ctx *gin.Context) {
	Db, _ := c.InfoDb.InitDB()
	var req requests.WordUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}

	wordService := services.NewWordService(Db)
	word := models.Word{
		Id:            req.Id,
		UnitId:        req.UnitId,
		Word:          req.Word,
		Definition:    req.Definition,
		Image:         req.Image,
		Pronunciation: req.Pronunciation,
		Description:   req.Description,
		Comment:       req.Comment,
		Order:         req.Order,
	}
	_ = ctx.ShouldBindJSON(&word)
	err := wordService.UpdateWord(word)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}

// @Summary Delete Word
// @Description Delete a word
// @Tags Word
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param title body requests.WordDeleteRequest true "Word object that needs to be deleted"
// @Success 200 {string} string "ok"
// @Router /my/word [delete]
func (c WordController) DeleteWord(ctx *gin.Context) {
	Db, _ := c.InfoDb.InitDB()
	var req requests.WordDeleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	wordService := services.NewWordService(Db)
	id := req.Id
	err := wordService.DeleteWord(id)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	responses.Ok(ctx)
}

// @Summary Get Words
// @Description Get words by unit id
// @Tags Word
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param unit_id query string true "unit id"
// @Success 200 {object} responses.WordResponse
// @Router /my/words [get]
func (c WordController) GetWords(ctx *gin.Context) {
	Db, _ := c.InfoDb.InitDB()
	unitId := ctx.Query("unit_id")
	wordService := services.NewWordService(Db)
	var wordsRes []responses.WordResponse
	var err error
	words, err := wordService.GetWordByUnitID(unitId)
	if err != nil {
		responses.FailWithMessage(err.Error(), ctx)
		return
	}
	for _, word := range words {
		wordsRes = append(wordsRes, responses.WordResponse{
			Id:            word.Id,
			UnitId:        word.UnitId,
			Word:          word.Word,
			Definition:    word.Definition,
			Image:         word.Image,
			Pronunciation: word.Pronunciation,
			Description:   word.Description,
			Comment:       word.Comment,
			Order:         word.Order,
		})
	}

	responses.OkWithData(wordsRes, ctx)
}
