package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"math/rand"
	"strconv"
	"time"
)

type Proofs struct {
	app.Compo
	rand          *rand.Rand
	inputValue    string
	Q1input       string
	Q2input       string
	Q3input       string
	Q4input       string
	Q5input       string
	Q6input       string
	Q7option      string
	question1     Question
	question2     Question
	question3     Question
	question4     Question
	question5     Question
	question6     Question
	theLevel      int64
	seedClicked   bool
	answerClicked bool
	currentValue  string
}

func (pr *Proofs) Render() app.UI {
	return app.Div().Class("container-fluid").Style("background-color", "#F0F8FF").Style("min-height", "100vh").Body(
		app.Div().Class("row bg-primary").Style("min-height", "50px").Body(),
		app.Div().Class("row").Style("background-color", "#add8e6").Body(
			app.H1().Style("width", "90%").Style("margin", "0 auto").Body(
				app.Text("HelloMath"),
			),
		),
		app.Div().Class("row").Body(
			app.Div().Class("col-2").Body(&Sidebar{}),
			app.Div().Class("col-10").Body(
				app.Div().Style("height", "8vh").Body(
					app.H6().
						Style("margin-top", "8vh").
						Style("text-align", "center").
						Text("Instructions HelloMath!"),
				),

				app.Div().Body(
					app.Input().
						ID("seedInput").
						Type("text").
						Value(pr.inputValue).
						OnChange(pr.OnInputChange),
					app.Button().
						Text("Generate Seed").
						OnClick(pr.GenerateSeed),
				),
				app.H5().Text("Seed Value: "+pr.currentValue),
				app.Div().Style("height", "3vh").Body(),
				app.If(pr.seedClicked,
					//question 1
					app.Div().Class("row").Body(
						app.Div().Class("col-auto").Body(
							app.H6().Text("Q1 :"),
						),
						app.Div().Class("col").Body(
							app.P().Text(pr.question1.statement1),
						),
					),
					app.Div().Class("row").Body(
						app.Div().Class("col-auto").Body(
							app.Input().
								ID("Q1input").
								Type("text").
								Value(pr.Q1input).
								Style("width", "200px").
								Style("height", "30px").
								Style("margin-left", "50px").
								Class("form-control mb-3").
								Placeholder("").
								AutoFocus(true),
						),
						app.If(pr.answerClicked,
							app.If(pr.Q1input == pr.question1.answer,
								//app.If(se.question1.answer == se.Q1input,
								app.Div().Class("col").Body(
									app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
								),
							),
						),

						app.If(pr.answerClicked,
							app.P().Text("Answer: "+pr.question1.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
						//question 2

						//
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(pr.question2.statement1),
							),
						),
						app.Input().
							ID("seedInput").
							Type("text").
							Value(pr.Q2input).
							Style("width", "200px").
							Style("height", "30px").
							Style("margin-left", "50px").
							Class("form-control mb-3").
							Placeholder("").
							AutoFocus(true),

						app.If(pr.answerClicked,
							app.P().Text("Answer: "+pr.question2.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
						//question 3
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(pr.question3.statement1),
							),
						),
						app.Input().
							ID("seedInput").
							Type("text").
							Value(pr.Q3input).
							Style("width", "200px").
							Style("height", "30px").
							Style("margin-left", "50px").
							Class("form-control mb-3").
							Placeholder("").
							AutoFocus(true),

						app.If(pr.answerClicked,
							app.If(pr.question1.answer == pr.Q1input,
								app.Div().Class("col").Body(
									app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
								),
							),
						),
						//question 4
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q4 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(pr.question4.statement1),
							),
						),

						app.Input().
							ID("seedInput").
							Type("text").
							Value(pr.Q4input).
							Style("width", "200px").
							Style("height", "30px").
							Style("margin-left", "50px").
							Class("form-control mb-3").
							Placeholder("").
							AutoFocus(true),

						app.If(pr.answerClicked,
							app.If(pr.question1.answer == pr.Q1input,
								app.Div().Class("col").Body(
									app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
								),
							),
						),
						//question 5
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q5 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(pr.question5.statement1),
							),
						),
						app.Input().
							ID("seedInput").
							Type("text").
							Value(pr.Q5input).
							Style("width", "200px").
							Style("height", "30px").
							Style("margin-left", "50px").
							Class("form-control mb-3").
							Placeholder("").
							AutoFocus(true),

						app.If(pr.answerClicked,
							app.If(pr.question1.answer == pr.Q1input,
								app.Div().Class("col").Body(
									app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
								),
							),
						),
						//question 6
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q6 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(pr.question6.statement1),
							),
						),

						app.Input().
							ID("seedInput").
							Type("text").
							Value(pr.Q6input).
							Style("width", "200px").
							Style("height", "30px").
							Style("margin-left", "50px").
							Class("form-control mb-3").
							Placeholder("").
							AutoFocus(true),
					),

					//Name("options")
					app.Div().Class("form-check").Body(
						app.Input().
							Type("radio").
							Name("options").
							Class("form-check-input").
							Value("Option A").
							ID("option-a"),
						app.Label().Class("form-check-label").Text("True"),
					),
					app.Div().Class("form-check").Body(
						app.Input().
							Type("radio").
							Name("options").
							Class("form-check-input").
							Value("Option B").
							ID("option-b"),
						app.Label().Class("form-check-label").For("option-b").Text("False"),
					),
					app.If(pr.answerClicked,
						app.If(pr.Q7option == "Option B",
							app.Div().Class("col").Body(
								app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
							),
						),
					),
					app.Button().
						Text("Finish").
						Style("margin-left", "50px").
						OnClick(pr.finishClicked),
				),
			),
		),
	)

}

func (pr *Proofs) OnInputChange(ctx app.Context, e app.Event) {
	pr.inputValue = ctx.JSSrc().Get("value").String()
	pr.Update()
}
func (pr *Proofs) GenerateSeed(ctx app.Context, e app.Event) {
	inputValue := app.Window().GetElementByID("seedInput").Get("value").String()
	fmt.Println("input value:", inputValue)
	seed, err := strconv.ParseInt(inputValue, 10, 64)
	if err != nil {
		fmt.Println("No input or error converting string to integer:", err)
		seconds := time.Now().UnixNano() / 1000000
		seed = (seconds % 9000) + 1000
		//seed = time.Now().UnixNano()/1000000 - 1713220000000
		//one day 86400 seconds
		pr.currentValue = strconv.FormatInt(seed, 10) + " (No input or error input use random seed)"

	} else {
		pr.currentValue = strconv.FormatInt(seed, 10)
	}
	pr.answerClicked = false
	source := rand.NewSource(seed)
	pr.rand = rand.New(source)
	pr.theLevel = seed % 10
	pr.question1 = pr.pr1()
	pr.seedClicked = true
	pr.Update()
}

func (pr *Proofs) finishClicked(ctx app.Context, e app.Event) {
	pr.answerClicked = true
	pr.Q1input = app.Window().GetElementByID("Q1input").Get("value").String()
	if app.Window().GetElementByID("option-a").Get("checked").Bool() {
		pr.Q7option = "Option A"
	} else if app.Window().GetElementByID("option-b").Get("checked").Bool() {
		pr.Q7option = "Option B"
	}
	pr.Update()
}

func (pr *Proofs) pr1() Question {

	statement := ""
	answer := ""
	randC := pr.rand.Intn(5)
	switch randC {
	case 0:
		statement = "Is there a difference in formality or correctness between an argument structured in English text and an argument written in symbols?"
		answer = "yes"
	case 1:
		statement = "For any integer n, the number n^2 − 1 is divisible by 4 if and only if n is an odd integer."

	}
	theStatement := fmt.Sprintf("Is the following statement true or false? %s", statement)

	question := Question{
		statement1: theStatement,
		answer:     answer,
	}
	return question
}
