package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"math/rand"
	"strconv"
	"time"
)

type Logic struct {
	app.Compo
	rand          *rand.Rand
	inputValue    string
	seed          int64
	Q1input       string
	Q2input       string
	Q3input       string
	Q4input       string
	Q5input       string
	Q6input       string
	Q1option      string
	Q2option      string
	Q3option      string
	Q4option      string
	Q5option      string
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

func (lo *Logic) Render() app.UI {
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
						Style("margin-top", "4vh").
						//Style("text-align", "center").
						Text("Instructions HelloMath!"),
				),
				app.Div().Body(
					app.Input().
						ID("seedInput").
						Type("text").
						Value(lo.inputValue).
						OnChange(lo.OnInputChange),
					app.Button().
						Text("Generate Seed").
						OnClick(lo.GenerateSeed),
					app.Button().
						Text("up").
						OnClick(lo.upLevel),

					app.Button().
						Text("down").
						OnClick(lo.downLevel),
				),

				app.H6().Text("Seed Value: "+lo.currentValue),
				app.Div().Style("height", "3vh").Body(),
				app.If(lo.seedClicked,
					//question 1
					///////////
					app.If(lo.question1.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q1 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(lo.question1.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q1input").
									Type("text").
									Value(lo.Q1input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(lo.answerClicked,
								app.If(lo.Q1input == lo.question1.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(lo.answerClicked,
							app.P().Text("Answer: "+lo.question1.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(lo.question1.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q1 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(lo.question1.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Div().Class("form-check").Style("margin-left", "50px").Body(
									app.Input().
										Type("radio").
										Name("options1").
										Class("form-check-input").
										Value("Option A").
										ID("Q1option-a"),
									app.Label().Class("form-check-label").Text("True"),
								),
							),
							app.If(lo.answerClicked,
								app.If(lo.Q1option == lo.question1.answer,
									app.Div().Class("col-auto").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.Div().Class("form-check").Style("margin-left", "50px").Body(
							app.Input().
								Type("radio").
								Name("options1").
								Class("form-check-input").
								Value("Option B").
								ID("Q1option-b"),
							app.Label().Class("form-check-label").For("option-b").Text("False"),
						),
						app.If(lo.answerClicked,
							app.P().Text("Answer: "+lo.question1.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					////question 2
					app.If(lo.question2.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(lo.question2.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q2input").
									Type("text").
									Value(lo.Q2input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(lo.answerClicked,
								app.If(lo.Q2input == lo.question2.answer,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(lo.answerClicked,
							app.P().Text("Answer: "+lo.question2.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(lo.question2.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(lo.question2.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Div().Class("form-check").Style("margin-left", "50px").Body(
									app.Input().
										Type("radio").
										Name("options2").
										Class("form-check-input").
										Value("Option A").
										ID("Q2option-a"),
									app.Label().Class("form-check-label").Text("True"),
								),
							),
							app.If(lo.answerClicked,
								app.If(lo.Q2option == lo.question2.answer,
									app.Div().Class("col-auto").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.Div().Class("form-check").Style("margin-left", "50px").Body(
							app.Input().
								Type("radio").
								Name("options2").
								Class("form-check-input").
								Value("Option B").
								ID("Q2option-b"),
							app.Label().Class("form-check-label").For("option-b").Text("False"),
						),
						app.If(lo.answerClicked,
							app.P().Text("Answer: "+lo.question2.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					///////3
					app.If(lo.question3.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(lo.question3.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q3input").
									Type("text").
									Value(lo.Q3input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(lo.answerClicked,
								app.If(lo.Q3input == lo.question3.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(lo.answerClicked,
							app.P().Text("Answer: "+lo.question3.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(lo.question3.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(lo.question3.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Div().Class("form-check").Style("margin-left", "50px").Body(
									app.Input().
										Type("radio").
										Name("options3").
										Class("form-check-input").
										Value("Option A").
										ID("Q3option-a"),
									app.Label().Class("form-check-label").Text("True"),
								),
							),
							app.If(lo.answerClicked,
								app.If(lo.Q3option == lo.question3.answer,
									app.Div().Class("col-auto").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.Div().Class("form-check").Style("margin-left", "50px").Body(
							app.Input().
								Type("radio").
								Name("options3").
								Class("form-check-input").
								Value("Option B").
								ID("Q3option-b"),
							app.Label().Class("form-check-label").For("option-b").Text("False"),
						),
						app.If(lo.answerClicked,
							app.P().Text("Answer: "+lo.question3.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					//////4
					app.If(lo.question4.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q4 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(lo.question4.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q4input").
									Type("text").
									Value(lo.Q4input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(lo.answerClicked,
								app.If(lo.Q4input == lo.question4.answer,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(lo.answerClicked,
							app.P().Text("Answer: "+lo.question4.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(lo.question4.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q4 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(lo.question4.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Div().Class("form-check").Style("margin-left", "50px").Body(
									app.Input().
										Type("radio").
										Name("options4").
										Class("form-check-input").
										Value("Option A").
										ID("Q4option-a"),
									app.Label().Class("form-check-label").Text("True"),
								),
							),
							app.If(lo.answerClicked,
								app.If(lo.Q4option == lo.question4.answer,
									app.Div().Class("col-auto").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.Div().Class("form-check").Style("margin-left", "50px").Body(
							app.Input().
								Type("radio").
								Name("options4").
								Class("form-check-input").
								Value("Option B").
								ID("Q4option-b"),
							app.Label().Class("form-check-label").For("option-b").Text("False"),
						),
						app.If(lo.answerClicked,
							app.P().Text("Answer: "+lo.question4.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					///////5
					app.If(lo.question5.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q5 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(lo.question5.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q5input").
									Type("text").
									Value(lo.Q5input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(lo.answerClicked,
								app.If(lo.Q5input == lo.question5.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(lo.answerClicked,
							app.P().Text("Answer: "+lo.question5.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(lo.question5.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q5 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(lo.question5.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Div().Class("form-check").Style("margin-left", "50px").Body(
									app.Input().
										Type("radio").
										Name("options5").
										Class("form-check-input").
										Value("Option A").
										ID("Q5option-a"),
									app.Label().Class("form-check-label").Text("True"),
								),
							),
							app.If(lo.answerClicked,
								app.If(lo.Q5option == lo.question5.answer,
									app.Div().Class("col-auto").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.Div().Class("form-check").Style("margin-left", "50px").Body(
							app.Input().
								Type("radio").
								Name("options5").
								Class("form-check-input").
								Value("Option B").
								ID("Q5option-b"),
							app.Label().Class("form-check-label").For("option-b").Text("False"),
						),
						app.If(lo.answerClicked,
							app.P().Text("Answer: "+lo.question5.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),

					app.Button().
						Text("Finish").
						Style("margin-left", "50px").
						OnClick(lo.finishClicked),
				),
			),
		),
	)
}

func (lo *Logic) OnInputChange(ctx app.Context, e app.Event) {
	lo.inputValue = ctx.JSSrc().Get("value").String()
	lo.Update()
}
func (lo *Logic) GenerateSeed(ctx app.Context, e app.Event) {
	inputValue := app.Window().GetElementByID("seedInput").Get("value").String()
	fmt.Println("input value:", inputValue)
	seed, err := strconv.ParseInt(inputValue, 10, 64)
	lo.seed = seed
	if err != nil {
		fmt.Println("No input or error converting string to integer:", err)
		seconds := time.Now().UnixNano() / 1000000
		lo.seed = (seconds % 9000) + 1000
		//seed = time.Now().UnixNano()/1000000 - 1713220000000
		//one day 86400 seconds
		lo.currentValue = strconv.FormatInt(lo.seed, 10) + " (No input or error input use random seed)"

	} else {
		lo.currentValue = strconv.FormatInt(lo.seed, 10)
	}

	lo.answerClicked = false
	source := rand.NewSource(lo.seed)
	lo.rand = rand.New(source)
	group1 := []func() Question{
		lo.lo1,
		lo.lo2,
		lo.lo3,
		lo.lo4,
		lo.lo5,
	}

	lo.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})

	lo.theLevel = lo.seed % 10
	lo.question1 = group1[0]()
	lo.question2 = group1[1]()
	lo.question3 = group1[2]()
	lo.question4 = group1[3]()
	lo.question5 = group1[4]()

	lo.seedClicked = true
	lo.Update()
}

func (lo *Logic) upLevel(ctx app.Context, e app.Event) {
	if lo.seed%10 == 9 {
		lo.seed = lo.seed + 10
	} else {
		lo.seed++
	}

	lo.currentValue = strconv.FormatInt(lo.seed, 10)
	lo.inputValue = strconv.FormatInt(lo.seed, 10)

	lo.answerClicked = false
	source := rand.NewSource(lo.seed)
	lo.rand = rand.New(source)
	group1 := []func() Question{
		lo.lo1,
		lo.lo2,
		lo.lo3,
		lo.lo4,
		lo.lo5,
	}

	lo.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})

	lo.theLevel = lo.seed % 10
	lo.question1 = group1[0]()
	lo.question2 = group1[1]()
	lo.question3 = group1[2]()
	lo.question4 = group1[3]()
	lo.question5 = group1[4]()

	lo.seedClicked = true
	lo.Update()
}
func (lo *Logic) downLevel(ctx app.Context, e app.Event) {
	if lo.seed%10 == 0 {
		lo.seed = lo.seed - 10
	} else {
		lo.seed--
	}

	lo.currentValue = strconv.FormatInt(lo.seed, 10)
	lo.inputValue = strconv.FormatInt(lo.seed, 10)

	lo.answerClicked = false
	source := rand.NewSource(lo.seed)
	lo.rand = rand.New(source)
	group1 := []func() Question{
		lo.lo1,
		lo.lo2,
		lo.lo3,
		lo.lo4,
		lo.lo5,
	}

	lo.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})

	lo.theLevel = lo.seed % 10
	lo.question1 = group1[0]()
	lo.question2 = group1[1]()
	lo.question3 = group1[2]()
	lo.question4 = group1[3]()
	lo.question5 = group1[4]()

	lo.seedClicked = true

	lo.Update()
}

func (lo *Logic) finishClicked(ctx app.Context, e app.Event) {
	lo.answerClicked = true
	//q1
	if lo.question1.questionType == 1 {
		lo.Q1input = app.Window().GetElementByID("Q1input").Get("value").String()
	}
	if lo.question1.questionType == 2 {
		if app.Window().GetElementByID("Q1option-a").Get("checked").Bool() {
			lo.Q1option = "true"
		} else if app.Window().GetElementByID("Q1option-b").Get("checked").Bool() {
			lo.Q1option = "false"
		}
	}
	//q2
	if lo.question2.questionType == 1 {
		lo.Q2input = app.Window().GetElementByID("Q2input").Get("value").String()
	}
	if lo.question2.questionType == 2 {
		if app.Window().GetElementByID("Q2option-a").Get("checked").Bool() {
			lo.Q2option = "true"
		} else if app.Window().GetElementByID("Q2option-b").Get("checked").Bool() {
			lo.Q2option = "false"
		}
	}
	//q3
	if lo.question3.questionType == 1 {
		lo.Q3input = app.Window().GetElementByID("Q3input").Get("value").String()
	}
	if lo.question3.questionType == 2 {
		if app.Window().GetElementByID("Q3option-a").Get("checked").Bool() {
			lo.Q3option = "true"
		} else if app.Window().GetElementByID("Q3option-b").Get("checked").Bool() {
			lo.Q3option = "false"
		}
	}
	//q4
	if lo.question4.questionType == 1 {
		lo.Q4input = app.Window().GetElementByID("Q4input").Get("value").String()
	}
	if lo.question4.questionType == 2 {
		if app.Window().GetElementByID("Q4option-a").Get("checked").Bool() {
			lo.Q4option = "true"
		} else if app.Window().GetElementByID("Q4option-b").Get("checked").Bool() {
			lo.Q4option = "false"
		}
	}
	//q5
	if lo.question5.questionType == 1 {
		lo.Q5input = app.Window().GetElementByID("Q5input").Get("value").String()
	}
	if lo.question5.questionType == 2 {
		if app.Window().GetElementByID("Q5option-a").Get("checked").Bool() {
			lo.Q5option = "true"
		} else if app.Window().GetElementByID("Q5option-b").Get("checked").Bool() {
			lo.Q5option = "false"
		}
	}
	lo.Update()
}

func (lo *Logic) lo1() Question {

	statement := ""
	answer := ""
	randC := lo.rand.Intn(7)
	switch randC {
	case 0:
		statement = "If a = true, b = true, what is a→b ?"
		answer = "true"
	case 1:
		statement = "If a = true, b = false, what is a→b ?"
		answer = "false"
	case 2:
		statement = "If a = false, b = true, what is a→b ?"
		answer = "true"
	case 3:
		statement = "If a = false, b = false, what is a↔b ?"
		answer = "true"
	case 4:
		statement = "If a = true, b = true, what is a↔b ?"
		answer = "true"
	case 5:
		statement = "If a = true, b = false, what is a↔b ?"
		answer = "false"
	case 6:
		statement = "If a = false, b = true, what is a↔b ?"
		answer = "false"
	default:
		statement = "If a = false, b = false, what is a→b ?"
		answer = "true"
	}

	question := Question{
		statement1:   statement,
		answer:       answer,
		questionType: 2,
	}
	return question
}

func (lo *Logic) lo2() Question {
	statement := ""
	answer := ""
	randC := lo.rand.Intn(7)
	switch randC {
	case 0:
		statement = "If a = true, b = true, what is a∧b ?"
		answer = "true"
	case 1:
		statement = "If a = true, b = false, what is a∧b ?"
		answer = "false"
	case 2:
		statement = "If a = false, b = true, what is a∧b ?"
		answer = "false"
	case 3:
		statement = "If a = false, b = false, what is a∧b ?"
		answer = "false"
	case 4:
		statement = "If a = true, b = true, what is a∨b ?"
		answer = "true"
	case 5:
		statement = "If a = true, b = false, what is a∨b ?"
		answer = "true"
	case 6:
		statement = "If a = false, b = true, what is a∨b ?"
		answer = "true"
	case 7:
		statement = "If a = false, b = false, what is a∨b ?"
		answer = "false"
	case 8:
		statement = "If a = false, b = false, what is a⊕b ?"
		answer = "false"
	case 9:
		statement = "If a = true, b = true, what is a⊕b ?"
		answer = "false"
	case 10:
		statement = "If a = true, b = false, what is a⊕b ?"
		answer = "true"
	default:
		statement = "If a = false, b = true, what is a⊕b ?"
		answer = "true"
	}
	return Question{
		statement1:   statement,
		answer:       answer,
		questionType: 2,
	}
}

func trueTableCheck(a int, b int, c string) int {
	var result int
	switch c {
	case "→":
		if a == 1 && b == 1 {
			result = 1
		}
		if a == 1 && b == 0 {
			result = 0
		}
		if a == 0 && b == 1 {
			result = 1
		} else {
			result = 1
		}
	case "↔":
		if a == 1 && b == 1 {
			result = 1
		}
		if a == 1 && b == 0 {
			result = 0
		}
		if a == 0 && b == 1 {
			result = 0
		} else {
			result = 1
		}
	case "∧":
		if a == 1 && b == 1 {
			result = 1
		}
		if a == 1 && b == 0 {
			result = 0
		}
		if a == 0 && b == 1 {
			result = 0
		} else {
			result = 0
		}
	case "∨":
		if a == 1 && b == 1 {
			result = 1
		}
		if a == 1 && b == 0 {
			result = 1
		}
		if a == 0 && b == 1 {
			result = 1
		} else {
			result = 0
		}
	case "⊕":
		if a == 1 && b == 1 {
			result = 0
		}
		if a == 1 && b == 0 {
			result = 1
		}
		if a == 0 && b == 1 {
			result = 1
		} else {
			result = 0
		}
	default:
		result = 0
	}
	return result
}

func negation(a int) int {
	if a == 1 {
		return 0
	}
	return 1
}

func (lo *Logic) lo3() Question {
	//var propositions = []string{"p", "q", "r", "s"}
	var operators = []string{"→", "↔", "∧", "∨", "⊕"}
	// Generate two random logical expressions
	randProposition1 := "p"
	randProposition2 := "q"
	randProposition3 := "r"
	randProposition4 := "s"
	randProposition5 := "z"
	randProposition6 := "y"
	randOperator1 := operators[lo.rand.Intn(len(operators))]
	randOperator2 := operators[lo.rand.Intn(len(operators))]
	randOperator3 := operators[lo.rand.Intn(len(operators))]
	randOperator4 := operators[lo.rand.Intn(len(operators))]
	randOperator5 := operators[lo.rand.Intn(len(operators))]
	//randOperator6 := operators[lo.rand.Intn(len(operators))]

	a := lo.rand.Intn(2)
	var aTruth string
	if a == 0 {
		aTruth = "false"
	} else {
		aTruth = "true"
	}
	b := lo.rand.Intn(2)
	var bTruth string
	if b == 0 {
		bTruth = "false"
	} else {
		bTruth = "true"
	}
	c := lo.rand.Intn(2)
	var cTruth string
	if c == 0 {
		cTruth = "false"
	} else {
		cTruth = "true"
	}
	d := lo.rand.Intn(2)
	var dTruth string
	if d == 0 {
		dTruth = "false"
	} else {
		dTruth = "true"
	}
	e := lo.rand.Intn(2)
	var eTruth string
	if e == 0 {
		eTruth = "false"
	} else {
		eTruth = "true"
	}
	var fTruth string
	f := lo.rand.Intn(2)
	if f == 0 {
		fTruth = "false"
	} else {
		fTruth = "true"
	}
	expression1 := ""
	expression2 := ""
	var result1 int

	randC := lo.rand.Intn(7)
	switch randC {
	case 0:
		//(p→q)∧(r→s)
		expression1 = "(" + randProposition1 + randOperator1 + randProposition2 + ")" + randOperator2 + "(" + randProposition3 + randOperator3 + randProposition4 + ")"
		expression2 = randProposition1 + " = " + aTruth + ", " + randProposition2 + " = " + bTruth + ", " + randProposition3 + " = " + cTruth + ", " + randProposition4 + " = " + dTruth
		result1 = trueTableCheck(trueTableCheck(a, b, randOperator1), trueTableCheck(c, d, randOperator3), randOperator2)
	case 1:
		//((p↔q)∧(r↔s))∧(p∨q)
		expression1 = "((" + randProposition1 + randOperator1 + randProposition2 + ")" + randOperator2 + "(" + randProposition3 + randOperator3 + randProposition4 + "))" + randOperator4 + "(" + randProposition5 + randOperator5 + randProposition6 + ")"
		expression2 = randProposition1 + " = " + aTruth + ", " + randProposition2 + " = " + bTruth + ", " + randProposition3 + " = " + cTruth + ", " + randProposition4 + " = " + dTruth + ", " + randProposition5 + " = " + eTruth + ", " + randProposition6 + " = " + fTruth
		result1 = trueTableCheck(trueTableCheck(trueTableCheck(a, b, randOperator1), trueTableCheck(c, d, randOperator3), randOperator2), trueTableCheck(e, f, randOperator5), randOperator4)
	case 2:
		//(p∧q)∨((r∧s)∧(p∨q))
		expression1 = "(" + randProposition1 + randOperator1 + randProposition2 + ")" + randOperator2 + "((" + randProposition3 + randOperator3 + randProposition4 + ")" + randOperator4 + "(" + randProposition5 + randOperator5 + randProposition6 + "))"
		expression2 = randProposition1 + " = " + aTruth + ", " + randProposition2 + " = " + bTruth + ", " + randProposition3 + " = " + cTruth + ", " + randProposition4 + " = " + dTruth + ", " + randProposition5 + " = " + eTruth + ", " + randProposition6 + " = " + fTruth
		result1 = trueTableCheck(trueTableCheck(a, b, randOperator1), trueTableCheck(trueTableCheck(c, d, randOperator3), trueTableCheck(e, f, randOperator5), randOperator4), randOperator2)
	case 3:
		//((p∧q)∨(r∧s))∧(p∨q)
		expression1 = "¬(" + randProposition1 + randOperator1 + randProposition2 + ")" + randOperator2 + "(" + randProposition3 + randOperator3 + randProposition4 + ")"
		expression2 = randProposition1 + " = " + aTruth + ", " + randProposition2 + " = " + bTruth + ", " + randProposition3 + " = " + cTruth + ", " + randProposition4 + " = " + dTruth
		result1 = trueTableCheck(negation(trueTableCheck(a, b, randOperator1)), trueTableCheck(c, d, randOperator3), randOperator2)
	case 4:
		expression1 = "(¬(" + randProposition1 + randOperator1 + randProposition2 + ")" + randOperator2 + "(" + randProposition3 + randOperator3 + randProposition4 + "))" + randOperator4 + "(" + randProposition5 + randOperator5 + randProposition6 + ")"
		expression2 = randProposition1 + " = " + aTruth + ", " + randProposition2 + " = " + bTruth + ", " + randProposition3 + " = " + cTruth + ", " + randProposition4 + " = " + dTruth + ", " + randProposition5 + " = " + eTruth + ", " + randProposition6 + " = " + fTruth
		result1 = trueTableCheck(trueTableCheck(negation(trueTableCheck(a, b, randOperator1)), trueTableCheck(c, d, randOperator3), randOperator2), trueTableCheck(e, f, randOperator5), randOperator4)
	default:
		expression1 = "¬((" + randProposition1 + randOperator1 + randProposition2 + ")" + randOperator2 + "(" + randProposition3 + randOperator3 + randProposition4 + "))" + randOperator4 + "(" + randProposition5 + randOperator5 + randProposition6 + ")"
		expression2 = randProposition1 + " = " + aTruth + ", " + randProposition2 + " = " + bTruth + ", " + randProposition3 + " = " + cTruth + ", " + randProposition4 + " = " + dTruth + ", " + randProposition5 + " = " + eTruth + ", " + randProposition6 + " = " + fTruth
		result1 = trueTableCheck(negation(trueTableCheck(trueTableCheck(a, b, randOperator1), trueTableCheck(c, d, randOperator3), randOperator2)), trueTableCheck(e, f, randOperator5), randOperator4)

	}
	theAnswer := ""
	if result1 == 1 {
		theAnswer = "true"
	} else {
		theAnswer = "false"
	}

	statement := fmt.Sprintf("If %v, what is the value of %v", expression2, expression1)

	return Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
	}
}

func (lo *Logic) lo4() Question {
	randC := lo.rand.Intn(5)
	sentence := ""
	theAnswer := ""
	switch randC {
	case 0:
		sentence = "'it will not rain or it will not be windy' and 'it will not rain and it will not be windy'"
		theAnswer = "false"
	case 1:
		sentence = "'it will rain or it will be windy' and 'it will rain and it will be windy'"
		theAnswer = "false"
	case 2:
		sentence = "'it will rain or it will be windy' and 'it will rain and it will be windy'"
		theAnswer = "false"
	case 3:
		sentence = "'it will not rain or it will be windy' and 'it will not rain and it will be windy'"
		theAnswer = "false"
	default:
		sentence = "'it will rain or it will be windy' and 'it will be windy or it will rain'"
		theAnswer = "true"

	}

	statement := fmt.Sprintf("Are the statements, %v logically equivalent?", sentence)
	return Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
	}
}

func (lo *Logic) lo5() Question {
	randC := lo.rand.Intn(9)
	sentence := ""
	theAnswer := ""
	switch randC {
	case 0:
		sentence = "P ∧ P to true"
		theAnswer = "false"
	case 1:
		sentence = "P ∨ P to true"
		theAnswer = "false"
	case 2:
		sentence = "P∧(Q∨R) to P∧Q∨P∧R"
		theAnswer = "true"
	case 3:
		sentence = "¬(¬(P∧Q)) to P∨Q"
		theAnswer = "false"
	case 4:
		sentence = "(P∨Q)∨R to P∨Q∨R"
		theAnswer = "true"
	case 5:
		sentence = "P→Q to ¬P∨Q"
		theAnswer = "true"
	case 6:
		sentence = "P∨(P∧Q) to P"
		theAnswer = "true"
	case 7:
		sentence = "(P∧Q)∨(P∧¬Q) to P"
		theAnswer = "true"
	default:
		sentence = "P∧P to P"
		theAnswer = "true"
	}
	//var operators = []string{"→", "↔", "∧", "∨", "⊕" ¬}

	statement := fmt.Sprintf("Determine if the simplification of the following expression is correct: %v", sentence)
	return Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
	}
}
