package utils

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	"github.com/adon988/go_api_example/internal/models"
)

func GenerateQuizContent(words []models.Word) (content string, err error) {

	return "quiz content", nil
}

var allowQuestionTypes = map[string]bool{
	"multiple_choice": true,
	"true_false":      true,
	"full_in_blank":   true,
}

func CheckQuestionTypes(questionType []string) error {
	for _, qt := range questionType {
		if _, ok := allowQuestionTypes[qt]; !ok {
			return fmt.Errorf("invalid question type")
		}
	}
	return nil
}

func GenerateMutipleChoiceContent(words []models.Word) (contentItems models.ContentItems) {
	//if less then 3 word, would not generate any content
	if len(words) < 4 {
		return contentItems
	}
	for i, word := range words {
		//correct option
		options := []models.Title{
			{Title: word.Word, Id: word.Id},
		}
		//used indexes
		usedIndexes := map[int]bool{i: true}
		//random 3 words
		for len(options) < 4 {
			randomIndex := rand.Intn(len(words))
			if !usedIndexes[randomIndex] && words[randomIndex].Id != word.Id {
				options = append(options, models.Title{Title: words[randomIndex].Word, Id: words[randomIndex].Id})
				usedIndexes[randomIndex] = true
			}
		}
		//shuffle options (洗牌)
		rand.Shuffle(len(options), func(i, j int) {
			options[i], options[j] = options[j], options[i]
		})
		ContentItem := models.ContentItem{
			QuestionType:  "multiple_choice",
			Question:      options,
			Answer:        word.Id,
			WordId:        word.Id,
			Word:          word.Word,
			Definition:    word.Definition,
			Pronunciation: word.Pronunciation,
		}
		contentItems.ContentItems = append(contentItems.ContentItems, ContentItem)
	}
	return contentItems
}

func GenerateTrueFalseContent(words []models.Word) (contentItems models.ContentItems) {
	//if less then 2 word, would not generate any content
	if len(words) <= 2 {
		return contentItems
	}
	for i, word := range words {

		//shuffle one word
		var randomWord = word
		// random true or false
		var answerTrueOrFalse = false
		if rand.Intn(2) == 1 {
			answerTrueOrFalse = true
		}

		//If the answer is false, then get a random wrong word
		for !answerTrueOrFalse && randomWord.Id == word.Id {
			randomIndex := rand.Intn(len(words))
			if randomIndex != i {
				randomWord = words[randomIndex]
			}
		}
		var Answer string = "0"
		if answerTrueOrFalse {
			Answer = "1"
		}
		contentItem := models.ContentItem{
			QuestionType: "true_false",
			Question: []models.Title{
				{Title: randomWord.Word, Id: randomWord.Id},
			},
			Answer:        Answer,
			WordId:        word.Id,
			Word:          word.Word,
			Definition:    word.Definition,
			Pronunciation: word.Pronunciation,
		}
		contentItems.ContentItems = append(contentItems.ContentItems, contentItem)

	}
	return contentItems
}

func GenerateFullInBlankContent(words []models.Word) (contentItems models.ContentItems) {
	contentItems = models.ContentItems{}
	for _, word := range words {

		//using []rune to get the lenght of the word slice, but not len(word.Word) to suppoort no english word
		var wordCount int = len([]rune(word.Word))

		//should more then 3 alphabets to generate content
		if wordCount <= 3 {
			continue
		}

		var blankCount int = GetBlankCount(wordCount)

		var randomNumber []int = GetRandomNumber(blankCount, wordCount)

		//將字元轉為切片
		blankWord, answerSlice := GetBlankWordAndAnswerSlice(word, randomNumber)

		//將切片轉為字串
		contentItem := models.ContentItem{
			QuestionType: "full_in_blank",
			Question: []models.Title{
				{Title: blankWord},
			},
			Answer:        strings.Join(answerSlice, ","),
			WordId:        word.Id,
			Word:          word.Word,
			Definition:    word.Definition,
			Pronunciation: word.Pronunciation,
		}

		contentItems.ContentItems = append(contentItems.ContentItems, contentItem)
	}
	return contentItems
}

func GetBlankWordAndAnswerSlice(word models.Word, randomNumber []int) (blankWord string, answerSlice []string) {
	var wordSlice = []rune(word.Word)
	used := make(map[int]bool)
	for _, index := range randomNumber {
		used[index] = true
	}
	for k, v := range wordSlice {
		if used[k] {
			blankWord = blankWord + "[*]"
			answerSlice = append(answerSlice, string(v))
			continue
		}
		blankWord = blankWord + string(v)
	}

	return blankWord, answerSlice
}

func GetBlankCount(wordCount int) (blankCount int) {
	if wordCount == 0 {
		return 0
	}
	var dividesNm int = 3
	blankCount = int(math.Floor(float64(wordCount)/float64(dividesNm)) + 1)
	if blankCount > dividesNm {
		blankCount = dividesNm
	}
	return int(blankCount)
}

func GetRandomNumber(blankCount, wordCount int) (number []int) {
	if blankCount >= wordCount {
		blankCount = wordCount
	}
	number = []int{}
	used := make(map[int]bool)
	for len(number) < blankCount {
		n := rand.Intn(wordCount)
		if !used[n] {
			number = append(number, n)
			used[n] = true
		}
	}
	return number
}
