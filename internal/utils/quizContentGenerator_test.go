package utils

import (
	"fmt"
	"testing"

	"github.com/adon988/go_api_example/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetBlankCount(t *testing.T) {
	expsNum := 3
	wordCount := 10
	blankCount := GetBlankCount(wordCount)
	assert.Equal(t, expsNum, blankCount)

	expsNum = 0
	wordCount = 0
	blankCount = GetBlankCount(wordCount)
	assert.Equal(t, expsNum, blankCount)

	expsNum = 1
	wordCount = 1
	blankCount = GetBlankCount(wordCount)
	assert.Equal(t, expsNum, blankCount)
}

func TestGetRandomNumber(t *testing.T) {
	wordCount := 0
	blankCount := GetBlankCount(wordCount)
	var randNumber []int = GetRandomNumber(blankCount, wordCount)
	assert.Equal(t, blankCount, len(randNumber))

	wordCount = 1
	blankCount = GetBlankCount(wordCount)
	randNumber = GetRandomNumber(blankCount, wordCount)
	assert.Equal(t, blankCount, len(randNumber))

	wordCount = 5
	blankCount = GetBlankCount(wordCount)
	randNumber = GetRandomNumber(blankCount, wordCount)
	assert.Equal(t, blankCount, len(randNumber))

	wordCount = 16
	blankCount = GetBlankCount(wordCount)
	randNumber = GetRandomNumber(blankCount, wordCount)
	assert.Equal(t, blankCount, len(randNumber))
}

func TestGetWordSlice(t *testing.T) {
	word := models.Word{
		Id:            "1",
		Word:          "應用程式集合",
		Definition:    "capital",
		Pronunciation: "ˈkapɪtl",
	}
	wordCount := len([]rune(word.Word))
	blankCount := GetBlankCount(wordCount)
	randNumber := GetRandomNumber(blankCount, wordCount)
	blankWord, answerSlice := GetBlankWordAndAnswerSlice(word, randNumber)
	assert.Equal(t, len(randNumber), len(answerSlice))
	assert.Equal(t, wordCount, len([]rune(blankWord))-2*blankCount)
}

func TestGenerateFullInBlankContent(t *testing.T) {
	words := []models.Word{
		{
			Id:            "1",
			Word:          "應用程式集合",
			Definition:    "Application Suite",
			Pronunciation: "ㄧㄥˋ ㄩㄝˋ ㄔㄥˊ ㄕˋ ㄐㄧˊㄏㄜˊ",
		},
		{
			Id:            "2",
			Word:          "country",
			Definition:    "國家",
			Pronunciation: "ˈkʌntri",
		},
		{
			Id:            "3",
			Word:          "an",
			Definition:    "是",
			Pronunciation: "ˈæn",
		},
	}
	contentItems := GenerateFullInBlankContent(words)
	//less then 3 word would not generate any content, so here plus 1
	assert.Equal(t, len(words), len(contentItems.ContentItems)+1)
	fmt.Println(contentItems)
	for i, contentItem := range contentItems.ContentItems {
		assert.Equal(t, "full_in_blank", contentItem.QuestionType)
		assert.Equal(t, words[i].Id, contentItem.WordId)
		assert.Equal(t, words[i].Definition, contentItem.Definition)
		assert.Equal(t, words[i].Pronunciation, contentItem.Pronunciation)
	}
}

func TestGenerateTrueFalseConent(t *testing.T) {
	words := []models.Word{
		{
			Id:            "1",
			Word:          "應用程式集合",
			Definition:    "Application Suite",
			Pronunciation: "ㄧㄥˋ ㄩㄝˋ ㄔㄥˊ ㄕˋ ㄐㄧˊㄏㄜˊ",
		},
		{
			Id:            "2",
			Word:          "country",
			Definition:    "國家",
			Pronunciation: "ˈkʌntri",
		},
		{
			Id:            "3",
			Word:          "an",
			Definition:    "是",
			Pronunciation: "ˈæn",
		}, {
			Id:            "4",
			Word:          "capital",
			Definition:    "首都",
			Pronunciation: "ˈkæpɪtl",
		},
		{
			Id:            "5",
			Word:          "application",
			Definition:    "應用程式",
			Pronunciation: "ˌæplɪˈkeɪʃən",
		},
	}

	contentItems := GenerateTrueFalseContent(words)
	fmt.Println(contentItems)
	assert.Equal(t, len(words), len(contentItems.ContentItems))

	//if less then 2 word, would not generate any content
	words = []models.Word{
		{
			Id:            "1",
			Word:          "應用程式集合",
			Definition:    "Application Suite",
			Pronunciation: "ㄧㄥˋ ㄩㄝˋ ㄔㄥˊ ㄕˋ ㄐㄧˊㄏㄜˊ",
		},
		{
			Id:            "2",
			Word:          "country",
			Definition:    "國家",
			Pronunciation: "ˈkʌntri",
		},
	}
	contentItems = GenerateTrueFalseContent(words)
	assert.Equal(t, 0, len(contentItems.ContentItems))
}

func TestGenerateMutipleChoiceContent(t *testing.T) {
	words := []models.Word{
		{
			Id:            "1",
			Word:          "應用程式集合",
			Definition:    "Application Suite",
			Pronunciation: "ㄧㄥˋ ㄩㄝˋ ㄔㄥˊ ㄕˋ ㄐㄧˊㄏㄜˊ",
		},
		{
			Id:            "2",
			Word:          "country",
			Definition:    "國家",
			Pronunciation: "ˈkʌntri",
		},
		{
			Id:            "3",
			Word:          "an",
			Definition:    "是",
			Pronunciation: "ˈæn",
		}, {
			Id:            "4",
			Word:          "capital",
			Definition:    "首都",
			Pronunciation: "ˈkæpɪtl",
		},
		{
			Id:            "5",
			Word:          "application",
			Definition:    "應用程式",
			Pronunciation: "ˌæplɪˈkeɪʃən",
		},
	}
	contetnItems := GenerateMutipleChoiceContent(words)
	fmt.Println(contetnItems)
	assert.Equal(t, len(words), len(contetnItems.ContentItems))

	//if less then 3 word, would not generate any content
	words = []models.Word{
		{
			Id:            "1",
			Word:          "應用程式集合",
			Definition:    "Application Suite",
			Pronunciation: "ㄧㄥˋ ㄩㄝˋ ㄔㄥˊ ㄕˋ ㄐㄧˊㄏㄜˊ",
		},
		{
			Id:            "2",
			Word:          "country",
			Definition:    "國家",
			Pronunciation: "ˈkʌntri",
		},
		{
			Id:            "3",
			Word:          "an",
			Definition:    "是",
			Pronunciation: "ˈæn",
		},
	}
	contetnItems = GenerateMutipleChoiceContent(words)
	assert.Equal(t, 0, len(contetnItems.ContentItems))
}
