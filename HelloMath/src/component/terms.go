package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Terms struct {
	app.Compo
	rand       *rand.Rand
	inputValue string
	seed       int64
	Q1input    string
	Q2input    string
	Q3input    string
	Q4input    string
	Q5input    string
	Q6input    string
	Q1option   string
	Q2option   string
	Q3option   string
	Q4option   string
	Q5option   string

	question1     Question
	question2     Question
	question3     Question
	question4     Question
	question5     Question
	question6     Question
	questionType  int
	theLevel      int64
	seedClicked   bool
	answerClicked bool
	currentValue  string
}

func (te *Terms) Render() app.UI {
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
						Value(te.inputValue).
						OnChange(te.OnInputChange),
					app.Button().
						Text("Generate Seed").
						OnClick(te.GenerateSeed),
					app.Button().
						Text("up").
						OnClick(te.upLevel),

					app.Button().
						Text("down").
						OnClick(te.downLevel),
				),

				app.H6().Text("Seed Value: "+te.currentValue),
				app.Div().Style("height", "3vh").Body(),
				app.If(te.seedClicked,
					//question 1
					///////////
					app.If(te.question1.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q1 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(te.question1.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q1input").
									Type("text").
									Value(te.Q1input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(te.answerClicked,
								app.If(te.Q1input == te.question1.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(te.answerClicked,
							app.P().Text("Answer: "+te.question1.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(te.question1.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q1 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(te.question1.statement1),
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
							app.If(te.answerClicked,
								app.If(te.Q1option == te.question1.answer,
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
						app.If(te.answerClicked,
							app.P().Text("Answer: "+te.question1.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					////question 2
					app.If(te.question2.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(te.question2.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q2input").
									Type("text").
									Value(te.Q2input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(te.answerClicked,
								app.If(te.Q2input == te.question2.answer,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(te.answerClicked,
							app.P().Text("Answer: "+te.question2.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(te.question2.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(te.question2.statement1),
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
							app.If(te.answerClicked,
								app.If(te.Q2option == te.question2.answer,
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
						app.If(te.answerClicked,
							app.P().Text("Answer: "+te.question2.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					///////3
					app.If(te.question3.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(te.question3.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q3input").
									Type("text").
									Value(te.Q3input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(te.answerClicked,
								app.If(te.Q3input == te.question3.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(te.answerClicked,
							app.P().Text("Answer: "+te.question3.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(te.question3.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(te.question3.statement1),
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
							app.If(te.answerClicked,
								app.If(te.Q3option == te.question3.answer,
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
						app.If(te.answerClicked,
							app.P().Text("Answer: "+te.question3.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					//////4
					app.If(te.question4.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q4 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(te.question4.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q4input").
									Type("text").
									Value(te.Q4input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(te.answerClicked,
								app.If(te.Q4input == te.question4.answer,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(te.answerClicked,
							app.P().Text("Answer: "+te.question4.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(te.question4.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q4 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(te.question4.statement1),
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
							app.If(te.answerClicked,
								app.If(te.Q4option == te.question4.answer,
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
						app.If(te.answerClicked,
							app.P().Text("Answer: "+te.question4.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					///////5
					app.If(te.question5.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q5 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(te.question5.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q5input").
									Type("text").
									Value(te.Q5input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(te.answerClicked,
								app.If(te.Q5input == te.question5.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(te.answerClicked,
							app.P().Text("Answer: "+te.question5.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(te.question5.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q5 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(te.question5.statement1),
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
							app.If(te.answerClicked,
								app.If(te.Q5option == te.question5.answer,
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
						app.If(te.answerClicked,
							app.P().Text("Answer: "+te.question5.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),

					app.Button().
						Text("Finish").
						Style("margin-left", "50px").
						OnClick(te.finishClicked),
				),
			),
		),
	)
}

func (te *Terms) OnInputChange(ctx app.Context, e app.Event) {
	te.inputValue = ctx.JSSrc().Get("value").String()
	te.Update()
}

func (te *Terms) GenerateSeed(ctx app.Context, e app.Event) {
	inputValue := app.Window().GetElementByID("seedInput").Get("value").String()
	fmt.Println("input value:", inputValue)
	seed, err := strconv.ParseInt(inputValue, 10, 64)
	te.seed = seed
	if err != nil {
		fmt.Println("No input or error converting string to integer:", err)
		seconds := time.Now().UnixNano() / 1000000
		te.seed = (seconds % 9000) + 1000
		//seed = time.Now().UnixNano()/1000000 - 1713220000000
		//one day 86400 seconds
		te.currentValue = strconv.FormatInt(te.seed, 10) + " (No input or error input use random seed)"

	} else {
		te.currentValue = strconv.FormatInt(te.seed, 10)
	}

	te.answerClicked = false
	source := rand.NewSource(te.seed)
	te.rand = rand.New(source)
	group1 := []func() Question{
		te.Te1,
		te.te3,
		te.te4,
		te.te8,
	}
	group2 := []func() Question{
		te.te7,
		te.te6,
		te.te5,
	}
	te.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})
	te.rand.Shuffle(len(group2), func(i, j int) {
		group2[i], group2[j] = group2[j], group2[i]
	})
	te.theLevel = te.seed % 10
	if te.theLevel >= 0 && te.theLevel < 5 {
		te.question1 = group1[0]()
		te.question2 = group1[1]()
		te.question3 = group1[2]()
		te.question4 = group1[3]()
		te.question5 = group2[0]()
	} else {
		te.question1 = group1[0]()
		te.question2 = group1[1]()
		te.question3 = group2[0]()
		te.question4 = group2[1]()
		te.question5 = group2[2]()
	}
	te.seedClicked = true
	te.Update()
}

func (te *Terms) upLevel(ctx app.Context, e app.Event) {
	if te.seed%10 == 9 {
		te.seed = te.seed + 10
	} else {
		te.seed++
	}

	te.currentValue = strconv.FormatInt(te.seed, 10)
	te.inputValue = strconv.FormatInt(te.seed, 10)

	te.answerClicked = false
	source := rand.NewSource(te.seed)
	te.rand = rand.New(source)
	group1 := []func() Question{
		te.Te1,
		te.te3,
		te.te4,
		te.te8,
	}
	group2 := []func() Question{
		te.te7,
		te.te6,
		te.te5,
	}
	te.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})
	te.rand.Shuffle(len(group2), func(i, j int) {
		group2[i], group2[j] = group2[j], group2[i]
	})
	te.theLevel = te.seed % 10
	if te.theLevel >= 0 && te.theLevel < 5 {
		te.question1 = group1[0]()
		te.question2 = group1[1]()
		te.question3 = group1[2]()
		te.question4 = group1[3]()
		te.question5 = group2[0]()
	} else {
		te.question1 = group1[0]()
		te.question2 = group1[1]()
		te.question3 = group2[0]()
		te.question4 = group2[1]()
		te.question5 = group2[2]()
	}

	te.seedClicked = true
	te.Update()
}
func (te *Terms) downLevel(ctx app.Context, e app.Event) {
	if te.seed%10 == 0 {
		te.seed = te.seed - 10
	} else {
		te.seed--
	}

	te.currentValue = strconv.FormatInt(te.seed, 10)
	te.inputValue = strconv.FormatInt(te.seed, 10)

	te.answerClicked = false
	source := rand.NewSource(te.seed)
	te.rand = rand.New(source)
	group1 := []func() Question{
		te.Te1,
		te.te3,
		te.te4,
		te.te8,
	}
	group2 := []func() Question{
		te.te7,
		te.te6,
		te.te5,
	}
	te.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})
	te.rand.Shuffle(len(group2), func(i, j int) {
		group2[i], group2[j] = group2[j], group2[i]
	})
	te.theLevel = te.seed % 10
	if te.theLevel >= 0 && te.theLevel < 5 {
		te.question1 = group1[0]()
		te.question2 = group1[1]()
		te.question3 = group1[2]()
		te.question4 = group1[3]()
		te.question5 = group2[0]()
	} else {
		te.question1 = group1[0]()
		te.question2 = group1[1]()
		te.question3 = group2[0]()
		te.question4 = group2[1]()
		te.question5 = group2[2]()
	}

	te.seedClicked = true
	te.Update()
}

func (te *Terms) finishClicked(ctx app.Context, e app.Event) {
	te.answerClicked = true
	//q1
	if te.question1.questionType == 1 {
		te.Q1input = app.Window().GetElementByID("Q1input").Get("value").String()
	}
	if te.question1.questionType == 2 {
		if app.Window().GetElementByID("Q1option-a").Get("checked").Bool() {
			te.Q1option = "true"
		} else if app.Window().GetElementByID("Q1option-b").Get("checked").Bool() {
			te.Q1option = "false"
		}
	}
	//q2
	if te.question2.questionType == 1 {
		te.Q2input = app.Window().GetElementByID("Q2input").Get("value").String()
	}
	if te.question2.questionType == 2 {
		if app.Window().GetElementByID("Q2option-a").Get("checked").Bool() {
			te.Q2option = "true"
		} else if app.Window().GetElementByID("Q2option-b").Get("checked").Bool() {
			te.Q2option = "false"
		}
	}
	//q3
	if te.question3.questionType == 1 {
		te.Q3input = app.Window().GetElementByID("Q3input").Get("value").String()
	}
	if te.question3.questionType == 2 {
		if app.Window().GetElementByID("Q3option-a").Get("checked").Bool() {
			te.Q3option = "true"
		} else if app.Window().GetElementByID("Q3option-b").Get("checked").Bool() {
			te.Q3option = "false"
		}
	}
	//q4
	if te.question4.questionType == 1 {
		te.Q4input = app.Window().GetElementByID("Q4input").Get("value").String()
	}
	if te.question4.questionType == 2 {
		if app.Window().GetElementByID("Q4option-a").Get("checked").Bool() {
			te.Q4option = "true"
		} else if app.Window().GetElementByID("Q4option-b").Get("checked").Bool() {
			te.Q4option = "false"
		}
	}
	//q5
	if te.question5.questionType == 1 {
		te.Q5input = app.Window().GetElementByID("Q5input").Get("value").String()
	}
	if te.question5.questionType == 2 {
		if app.Window().GetElementByID("Q5option-a").Get("checked").Bool() {
			te.Q5option = "true"
		} else if app.Window().GetElementByID("Q5option-b").Get("checked").Bool() {
			te.Q5option = "false"
		}
	}
	te.Update()
}

func (te *Terms) Te1() Question {
	//The set Zn contains natural number less than
	theN := te.rand.Intn(8) + 1
	randC := te.rand.Intn(9)
	statement := ""
	theAnswer := ""

	switch randC {
	case 0:
		statement = fmt.Sprintf("Does the set ℤ%v include the element 0?", theN)
		theAnswer = "true"

	case 1:
		statement = fmt.Sprintf("Does the set ℤ%v include the element %v?", theN, theN)
		theAnswer = "false"

	case 2:
		statement = fmt.Sprintf("Is the number π(pi) in the set Q?")
		theAnswer = "false"
	case 3:
		statement = fmt.Sprintf("Is the number π(pi) in the set R?")
		theAnswer = "true"
	case 4:
		statement = fmt.Sprintf("Is the number %v in the set Q?", theN)
		theAnswer = "true"
	case 5:
		statement = fmt.Sprintf("Is the square root of 2 in the set Q??")
		theAnswer = "false"
	case 6:
		statement = fmt.Sprintf("Is the square root of 3 in the set Q??")
		theAnswer = "false"
	case 7:
		statement = fmt.Sprintf("Is the e (Euler's number) in the set Q?")
		theAnswer = "false"
	case 8:
		statement = fmt.Sprintf("Is the repeating decimal 0.3333... in the set Q?")
		theAnswer = "true"

	default:
		statement = fmt.Sprintf("Does the set ℤn include negative elements?")
		theAnswer = "false"
	}

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}

// multi-choice question
func (te *Terms) te2() Question {
	statement := ""
	Crand := te.rand.Intn(3)

	theAnswer := ""
	switch Crand {
	case 0:
		statement = "Which syntax rule is demonstrated by the expression a+b ⊢ b+a?"
		theAnswer = "commutative law of addition"
	case 1:
		statement = "Which syntax rule is demonstrated by the expression a+(b+c) ⊢ (a+b)+c?"
		theAnswer = "associative law of addition"
	case 2:
		statement = "Which syntax rule is demonstrated by the expression a×1 ⊢ a?"
		theAnswer = "multiplicative identity law"
	default:
		statement = "Which syntax rule is demonstrated by the expression a×(b×c) ⊢ (a×b)×c?"
		theAnswer = "associative law of multiplication"
	}

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 3,
		level:        1,
	}
	return question
}

func (te *Terms) te3() Question {
	a := te.rand.Intn(30) + 1
	b := te.rand.Intn(30) + 1
	c := te.rand.Intn(30) + 1
	d := te.rand.Intn(30) + 1

	statement := fmt.Sprintf("Given the syntax rule for the Associative Law of addition: (a+b)+c ⊢ a+(b+c), perform a derivation to simplify the expression ((%v+%v)+%v)+%v?", a, b, c, d)
	theAnswer := fmt.Sprintf("%v+(%v+(%v+%v))", a, b, c, d)
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}

func (te *Terms) te4() Question {
	Crand := te.rand.Intn(2)
	statement := ""
	theAnswer := ""

	if Crand == 0 {
		statement = "If terms A, B are syntactically equivalent, A and B are the same in all contexts."
		theAnswer = "false"
	} else {
		statement = "If terms A, B are syntactically equivalent, which mean A can be derived from B and B can be derived from A."
		theAnswer = "true"
	}
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
	}
	return question
}

func (te *Terms) te5() Question {
	randN1 := te.rand.Intn(10) + 1
	randN2 := te.rand.Intn(10) + 1
	C := randN2 + 1
	CFalse := randN2 + te.rand.Intn(4) + 2
	randTf := te.rand.Intn(3)
	statement := ""
	theAnswer := ""
	if randTf == 0 {
		statement = fmt.Sprintf("If given the syntax rules: x ⊢ x + %v, x + %v ⊢ x. Determine if the terms %v and terms %v are syntactically equivalent.", randN1, randN1, randN2, CFalse)
		theAnswer = "false"
	} else {
		statement = fmt.Sprintf("If given the syntax rules: x ⊢ x + %v, x + %v ⊢ x. Determine if the terms %v and terms %v are syntactically equivalent.", randN1, randN1, randN2, C)
		theAnswer = "true"
	}
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
	}
	return question
}

func (te *Terms) te6() Question {
	C := te.rand.Intn(10)
	expression := ""
	theAnswer := ""
	switch C {
	case 0:
		expression = "a+(b+c)"
		theAnswer = "true"
	case 1:
		expression = "(a*(b+c))"
		theAnswer = "true"
	case 2:
		expression = "a+(b*c)"
		theAnswer = "false"
	case 3:
		expression = "(a+b)*(c+d)"
		theAnswer = "false"
	default:
		expression = "(a×b)×c"
		theAnswer = "true"
	}
	statement := fmt.Sprintf("Does the expression %v contain any unnecessary brackets?", expression)
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
	}
	return question

}

func (te *Terms) te7() Question {
	operators := []string{"+", "-", "*", "/"}
	var rpnExpression string
	var stack []string
	randomNumber := te.rand.Intn(8) + 1
	randLength := te.rand.Intn(2) + 3
	rpnExpression += strconv.Itoa(randomNumber)
	switch randLength {
	case 3:
		for i := 0; i < 3; i++ {
			randomNumber = te.rand.Intn(8) + 1
			rpnExpression += " " + strconv.Itoa(randomNumber)
			randomOperator := operators[te.rand.Intn(len(operators))]
			rpnExpression += " " + randomOperator
		}
	case 4:
		for i := 0; i < 4; i++ {
			randomNumber = te.rand.Intn(8) + 1
			rpnExpression += " " + strconv.Itoa(randomNumber)
			randomOperator := operators[te.rand.Intn(len(operators))]
			rpnExpression += " " + randomOperator
		}
	default:
		for i := 0; i < 5; i++ {
			randomNumber = te.rand.Intn(8) + 1
			rpnExpression += " " + strconv.Itoa(randomNumber)
			randomOperator := operators[te.rand.Intn(len(operators))]
			rpnExpression += " " + randomOperator
		}

	}
	theRpn := rpnExpression
	tokens := strings.Split(theRpn, " ")
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			infix := fmt.Sprintf("(%s %s %s)", operand1, token, operand2)
			stack = append(stack, infix)
		default:
			stack = append(stack, token)
		}
	}
	theInfixWithBracelet := stack[0]
	theInfix := theInfixWithBracelet[1 : len(theInfixWithBracelet)-1]
	randC := te.rand.Intn(2)
	statement := ""
	theAnswer := ""
	if randC == 0 {
		statement = fmt.Sprintf("Given the RPN expression: %v, convert it to infix notation.", theRpn)
		theAnswer = theInfix
	} else {
		statement = fmt.Sprintf("Given the infix expression: %v, convert it to RPN notation.", theInfix)
		theAnswer = theRpn

	}
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}

func (te *Terms) te8() Question {
	theAnswer := ""
	randC := te.rand.Intn(5)
	theSet := ""
	var theNumber int
	switch randC {
	case 0:
		theSet = "N(Natural Numbers)"
		theNumber = te.rand.Intn(100)
		theAnswer = "true"
	case 1:
		theSet = "N(Natural Numbers)"
		theNumber = -te.rand.Intn(100) - 1
		theAnswer = "false"
	case 2:
		theSet = "Z(integers)"
		theNumber = te.rand.Intn(200) - 100
		theAnswer = "true"
	case 3:
		theSet = "Q(Rational Numbers)"
		theNumber = te.rand.Intn(200) - 100
		theAnswer = "true"
	default:
		theSet = "R(Real Numbers)"
		theNumber = te.rand.Intn(200) - 100
		theAnswer = "true"

	}
	statement := fmt.Sprintf("Does the set %v include the number %v?", theSet, theNumber)
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
	}
	return question
}

//
//func chooseFunction(level int) {
//	switch level {
//	case 1:
//		function1()
//	case 2:
//		function2()
//	case 3:
//		function3()
//	default:
//		fmt.Println("Invalid level")
//	}
//}

//tiankong 1
//xuanze 2
