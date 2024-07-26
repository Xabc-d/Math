package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"math/rand"
	"strconv"
	"time"
)

type Induction struct {
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

func (in *Induction) Render() app.UI {
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
						Value(in.inputValue).
						OnChange(in.OnInputChange),
					app.Button().
						Text("Generate Seed").
						OnClick(in.GenerateSeed),
				),
				app.H5().Text("Seed Value: "+in.currentValue),
				app.Div().Style("height", "3vh").Body(),
				app.If(in.seedClicked,
					//question 1
					app.Div().Class("row").Body(
						app.Div().Class("col-auto").Body(
							app.H6().Text("Q1 :"),
						),
						app.Div().Class("col").Body(
							app.P().Text(in.question1.statement1),
						),
					),
					app.Div().Class("row").Body(
						app.Div().Class("col-auto").Body(
							app.Input().
								ID("Q1input").
								Type("text").
								Value(in.Q1input).
								Style("width", "200px").
								Style("height", "30px").
								Style("margin-left", "50px").
								Class("form-control mb-3").
								Placeholder("").
								AutoFocus(true),
						),
						app.If(in.answerClicked,
							app.If(in.Q1input == in.question1.answer,
								//app.If(se.question1.answer == se.Q1input,
								app.Div().Class("col").Body(
									app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
								),
							),
						),

						app.If(in.answerClicked,
							app.P().Text("Answer: "+in.question1.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
						//question 2

						//
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(in.question2.statement1),
							),
						),
						app.Input().
							ID("seedInput").
							Type("text").
							Value(in.Q2input).
							Style("width", "200px").
							Style("height", "30px").
							Style("margin-left", "50px").
							Class("form-control mb-3").
							Placeholder("").
							AutoFocus(true),

						app.If(in.answerClicked,
							app.P().Text("Answer: "+in.question2.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
						//question 3
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(in.question3.statement1),
							),
						),
						app.Input().
							ID("seedInput").
							Type("text").
							Value(in.Q3input).
							Style("width", "200px").
							Style("height", "30px").
							Style("margin-left", "50px").
							Class("form-control mb-3").
							Placeholder("").
							AutoFocus(true),

						app.If(in.answerClicked,
							app.If(in.question1.answer == in.Q1input,
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
								app.P().Text(in.question4.statement1),
							),
						),

						app.Input().
							ID("seedInput").
							Type("text").
							Value(in.Q4input).
							Style("width", "200px").
							Style("height", "30px").
							Style("margin-left", "50px").
							Class("form-control mb-3").
							Placeholder("").
							AutoFocus(true),

						app.If(in.answerClicked,
							app.If(in.question1.answer == in.Q1input,
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
								app.P().Text(in.question5.statement1),
							),
						),
						app.Input().
							ID("seedInput").
							Type("text").
							Value(in.Q5input).
							Style("width", "200px").
							Style("height", "30px").
							Style("margin-left", "50px").
							Class("form-control mb-3").
							Placeholder("").
							AutoFocus(true),

						app.If(in.answerClicked,
							app.If(in.question1.answer == in.Q1input,
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
								app.P().Text(in.question6.statement1),
							),
						),

						app.Input().
							ID("seedInput").
							Type("text").
							Value(in.Q6input).
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
					app.If(in.answerClicked,
						app.If(in.Q7option == "Option B",
							app.Div().Class("col").Body(
								app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
							),
						),
					),
					app.Button().
						Text("Finish").
						Style("margin-left", "50px").
						OnClick(in.finishClicked),
				),
			),
		),
	)

}

func (in *Induction) OnInputChange(ctx app.Context, e app.Event) {
	in.inputValue = ctx.JSSrc().Get("value").String()
	in.Update()
}
func (in *Induction) GenerateSeed(ctx app.Context, e app.Event) {
	inputValue := app.Window().GetElementByID("seedInput").Get("value").String()
	fmt.Println("input value:", inputValue)
	seed, err := strconv.ParseInt(inputValue, 10, 64)
	if err != nil {
		fmt.Println("No input or error converting string to integer:", err)
		seconds := time.Now().UnixNano() / 1000000
		seed = (seconds % 9000) + 1000
		//seed = time.Now().UnixNano()/1000000 - 1713220000000
		//one day 86400 seconds
		in.currentValue = strconv.FormatInt(seed, 10) + " (No input or error input use random seed)"

	} else {
		in.currentValue = strconv.FormatInt(seed, 10)
	}
	in.answerClicked = false
	//convert seed to string
	source := rand.NewSource(seed)
	in.rand = rand.New(source)
	in.theLevel = seed % 10
	in.question1 = in.in1()
	//se.question1 = se.symmetricSingle()

	in.seedClicked = true
	in.Update()
}

func (in *Induction) finishClicked(ctx app.Context, e app.Event) {
	in.answerClicked = true
	in.Q1input = app.Window().GetElementByID("Q1input").Get("value").String()
	if app.Window().GetElementByID("option-a").Get("checked").Bool() {
		in.Q7option = "Option A"
	} else if app.Window().GetElementByID("option-b").Get("checked").Bool() {
		in.Q7option = "Option B"
	}
	in.Update()
}

func (in *Induction) in1() Question {

	question := Question{
		statement1: "statement",
		answer:     "",
	}
	return question
}

//Which of the following statements is typically proven using mathematical induction?
