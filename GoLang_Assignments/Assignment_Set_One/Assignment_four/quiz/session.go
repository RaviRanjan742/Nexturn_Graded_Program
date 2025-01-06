// quiz/session.go
package quiz

type QuizSession struct {
    Questions  []*Question
    Answers    []int
    Score      int
    Percentage float64
}

func NewQuizSession(questions []*Question) *QuizSession {
    return &QuizSession{
        Questions: questions,
        Answers:   make([]int, len(questions)),
        Score:     0,
    }
}

func (qs *QuizSession) RecordAnswer(questionIndex, answer int) {
    qs.Answers[questionIndex] = answer
}

func (qs *QuizSession) CalculateScore() {
    qs.Score = 0
    for i, question := range qs.Questions {
        if qs.Answers[i] == question.correctAnswer {
            qs.Score++
        }
    }
    qs.Percentage = float64(qs.Score) / float64(len(qs.Questions)) * 100
}

func (qs *QuizSession) GetPerformanceRating() string {
    switch {
    case qs.Percentage >= ExcellentThreshold:
        return "Excellent"
    case qs.Percentage >= GoodThreshold:
        return "Good"
    case qs.Percentage >= PassThreshold:
        return "Pass"
    default:
        return "Needs Improvement"
    }
}