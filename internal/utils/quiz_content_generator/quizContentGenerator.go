package quizcontentgenerator

import (
	"math"
	"math/rand"

	"github.com/adon988/go_api_example/internal/models"
)

func GenerateQuizContent(words []models.Word) (content string, err error) {

	return "quiz content", nil
}

func GenerateMutipleChoiceContent(words []models.Word) (contentItems models.ContentItems, err error) {
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
			Word:          word.Word,
			Definition:    word.Definition,
			Pronunciation: word.Pronunciation,
		}
		contentItems.ContentItems = append(contentItems.ContentItems, ContentItem)
	}
	return contentItems, nil
}

func GenerateTrueFalseContent(words []models.Word) (contentItems models.ContentItems, err error) {
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

		contentItem := models.ContentItem{
			QuestionType: "true_false",
			Question: []models.Title{
				{Title: randomWord.Word, Id: randomWord.Id},
			},
			Answer:        word.Id,
			WordId:        word.Id,
			Word:          word.Word,
			Definition:    word.Definition,
			Pronunciation: word.Pronunciation,
		}
		contentItems.ContentItems = append(contentItems.ContentItems, contentItem)

	}
	return contentItems, nil
}

func GenerateFullInBlankContent(words []models.Word) (contentItems models.ContentItems, err error) {
	for _, word := range words {
		var wordCount float64

		wordCount = math.Floor(float64(wordCount)/4) + 1
		if wordCount > 4 {
			wordCount = 4
		}
		
		random.Intn(len(word.Word))

	}
	return models.ContentItems{}, nil
}
