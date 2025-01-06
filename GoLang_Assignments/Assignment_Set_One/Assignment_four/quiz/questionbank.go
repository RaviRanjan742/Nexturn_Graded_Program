
package quiz

type QuestionBank struct {
    questions []*Question
}

func NewQuestionBank() *QuestionBank {
    return &QuestionBank{
        questions: make([]*Question, 0),
    }
}

func (qb *QuestionBank) AddQuestion(question *Question) {
    qb.questions = append(qb.questions, question)
}

func (qb *QuestionBank) GetQuestions() []*Question {
    return qb.questions
}
