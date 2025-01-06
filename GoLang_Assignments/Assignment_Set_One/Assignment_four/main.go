package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "Assignment_four/quiz"
    "time"
)

func readString(reader *bufio.Reader, prompt string) string {
    fmt.Print(prompt)
    str, _ := reader.ReadString('\n')
    return strings.TrimSpace(str)
}

func readInt(reader *bufio.Reader, prompt string) (int, error) {
    str := readString(reader, prompt)
    return strconv.Atoi(str)
}

func showMenu() {
    fmt.Println("\n=== Online Examination System ===")
    fmt.Println("1. Start New Quiz")
    fmt.Println("2. Add New Question")
    fmt.Println("3. View Question Bank")
    fmt.Println("4. Exit")
    fmt.Print("Enter your choice: ")
}

func addQuestion(reader *bufio.Reader, qb *quiz.QuestionBank) {
    fmt.Println("\n--- Add New Question ---")
    
    question := readString(reader, "Enter Question: ")
    if question == "" {
        fmt.Println("Question cannot be empty")
        return
    }

    var options []string
    for i := 0; i < 4; i++ {
        option := readString(reader, fmt.Sprintf("Enter Option %d: ", i+1))
        if option == "" {
            fmt.Println("Option cannot be empty")
            return
        }
        options = append(options, option)
    }

    correctAnswer, err := readInt(reader, "Enter Correct Option Number (1-4): ")
    if err != nil || correctAnswer < 1 || correctAnswer > 4 {
        fmt.Println("Invalid correct answer. Must be between 1 and 4")
        return
    }

    q := quiz.NewQuestion(question, options, correctAnswer-1)
    qb.AddQuestion(q)
    fmt.Println("Question added successfully!")
}

func takeQuiz(reader *bufio.Reader, qb *quiz.QuestionBank) {
    if len(qb.GetQuestions()) == 0 {
        fmt.Println("No questions available in the question bank!")
        return
    }

    quizSession := quiz.NewQuizSession(qb.GetQuestions())
    fmt.Println("\n=== Quiz Started ===")
    fmt.Printf("You have %d seconds per question\n", quiz.QuestionTimeLimit/time.Second)
    fmt.Println("Type 'exit' to end the quiz early")
    fmt.Println("Press Enter to start...")
    reader.ReadString('\n')

    for i := 0; i < len(quizSession.Questions); i++ {
        q := quizSession.Questions[i]
        fmt.Printf("\nQuestion %d of %d:\n", i+1, len(quizSession.Questions))
        fmt.Println(q.GetQuestion())
        
        for j, option := range q.GetOptions() {
            fmt.Printf("%d. %s\n", j+1, option)
        }

        
        answerCh := make(chan int)
        timeoutCh := time.After(quiz.QuestionTimeLimit)

        
        go func() {
            for {
                answer := readString(reader, "Your answer (1-4): ")
                if strings.ToLower(answer) == "exit" {
                    answerCh <- -1
                    return
                }

                num, err := strconv.Atoi(answer)
                if err != nil || num < 1 || num > 4 {
                    fmt.Println("Invalid input. Please enter a number between 1-4")
                    continue
                }
                answerCh <- num - 1
                return
            }
        }()

        
        select {
        case answer := <-answerCh:
            if answer == -1 {
                fmt.Println("\nQuiz terminated early!")
                quizSession.CalculateScore()
                displayResults(quizSession)
                return
            }
            quizSession.RecordAnswer(i, answer)
        case <-timeoutCh:
            fmt.Println("\nTime's up for this question!")
            quizSession.RecordAnswer(i, -1)
        }
    }

    quizSession.CalculateScore()
    displayResults(quizSession)
}

func displayResults(session *quiz.QuizSession) {
    fmt.Println("\n=== Quiz Results ===")
    fmt.Printf("Score: %d/%d (%.1f%%)\n", 
        session.Score, len(session.Questions), session.Percentage)
    fmt.Printf("Performance: %s\n", session.GetPerformanceRating())
    
    fmt.Println("\nQuestion Review:")
    for i, q := range session.Questions {
        fmt.Printf("\nQ%d: %s\n", i+1, q.GetQuestion())
        fmt.Printf("Your Answer: ")
        if session.Answers[i] == -1 {
            fmt.Println("No answer (Time out)")
        } else {
            fmt.Printf("%s\n", q.GetOptions()[session.Answers[i]])
        }
        fmt.Printf("Correct Answer: %s\n", q.GetOptions()[q.GetCorrectAnswer()])
    }
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    qb := quiz.NewQuestionBank()

    
    qb.AddQuestion(quiz.NewQuestion(
        "What is the capital of France?",
        []string{"London", "Berlin", "Paris", "Madrid"},
        2,
    ))
    qb.AddQuestion(quiz.NewQuestion(
        "Which planet is known as the Red Planet?",
        []string{"Venus", "Mars", "Jupiter", "Saturn"},
        1,
    ))

    for {
        showMenu()
        choice, err := readInt(reader, "")
        if err != nil {
            fmt.Println("Invalid input. Please try again.")
            continue
        }

        switch choice {
        case 1:
            takeQuiz(reader, qb)

        case 2:
            addQuestion(reader, qb)

        case 3:
            questions := qb.GetQuestions()
            if len(questions) == 0 {
                fmt.Println("No questions in the question bank")
                continue
            }
            
            fmt.Println("\n=== Question Bank ===")
            for i, q := range questions {
                fmt.Printf("\nQuestion %d:\n%s\n", i+1, q.GetQuestion())
                for j, option := range q.GetOptions() {
                    fmt.Printf("%d. %s\n", j+1, option)
                }
                fmt.Printf("Correct Answer: %d\n", q.GetCorrectAnswer()+1)
            }

        case 4:
            fmt.Println("Thank you for using the Online Examination System. Goodbye!")
            return

        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}
