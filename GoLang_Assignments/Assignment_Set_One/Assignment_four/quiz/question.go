
package quiz

type Question struct {
    question      string
    options       []string
    correctAnswer int
}

func NewQuestion(question string, options []string, correctAnswer int) *Question {
    return &Question{
        question:      question,
        options:       options,
        correctAnswer: correctAnswer,
    }
}

func (q *Question) GetQuestion() string {
    return q.question
}

func (q *Question) GetOptions() []string {
    return q.options
}

func (q *Question) GetCorrectAnswer() int {
    return q.correctAnswer
}
