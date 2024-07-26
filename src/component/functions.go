package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"math/rand"
	"strconv"
	"time"
)

type Functions struct {
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
	questionType  int
	seedClicked   bool
	answerClicked bool
	currentValue  string
}

func (fu *Functions) Render() app.UI {
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
						Value(fu.inputValue).
						OnChange(fu.OnInputChange),
					app.Button().
						Text("Generate Seed").
						OnClick(fu.GenerateSeed),
					app.Button().
						Text("up").
						OnClick(fu.upLevel),

					app.Button().
						Text("down").
						OnClick(fu.downLevel),
				),

				app.H6().Text("Seed Value: "+fu.currentValue),
				app.Div().Style("height", "3vh").Body(),
				app.If(fu.seedClicked,
					//question 1
					///////////
					app.If(fu.question1.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q1 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(fu.question1.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q1input").
									Type("text").
									Value(fu.Q1input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(fu.answerClicked,
								app.If(fu.Q1input == fu.question1.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(fu.answerClicked,
							app.P().Text("Answer: "+fu.question1.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(fu.question1.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q1 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(fu.question1.statement1),
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
							app.If(fu.answerClicked,
								app.If(fu.Q1option == fu.question1.answer,
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
						app.If(fu.answerClicked,
							app.P().Text("Answer: "+fu.question1.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					////question 2
					app.If(fu.question2.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(fu.question2.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q2input").
									Type("text").
									Value(fu.Q2input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(fu.answerClicked,
								app.If(fu.Q2input == fu.question2.answer,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(fu.answerClicked,
							app.P().Text("Answer: "+fu.question2.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(fu.question2.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(fu.question2.statement1),
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
							app.If(fu.answerClicked,
								app.If(fu.Q2option == fu.question2.answer,
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
						app.If(fu.answerClicked,
							app.P().Text("Answer: "+fu.question2.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					///////3
					app.If(fu.question3.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(fu.question3.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q3input").
									Type("text").
									Value(fu.Q3input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(fu.answerClicked,
								app.If(fu.Q3input == fu.question3.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(fu.answerClicked,
							app.P().Text("Answer: "+fu.question3.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(fu.question3.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(fu.question3.statement1),
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
							app.If(fu.answerClicked,
								app.If(fu.Q3option == fu.question3.answer,
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
						app.If(fu.answerClicked,
							app.P().Text("Answer: "+fu.question3.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					//////4
					app.If(fu.question4.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q4 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(fu.question4.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q4input").
									Type("text").
									Value(fu.Q4input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(fu.answerClicked,
								app.If(fu.Q4input == fu.question4.answer,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(fu.answerClicked,
							app.P().Text("Answer: "+fu.question4.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(fu.question4.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q4 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(fu.question4.statement1),
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
							app.If(fu.answerClicked,
								app.If(fu.Q4option == fu.question4.answer,
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
						app.If(fu.answerClicked,
							app.P().Text("Answer: "+fu.question4.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					///////5
					app.If(fu.question5.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q5 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(fu.question5.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q5input").
									Type("text").
									Value(fu.Q5input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(fu.answerClicked,
								app.If(fu.Q5input == fu.question5.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(fu.answerClicked,
							app.P().Text("Answer: "+fu.question5.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(fu.question5.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q5 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(fu.question5.statement1),
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
							app.If(fu.answerClicked,
								app.If(fu.Q5option == fu.question5.answer,
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
						app.If(fu.answerClicked,
							app.P().Text("Answer: "+fu.question5.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),

					app.Button().
						Text("Finish").
						Style("margin-left", "50px").
						OnClick(fu.finishClicked),
				),
			),
		),
	)
}

func (fu *Functions) OnInputChange(ctx app.Context, e app.Event) {
	fu.inputValue = ctx.JSSrc().Get("value").String()
	fu.Update()
}
func (fu *Functions) GenerateSeed(ctx app.Context, e app.Event) {
	inputValue := app.Window().GetElementByID("seedInput").Get("value").String()
	fmt.Println("input value:", inputValue)
	seed, err := strconv.ParseInt(inputValue, 10, 64)
	fu.seed = seed
	if err != nil {
		fmt.Println("No input or error converting string to integer:", err)
		seconds := time.Now().UnixNano() / 1000000
		fu.seed = (seconds % 9000) + 1000
		//seed = time.Now().UnixNano()/1000000 - 1713220000000
		//one day 86400 seconds
		fu.currentValue = strconv.FormatInt(fu.seed, 10) + " (No input or error input use random seed)"

	} else {
		fu.currentValue = strconv.FormatInt(fu.seed, 10)
	}

	fu.answerClicked = false
	source := rand.NewSource(fu.seed)
	fu.rand = rand.New(source)
	group1 := []func() Question{
		fu.func1,
		fu.func2,
		fu.func3,
		fu.func4,
		fu.func5,
	}

	fu.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})

	fu.theLevel = fu.seed % 10
	fu.question1 = group1[0]()
	fu.question2 = group1[1]()
	fu.question3 = group1[2]()
	fu.question4 = group1[3]()
	fu.question5 = group1[4]()

	fu.seedClicked = true
	fu.Update()
}

func (fu *Functions) upLevel(ctx app.Context, e app.Event) {
	if fu.seed%10 == 9 {
		fu.seed = fu.seed + 10
	} else {
		fu.seed++
	}

	fu.currentValue = strconv.FormatInt(fu.seed, 10)
	fu.inputValue = strconv.FormatInt(fu.seed, 10)

	fu.answerClicked = false

	source := rand.NewSource(fu.seed)
	fu.rand = rand.New(source)
	group1 := []func() Question{
		fu.func1,
		fu.func2,
		fu.func3,
		fu.func4,
		fu.func5,
	}

	fu.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})

	fu.theLevel = fu.seed % 10
	fu.question1 = group1[0]()
	fu.question2 = group1[1]()
	fu.question3 = group1[2]()
	fu.question4 = group1[3]()
	fu.question5 = group1[4]()

	fu.seedClicked = true
	fu.Update()
}
func (fu *Functions) downLevel(ctx app.Context, e app.Event) {
	if fu.seed%10 == 0 {
		fu.seed = fu.seed - 10
	} else {
		fu.seed--
	}
	fu.currentValue = strconv.FormatInt(fu.seed, 10)
	fu.inputValue = strconv.FormatInt(fu.seed, 10)

	fu.answerClicked = false
	source := rand.NewSource(fu.seed)
	fu.rand = rand.New(source)
	group1 := []func() Question{
		fu.func1,
		fu.func2,
		fu.func3,
		fu.func4,
		fu.func5,
	}

	fu.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})

	fu.theLevel = fu.seed % 10
	fu.question1 = group1[0]()
	fu.question2 = group1[1]()
	fu.question3 = group1[2]()
	fu.question4 = group1[3]()
	fu.question5 = group1[4]()

	fu.seedClicked = true
	fu.Update()
}

func (fu *Functions) finishClicked(ctx app.Context, e app.Event) {
	fu.answerClicked = true
	//q1
	if fu.question1.questionType == 1 {
		fu.Q1input = app.Window().GetElementByID("Q1input").Get("value").String()
	}
	if fu.question1.questionType == 2 {
		if app.Window().GetElementByID("Q1option-a").Get("checked").Bool() {
			fu.Q1option = "true"
		} else if app.Window().GetElementByID("Q1option-b").Get("checked").Bool() {
			fu.Q1option = "false"
		}
	}
	//q2
	if fu.question2.questionType == 1 {
		fu.Q2input = app.Window().GetElementByID("Q2input").Get("value").String()
	}
	if fu.question2.questionType == 2 {
		if app.Window().GetElementByID("Q2option-a").Get("checked").Bool() {
			fu.Q2option = "true"
		} else if app.Window().GetElementByID("Q2option-b").Get("checked").Bool() {
			fu.Q2option = "false"
		}
	}
	//q3
	if fu.question3.questionType == 1 {
		fu.Q3input = app.Window().GetElementByID("Q3input").Get("value").String()
	}
	if fu.question3.questionType == 2 {
		if app.Window().GetElementByID("Q3option-a").Get("checked").Bool() {
			fu.Q3option = "true"
		} else if app.Window().GetElementByID("Q3option-b").Get("checked").Bool() {
			fu.Q3option = "false"
		}
	}
	//q4
	if fu.question4.questionType == 1 {
		fu.Q4input = app.Window().GetElementByID("Q4input").Get("value").String()
	}
	if fu.question4.questionType == 2 {
		if app.Window().GetElementByID("Q4option-a").Get("checked").Bool() {
			fu.Q4option = "true"
		} else if app.Window().GetElementByID("Q4option-b").Get("checked").Bool() {
			fu.Q4option = "false"
		}
	}
	//q5
	if fu.question5.questionType == 1 {
		fu.Q5input = app.Window().GetElementByID("Q5input").Get("value").String()
	}
	if fu.question5.questionType == 2 {
		if app.Window().GetElementByID("Q5option-a").Get("checked").Bool() {
			fu.Q5option = "true"
		} else if app.Window().GetElementByID("Q5option-b").Get("checked").Bool() {
			fu.Q5option = "false"
		}
	}
	fu.Update()
}

type Element interface{}

func (fu *Functions) func1() Question {
	randFunc := fu.rand.Intn(5)
	theFunc := ""
	C := fu.rand.Intn(20) + 1
	D := fu.rand.Intn(20) + 1
	theAnswer := ""
	do := ""
	switch randFunc {
	case 0:
		theFunc = fmt.Sprintf("f(x) = %dx + %d", C, D)
		theAnswer = "bijective"
		do = "R→R"
	case 1:
		theFunc = fmt.Sprintf("f(x) = %dx^2 + %d", C, D)
		theAnswer = "none"
		do = "R→R"
	case 2:
		theFunc = fmt.Sprintf("f(x) = e^x + %d", C)
		theAnswer = "injective"
		do = "R→R"
	case 3:
		theFunc = fmt.Sprintf("f(x) = sin(x)")
		do = "[0,π]→[−1,1]"
		theAnswer = "bijective"
	default:
		theFunc = fmt.Sprintf("f(x) = sin(x)")
		do = "[0,π]→R"
		theAnswer = "injective"
	}

	statement := fmt.Sprintf("Determine whether the function f:%v, defined by %v, is injective, surjective, bijective or none.", do, theFunc)

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}

func (fu *Functions) func2() Question {
	//inverse
	theAnswer := ""
	theFunc := ""
	randC := fu.rand.Intn(6) + 1
	randD := fu.rand.Intn(10) + 1
	randE := fu.rand.Intn(10) + 1
	switch randC {
	case 1:
		theFunc = fmt.Sprintf("(2x-3)/(x+1)")
		theAnswer = "(3x+3)/(2-x)"
	case 2:
		theFunc = fmt.Sprintf("x/(1-x)")
		theAnswer = "x/(1+x)"

	case 3:
		theFunc = fmt.Sprintf("x/(1+x)")
		theAnswer = "x/(1-x)"
	case 4:
		theFunc = fmt.Sprintf("(2x-3)/(x+1)")
		theAnswer = "(3x+3)/(2-x)"
	default:
		theFunc = fmt.Sprintf(" %vx + %v", randD, randE)
		theAnswer = fmt.Sprintf("(x-%v)/%v", randE, randD)
	}

	statement := fmt.Sprintf("Given the function f(x) = %v, find the inverse of the function.", theFunc)

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}

func (fu *Functions) func3() Question {
	//image
	randC := fu.rand.Intn(3)
	theFunc := ""
	theAnswer := ""
	theX := fu.rand.Intn(20) + 1
	randD := fu.rand.Intn(20) + 1
	randE := fu.rand.Intn(20) + 1
	//y = kx + c
	LinearFunc := fmt.Sprintf("%vx + %v", randD, randE)
	LinearImage := randD*theX + randE
	//y = x^2 + b+ c
	QuadraticFunc := fmt.Sprintf("x^2 + %vx + %v", randD, randE)
	QuadraticImage := theX*theX + randD*theX + randE
	//y = x^3
	CubicFunc := fmt.Sprintf("x^3")
	CubicImage := theX * theX * theX

	switch randC {
	case 0:
		theFunc = LinearFunc
		theAnswer = fmt.Sprintf("%v", LinearImage)

	case 1:
		theFunc = QuadraticFunc
		theAnswer = fmt.Sprintf("%v", QuadraticImage)
	default:
		theFunc = CubicFunc
		theAnswer = fmt.Sprintf("%v", CubicImage)

	}

	statement := fmt.Sprintf("Given the function f(x) = %v, find the image of x= %v", theFunc, theX)
	return Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}

}

func (fu *Functions) func4() Question {
	randC := fu.rand.Intn(3)
	theFunc := ""
	theAnswer := ""
	theX := fu.rand.Intn(20) + 1
	randD := fu.rand.Intn(20) + 1
	randE := fu.rand.Intn(20) + 1
	//y = kx + c
	LinearFunc := fmt.Sprintf("%vx + %v", randD, randE)
	LinearImage := randD*theX + randE
	X2 := -randD - theX
	//y = x^2 + b+ c
	QuadraticFunc := fmt.Sprintf("x^2 + %vx + %v", randD, randE)
	QuadraticImage := theX*theX + randD*theX + randE

	//y = x^3
	CubicFunc := fmt.Sprintf("x^3")
	CubicImage := theX * theX * theX
	var image int

	switch randC {
	case 0:
		theFunc = LinearFunc
		image = LinearImage
		theAnswer = fmt.Sprintf("%v", theX)

	case 1:
		theFunc = QuadraticFunc
		image = QuadraticImage
		if X2 == theX {
			theAnswer = fmt.Sprintf("%v", theX)
		} else {
			theAnswer = fmt.Sprintf("%v %v", theX, X2)
		}

		//theAnswer = fmt.Sprintf("%v", QuadraticImage)
	default:
		theFunc = CubicFunc
		image = CubicImage
		theAnswer = fmt.Sprintf("%v", theX)

	}

	statement := fmt.Sprintf("Given the function f(x) = %v, find the Preimage of f(x)= %v", theFunc, image)
	return Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}

}

func (fu *Functions) func5() Question {
	//composition
	fun1 := ""
	fun2 := ""
	theAnswer := ""
	randC := fu.rand.Intn(3)
	randD := fu.rand.Intn(20) + 1
	randE := fu.rand.Intn(20) + 1
	randF := fu.rand.Intn(20) + 1
	randG := fu.rand.Intn(20) + 1

	LinearFunc1 := fmt.Sprintf("%vx + %v", randD, randE)
	LinearFunc2 := fmt.Sprintf("%vx + %v", randF, randG)
	CubicFunc := fmt.Sprintf("x^3")
	switch randC {
	case 0:
		fun1 = LinearFunc1
		fun2 = LinearFunc2
		theAnswer = fmt.Sprintf("%vx + %v", randF*randD, randD*randG+randE)
	case 1:
		fun1 = LinearFunc1
		fun2 = CubicFunc
		theAnswer = fmt.Sprintf("%vx^3 + %v", randD, randE)
	default:
		fun1 = fmt.Sprintf("1/(x+%v)", randD)
		fun2 = fmt.Sprintf("%vx", randE)
		theAnswer = fmt.Sprintf("1/(%vx+%v)", randE, randD)

	}

	statement := fmt.Sprintf("If f(x) = %v and g(x) = %v, find f(g(x))", fun1, fun2)
	return Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
}

func (fu *Functions) randomSet() []int {
	//randomSet
	setLength := fu.rand.Intn(3) + 3
	elements := make(map[int]bool)
	set := make([]int, 0, setLength)

	for len(set) < cap(set) {
		element := fu.rand.Intn(20) + 1
		if !elements[element] {
			set = append(set, element)
			elements[element] = true
		}

	}
	return set
}
