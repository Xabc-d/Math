package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"math/rand"
	"strconv"
	"time"
)

type All struct {
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

func (all *All) Render() app.UI {
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
						Value(all.inputValue).
						OnChange(all.OnInputChange),
					app.Button().
						Text("Generate Seed").
						OnClick(all.GenerateSeed),
					app.Button().
						Text("up").
						OnClick(all.upLevel),

					app.Button().
						Text("down").
						OnClick(all.downLevel),
				),

				app.H6().Text("Seed Value: "+all.currentValue),
				app.Div().Style("height", "3vh").Body(),
				app.If(all.seedClicked,
					//question 1
					///////////
					app.If(all.question1.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q1 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(all.question1.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q1input").
									Type("text").
									Value(all.Q1input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(all.answerClicked,
								app.If(all.Q1input == all.question1.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(all.answerClicked,
							app.P().Text("Answer: "+all.question1.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(all.question1.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q1 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(all.question1.statement1),
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
							app.If(all.answerClicked,
								app.If(all.Q1option == all.question1.answer,
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
						app.If(all.answerClicked,
							app.P().Text("Answer: "+all.question1.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					////question 2
					app.If(all.question2.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(all.question2.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q2input").
									Type("text").
									Value(all.Q2input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(all.answerClicked,
								app.If(all.Q2input == all.question2.answer,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(all.answerClicked,
							app.P().Text("Answer: "+all.question2.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(all.question2.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(all.question2.statement1),
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
							app.If(all.answerClicked,
								app.If(all.Q2option == all.question2.answer,
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
						app.If(all.answerClicked,
							app.P().Text("Answer: "+all.question2.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					///////3
					app.If(all.question3.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(all.question3.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q3input").
									Type("text").
									Value(all.Q3input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(all.answerClicked,
								app.If(all.Q3input == all.question3.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(all.answerClicked,
							app.P().Text("Answer: "+all.question3.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(all.question3.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(all.question3.statement1),
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
							app.If(all.answerClicked,
								app.If(all.Q3option == all.question3.answer,
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
						app.If(all.answerClicked,
							app.P().Text("Answer: "+all.question3.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					//////4
					app.If(all.question4.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q4 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(all.question4.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q4input").
									Type("text").
									Value(all.Q4input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(all.answerClicked,
								app.If(all.Q4input == all.question4.answer,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(all.answerClicked,
							app.P().Text("Answer: "+all.question4.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(all.question4.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q4 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(all.question4.statement1),
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
							app.If(all.answerClicked,
								app.If(all.Q4option == all.question4.answer,
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
						app.If(all.answerClicked,
							app.P().Text("Answer: "+all.question4.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					///////5
					app.If(all.question5.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q5 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(all.question5.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q5input").
									Type("text").
									Value(all.Q5input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(all.answerClicked,
								app.If(all.Q5input == all.question5.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(all.answerClicked,
							app.P().Text("Answer: "+all.question5.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(all.question5.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q5 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(all.question5.statement1),
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
							app.If(all.answerClicked,
								app.If(all.Q5option == all.question5.answer,
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
						app.If(all.answerClicked,
							app.P().Text("Answer: "+all.question5.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),

					app.Button().
						Text("Finish").
						Style("margin-left", "50px").
						OnClick(all.finishClicked),
				),
			),
		),
	)
}

func (all *All) OnInputChange(ctx app.Context, e app.Event) {
	all.inputValue = ctx.JSSrc().Get("value").String()
	all.Update()
}

func (all *All) GenerateSeed(ctx app.Context, e app.Event) {
	inputValue := app.Window().GetElementByID("seedInput").Get("value").String()
	fmt.Println("input value:", inputValue)
	seed, err := strconv.ParseInt(inputValue, 10, 64)
	all.seed = seed
	if err != nil {
		fmt.Println("No input or error converting string to integer:", err)
		seconds := time.Now().UnixNano() / 1000000
		all.seed = (seconds % 9000) + 1000
		//seed = time.Now().UnixNano()/1000000 - 1713220000000
		//one day 86400 seconds
		all.currentValue = strconv.FormatInt(all.seed, 10) + " (No input or error input use random seed)"

	} else {
		all.currentValue = strconv.FormatInt(all.seed, 10)
	}
	all.answerClicked = false
	source := rand.NewSource(all.seed)
	all.rand = rand.New(source)
	terms := &Terms{}
	logic := &Logic{}
	function := &Functions{}
	relation := &Relation{}
	set := &Set{}
	all.question2 = terms.Te1()
	all.question1 = all.al()
	group1 := []func() Question{
		terms.Te1,
		terms.te3,
		terms.te4,
		terms.te8,
		set.set1,
		set.set2,
		set.set4,
		set.set6,
		set.set7,
		relation.createRelation,
		relation.inverse,
		relation.inverseAlp,
		relation.inverseFix,
		relation.reflexiveSingle,
		relation.antisymmetricSingle,
		relation.transitiveSingle,
		relation.symmetricSingle,
		function.func1,
		function.func2,
		function.func3,
		function.func4,
		logic.lo1,
		logic.lo4,
		logic.lo2,
	}
	group2 := []func() Question{
		terms.te7,
		terms.te6,
		set.set3,
		relation.composition,
		relation.partialOrder,
		logic.lo3,
		logic.lo5,
		function.func5,

		terms.te5,
	}
	all.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})
	all.rand.Shuffle(len(group2), func(i, j int) {
		group2[i], group2[j] = group2[j], group2[i]
	})
	all.theLevel = all.seed % 10

	//if all.theLevel >= 0 && all.theLevel < 5 {
	//
	//} else {
	//
	//}

	all.seedClicked = true
	all.Update()
}
func (all *All) upLevel(ctx app.Context, e app.Event) {

}
func (all *All) downLevel(ctx app.Context, e app.Event) {

}
func (all *All) finishClicked(ctx app.Context, e app.Event) {
	all.answerClicked = true
	//q1
	if all.question1.questionType == 1 {
		all.Q1input = app.Window().GetElementByID("Q1input").Get("value").String()
	}
	if all.question1.questionType == 2 {
		if app.Window().GetElementByID("Q1option-a").Get("checked").Bool() {
			all.Q1option = "true"
		} else if app.Window().GetElementByID("Q1option-b").Get("checked").Bool() {
			all.Q1option = "false"
		}
	}
	//q2
	if all.question2.questionType == 1 {
		all.Q2input = app.Window().GetElementByID("Q2input").Get("value").String()
	}
	if all.question2.questionType == 2 {
		if app.Window().GetElementByID("Q2option-a").Get("checked").Bool() {
			all.Q2option = "true"
		} else if app.Window().GetElementByID("Q2option-b").Get("checked").Bool() {
			all.Q2option = "false"
		}
	}
	//q3
	if all.question3.questionType == 1 {
		all.Q3input = app.Window().GetElementByID("Q3input").Get("value").String()
	}
	if all.question3.questionType == 2 {
		if app.Window().GetElementByID("Q3option-a").Get("checked").Bool() {
			all.Q3option = "true"
		} else if app.Window().GetElementByID("Q3option-b").Get("checked").Bool() {
			all.Q3option = "false"
		}
	}
	//q4
	if all.question4.questionType == 1 {
		all.Q4input = app.Window().GetElementByID("Q4input").Get("value").String()
	}
	if all.question4.questionType == 2 {
		if app.Window().GetElementByID("Q4option-a").Get("checked").Bool() {
			all.Q4option = "true"
		} else if app.Window().GetElementByID("Q4option-b").Get("checked").Bool() {
			all.Q4option = "false"
		}
	}
	//q5
	if all.question5.questionType == 1 {
		all.Q5input = app.Window().GetElementByID("Q5input").Get("value").String()
	}
	if all.question5.questionType == 2 {
		if app.Window().GetElementByID("Q5option-a").Get("checked").Bool() {
			all.Q5option = "true"
		} else if app.Window().GetElementByID("Q5option-b").Get("checked").Bool() {
			all.Q5option = "false"
		}
	}
	all.Update()
}

func (all *All) al() Question {
	a := all.rand.Intn(30) + 1
	b := all.rand.Intn(30) + 1
	c := all.rand.Intn(30) + 1
	d := all.rand.Intn(30) + 1

	statement := fmt.Sprintf("Given the syntax rule for the Associative Law of addition: (a+b)+c ⊢ a+(b+c), perform a derivation to simplify the expression ((%v+%v)+%v)+%v?", a, b, c, d)
	theAnswer := fmt.Sprintf("%v+(%v+(%v+%v))", a, b, c, d)
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}
