package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Set struct {
	app.Compo
	rand       *rand.Rand
	inputValue string
	seed       int64

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
	questionType  int
	theLevel      int64
	seedClicked   bool
	answerClicked bool
	currentValue  string
}

func (se *Set) Render() app.UI {
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
						Value(se.inputValue).
						OnChange(se.OnInputChange),
					app.Button().
						Text("Generate Seed").
						OnClick(se.GenerateSeed),
					app.Button().
						Text("up").
						OnClick(se.upLevel),

					app.Button().
						Text("down").
						OnClick(se.downLevel),
				),

				app.H6().Text("Seed Value: "+se.currentValue),
				app.Div().Style("height", "3vh").Body(),
				app.If(se.seedClicked,
					//question 1
					app.If(se.question1.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q1 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(se.question1.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q1input").
									Type("text").
									Value(se.Q1input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(se.answerClicked,
								app.If(se.Q1input == se.question1.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(se.answerClicked,
							app.P().Text("Answer: "+se.question1.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(se.question1.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q1 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(se.question1.statement1),
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
							app.If(se.answerClicked,
								app.If(se.Q1option == se.question1.answer,
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
						app.If(se.answerClicked,
							app.P().Text("Answer: "+se.question1.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					////question 2
					app.If(se.question2.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(se.question2.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q2input").
									Type("text").
									Value(se.Q2input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(se.answerClicked,
								app.If(se.Q2input == se.question2.answer,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(se.answerClicked,
							app.P().Text("Answer: "+se.question2.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(se.question2.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(se.question2.statement1),
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
							app.If(se.answerClicked,
								app.If(se.Q2option == se.question2.answer,
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
						app.If(se.answerClicked,
							app.P().Text("Answer: "+se.question2.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					///////3
					app.If(se.question3.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(se.question3.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q3input").
									Type("text").
									Value(se.Q3input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(se.answerClicked,
								app.If(se.Q3input == se.question3.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(se.answerClicked,
							app.P().Text("Answer: "+se.question3.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(se.question3.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(se.question3.statement1),
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
							app.If(se.answerClicked,
								app.If(se.Q3option == se.question3.answer,
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
						app.If(se.answerClicked,
							app.P().Text("Answer: "+se.question3.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					//////4
					app.If(se.question4.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q4 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(se.question4.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q4input").
									Type("text").
									Value(se.Q4input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(se.answerClicked,
								app.If(se.Q4input == se.question4.answer,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(se.answerClicked,
							app.P().Text("Answer: "+se.question4.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(se.question4.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q4 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(se.question4.statement1),
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
							app.If(se.answerClicked,
								app.If(se.Q4option == se.question4.answer,
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
						app.If(se.answerClicked,
							app.P().Text("Answer: "+se.question4.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					///////5
					app.If(se.question5.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q5 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(se.question5.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q5input").
									Type("text").
									Value(se.Q5input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(se.answerClicked,
								app.If(se.Q5input == se.question5.answer,
									//app.If(se.question1.answer == se.Q1input,
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(se.answerClicked,
							app.P().Text("Answer: "+se.question5.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(se.question5.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q5 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(se.question5.statement1),
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
							app.If(se.answerClicked,
								app.If(se.Q5option == se.question5.answer,
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
						app.If(se.answerClicked,
							app.P().Text("Answer: "+se.question5.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),

					app.Button().
						Text("Finish").
						Style("margin-left", "50px").
						OnClick(se.finishClicked),
				),
			),
		),
	)
}

func (se *Set) OnInputChange(ctx app.Context, e app.Event) {
	se.inputValue = ctx.JSSrc().Get("value").String()
	se.Update()
}
func (se *Set) GenerateSeed(ctx app.Context, e app.Event) {
	inputValue := app.Window().GetElementByID("seedInput").Get("value").String()
	fmt.Println("input value:", inputValue)
	seed, err := strconv.ParseInt(inputValue, 10, 64)
	se.seed = seed
	if err != nil {
		fmt.Println("No input or error converting string to integer:", err)
		seconds := time.Now().UnixNano() / 1000000
		se.seed = (seconds % 9000) + 1000
		//seed = time.Now().UnixNano()/1000000 - 1713220000000
		//one day 86400 seconds
		se.currentValue = strconv.FormatInt(se.seed, 10) + " (No input or error input use random seed)"

	} else {
		se.currentValue = strconv.FormatInt(se.seed, 10)
	}

	se.answerClicked = false
	source := rand.NewSource(se.seed)
	se.rand = rand.New(source)
	group1 := []func() Question{
		se.set1,
		se.set2,
		se.set3,
		se.set4,
		se.set6,
		se.set7,
	}
	se.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})

	se.theLevel = se.seed % 10
	se.question1 = group1[0]()
	se.question2 = group1[1]()
	se.question3 = group1[2]()
	se.question4 = group1[3]()
	se.question5 = group1[4]()

	se.seedClicked = true
	se.Update()
}

func (se *Set) upLevel(ctx app.Context, e app.Event) {
	if se.seed%10 == 9 {
		se.seed = se.seed + 10
	} else {
		se.seed++
	}

	se.currentValue = strconv.FormatInt(se.seed, 10)
	se.inputValue = strconv.FormatInt(se.seed, 10)

	se.answerClicked = false
	source := rand.NewSource(se.seed)
	se.rand = rand.New(source)
	group1 := []func() Question{
		se.set1,
		se.set2,
		se.set3,
		se.set4,
		se.set6,
		se.set7,
	}
	se.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})

	se.theLevel = se.seed % 10
	se.question1 = group1[0]()
	se.question2 = group1[1]()
	se.question3 = group1[2]()
	se.question4 = group1[3]()
	se.question5 = group1[4]()

	se.seedClicked = true
	se.Update()

}
func (se *Set) downLevel(ctx app.Context, e app.Event) {
	if se.seed%10 == 0 {
		se.seed = se.seed - 10
	} else {
		se.seed--
	}
	se.currentValue = strconv.FormatInt(se.seed, 10)
	se.inputValue = strconv.FormatInt(se.seed, 10)

	se.answerClicked = false
	source := rand.NewSource(se.seed)
	se.rand = rand.New(source)
	group1 := []func() Question{
		se.set1,
		se.set2,
		se.set3,
		se.set4,
		se.set6,
		se.set7,
	}
	se.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})

	se.theLevel = se.seed % 10
	se.question1 = group1[0]()
	se.question2 = group1[1]()
	se.question3 = group1[2]()
	se.question4 = group1[3]()
	se.question5 = group1[4]()

	se.seedClicked = true
	se.Update()

}

func (se *Set) finishClicked(ctx app.Context, e app.Event) {
	se.answerClicked = true
	//q1
	if se.question1.questionType == 1 {
		se.Q1input = app.Window().GetElementByID("Q1input").Get("value").String()
	}
	if se.question1.questionType == 2 {
		if app.Window().GetElementByID("Q1option-a").Get("checked").Bool() {
			se.Q1option = "true"
		} else if app.Window().GetElementByID("Q1option-b").Get("checked").Bool() {
			se.Q1option = "false"
		}
	}
	//q2
	if se.question2.questionType == 1 {
		se.Q2input = app.Window().GetElementByID("Q2input").Get("value").String()
	}
	if se.question2.questionType == 2 {
		if app.Window().GetElementByID("Q2option-a").Get("checked").Bool() {
			se.Q2option = "true"
		} else if app.Window().GetElementByID("Q2option-b").Get("checked").Bool() {
			se.Q2option = "false"
		}
	}
	//q3
	if se.question3.questionType == 1 {
		se.Q3input = app.Window().GetElementByID("Q3input").Get("value").String()
	}
	if se.question3.questionType == 2 {
		if app.Window().GetElementByID("Q3option-a").Get("checked").Bool() {
			se.Q3option = "true"
		} else if app.Window().GetElementByID("Q3option-b").Get("checked").Bool() {
			se.Q3option = "false"
		}
	}
	//q4
	if se.question4.questionType == 1 {
		se.Q4input = app.Window().GetElementByID("Q4input").Get("value").String()
	}
	if se.question4.questionType == 2 {
		if app.Window().GetElementByID("Q4option-a").Get("checked").Bool() {
			se.Q4option = "true"
		} else if app.Window().GetElementByID("Q4option-b").Get("checked").Bool() {
			se.Q4option = "false"
		}
	}
	//q5
	if se.question5.questionType == 1 {
		se.Q5input = app.Window().GetElementByID("Q5input").Get("value").String()
	}
	if se.question5.questionType == 2 {
		if app.Window().GetElementByID("Q5option-a").Get("checked").Bool() {
			se.Q5option = "true"
		} else if app.Window().GetElementByID("Q5option-b").Get("checked").Bool() {
			se.Q5option = "false"
		}
	}
	se.Update()
}

// Cartesian product A x B
func (se *Set) set1() Question {
	ALength := se.rand.Intn(2) + 3
	ASlice := make([]int, 0, ALength)
	AExist := make(map[int]bool)
	for len(ASlice) < cap(ASlice) {
		a := se.rand.Intn(20) + 1
		if !AExist[a] {
			ASlice = append(ASlice, a)
			AExist[a] = true
		}
	}
	sort.Slice(ASlice, func(i, j int) bool {
		return ASlice[i] < ASlice[j]
	})
	BLength := se.rand.Intn(2) + 3
	BSlice := make([]int, 0, BLength)
	BExist := make(map[int]bool)
	for len(BSlice) < cap(BSlice) {
		b := se.rand.Intn(20) + 1
		if !BExist[b] {
			BSlice = append(BSlice, b)
			BExist[b] = true
		}
	}
	sort.Slice(BSlice, func(i, j int) bool {
		return BSlice[i] < BSlice[j]
	})
	interfaceSliceA := make([]interface{}, len(ASlice))
	interfaceSliceB := make([]interface{}, len(BSlice))
	for i, v := range ASlice {
		interfaceSliceA[i] = v
	}
	for i, v := range BSlice {
		interfaceSliceB[i] = v
	}
	interfaceSliceX := []interface{}{"x", "y", "z"}
	interfaceSliceD := []interface{}{"x", "y", "z"}
	Sets := [][]interface{}{interfaceSliceA, interfaceSliceB, interfaceSliceX, interfaceSliceD}
	twoSet := se.rand.Perm(len(Sets))
	firstSet := Sets[twoSet[0]]
	secondSet := Sets[twoSet[1]]
	var product [][]interface{}
	for _, a := range firstSet {
		for _, b := range secondSet {
			product = append(product, []interface{}{a, b})
		}
	}
	firstSetOutput := interfaceSingleOutput(firstSet)
	secondSetOutput := interfaceSingleOutput(secondSet)
	//randC := se.rand.Intn(6)
	statement := ""
	theAnswer := ""
	statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of Cartesian product A x B.", firstSetOutput, secondSetOutput)
	theAnswer = interfaceTwoOutput(product)
	//randC := se.rand.Intn(6)
	//switch randC {
	//case 0:
	//	statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of Cartesian product A x B.", firstSetOutput, secondSetOutput)
	//	theAnswer = interfaceTwoOutput(product)
	//case 1:
	//	statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of Cartesian product B x A.", firstSetOutput, secondSetOutput)
	//	theAnswer = interfaceTwoOutput(product)
	//case 2:
	//	statement = fmt.Sprintf("Given the sets A = %v A = %v. List all elements of Cartesian product A x A.", firstSetOutput, firstSetOutput)
	//	theAnswer = interfaceTwoOutput(product)
	//case 3:
	//	statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of Cartesian product A x B x B.", secondSetOutput, secondSetOutput)
	//	theAnswer = interfaceTwoOutput(product)
	//default:
	//	statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of Cartesian product A x B x A.", firstSetOutput, secondSetOutput)
	//	theAnswer = interfaceTwoOutput(product)
	//}

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}

func (se *Set) set2() Question {

	C := se.rand.Intn(15)

	statement := fmt.Sprintf("If set A contains exactly %v distinct elements, How many elements does the power set of A contain?", C)
	theAnswer := strconv.Itoa(1 << uint(C))

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}
func (se *Set) set3() Question {
	ALength := se.rand.Intn(4)
	ADou := se.rand.Intn(3)
	ASlice := make([]int, 0, ALength)
	for len(ASlice) < cap(ASlice) {
		if len(ASlice) == ADou {
			a := se.rand.Intn(20) + 1
			ASlice = append(ASlice, a)
			ASlice = append(ASlice, a)
		} else {
			a := se.rand.Intn(20) + 1
			ASlice = append(ASlice, a)
		}

	}
	ASliceDE := make(map[int]bool)
	for _, value := range ASlice {
		ASliceDE[value] = true
	}
	DE := len(ASliceDE)
	theAnswer := fmt.Sprintf("%v", DE)

	ASliceOutput := setOutput(ASlice)
	statement := fmt.Sprintf("What is the cardinality of this set A = %v?", ASliceOutput)

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}
func (se *Set) set4() Question {
	ALength := se.rand.Intn(2) + 3
	ASlice := make([]int, 0, ALength)
	AExist := make(map[int]bool)
	for len(ASlice) < cap(ASlice) {
		a := se.rand.Intn(30)
		if !AExist[a] {
			ASlice = append(ASlice, a)
			AExist[a] = true
		}
	}
	sort.Slice(ASlice, func(i, j int) bool {
		return ASlice[i] < ASlice[j]
	})
	BLength := se.rand.Intn(2) + 3
	BSlice := make([]int, 0, BLength)
	BExist := make(map[int]bool)
	for len(BSlice) < cap(BSlice) {
		b := se.rand.Intn(30)
		if !BExist[b] {
			BSlice = append(BSlice, b)
			BExist[b] = true
		}
	}
	sort.Slice(BSlice, func(i, j int) bool {
		return BSlice[i] < BSlice[j]
	})

	C := se.rand.Intn(3)
	newSet := make([]int, 0)
	statement := ""
	theAnswer := ""
	if C == 1 {
		//union
		newSetMapDu := make(map[int]bool)
		for _, a := range ASlice {
			newSetMapDu[a] = true
		}
		for _, b := range BSlice {
			newSetMapDu[b] = true
		}
		newSet = make([]int, 0, len(newSetMapDu))
		for key := range newSetMapDu {
			newSet = append(newSet, key)
		}
		sort.Slice(newSet, func(i, j int) bool {
			return newSet[i] < newSet[j]
		})
		theAnswer = setOutput(newSet)
		statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of the union A U B.", setOutput(ASlice), setOutput(BSlice))

	} else if C == 2 {
		//intersection
		for _, a := range ASlice {
			for _, b := range BSlice {
				if a == b {
					newSet = append(newSet, a)
				}
			}
		}
		sort.Slice(newSet, func(i, j int) bool {
			return newSet[i] < newSet[j]
		})
		theAnswer = setOutput(newSet)
		statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of the intersection A ∩ B.", setOutput(ASlice), setOutput(BSlice))
	} else {
		//difference
		newSetMapDu := make(map[int]bool)
		for _, b := range BSlice {
			newSetMapDu[b] = true
		}
		for _, a := range ASlice {
			if !newSetMapDu[a] {
				newSet = append(newSet, a)
			}
		}

		sort.Slice(newSet, func(i, j int) bool {
			return newSet[i] < newSet[j]
		})
		theAnswer = setOutput(newSet)
		statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of the difference A \\ B.", setOutput(ASlice), setOutput(BSlice))
	}

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}

// choice question
func (se *Set) randomSet() []int {
	//randomSet
	setLength := se.rand.Intn(3) + 3
	elements := make(map[int]bool)
	set := make([]int, 0, setLength)

	for len(set) < cap(set) {
		element := se.rand.Intn(20) + 1
		if !elements[element] {
			set = append(set, element)
			elements[element] = true
		}

	}
	return set
}

func (se *Set) set6() Question {
	ALength := se.rand.Intn(2) + 3
	ASlice := make([]int, 0, ALength)
	//AExist := make(map[int]bool)
	a := 0
	for len(ASlice) < cap(ASlice) {
		ASlice = append(ASlice, a)
		a++
		//if !AExist[a] {
		//	ASlice = append(ASlice, a)
		//	a++
		//	AExist[a] = true
		//}
	}
	//sort.Slice(ASlice, func(i, j int) bool {
	//	return ASlice[i] < ASlice[j]
	//})
	BLength := se.rand.Intn(2) + 3
	BSlice := make([]int, 0, BLength)
	BExist := make(map[int]bool)
	for len(BSlice) < cap(BSlice) {
		b := se.rand.Intn(30)
		if !BExist[b] {
			BSlice = append(BSlice, b)
			BExist[b] = true
		}
	}
	sort.Slice(BSlice, func(i, j int) bool {
		return BSlice[i] < BSlice[j]
	})

	C := se.rand.Intn(3)
	newSet := make([]int, 0)
	statement := ""
	theAnswer := ""
	//Arestriction :=
	if C == 1 {
		//union
		newSetMapDu := make(map[int]bool)
		for _, a := range ASlice {
			newSetMapDu[a] = true
		}
		for _, b := range BSlice {
			newSetMapDu[b] = true
		}
		newSet = make([]int, 0, len(newSetMapDu))
		for key := range newSetMapDu {
			newSet = append(newSet, key)
		}
		sort.Slice(newSet, func(i, j int) bool {
			return newSet[i] < newSet[j]
		})
		theAnswer = setOutput(newSet)
		statement = fmt.Sprintf("Given the sets A = {x ∈ N | x<%v} B = %v. List all elements of the union A U B.", ALength, setOutput(BSlice))

	} else if C == 2 {
		//intersection
		for _, a := range ASlice {
			for _, b := range BSlice {
				if a == b {
					newSet = append(newSet, a)
				}
			}
		}
		sort.Slice(newSet, func(i, j int) bool {
			return newSet[i] < newSet[j]
		})
		theAnswer = setOutput(newSet)
		statement = fmt.Sprintf("Given the sets A = {x ∈ N | x<%v} B = %v. List all elements of the intersection A ∩ B.", ALength, setOutput(BSlice))
	} else {
		//difference
		newSetMapDu := make(map[int]bool)
		for _, b := range BSlice {
			newSetMapDu[b] = true
		}
		for _, a := range ASlice {
			if !newSetMapDu[a] {
				newSet = append(newSet, a)
			}
		}

		sort.Slice(newSet, func(i, j int) bool {
			return newSet[i] < newSet[j]
		})
		theAnswer = setOutput(newSet)
		statement = fmt.Sprintf("Given the sets A = {x ∈ N | x<%v} B = %v. List all elements of the difference A \\ B.", ALength, setOutput(BSlice))
	}

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}

func (se *Set) set7() Question {
	setLength := se.rand.Intn(3) + 3
	elements := make(map[int]bool)
	set := make([]int, 0, setLength)
	for len(set) < cap(set) {
		element := se.rand.Intn(20) + 1
		if !elements[element] {
			set = append(set, element)
			elements[element] = true
		}
	}
	statement := fmt.Sprintf("Given the set A = %v. How many elements does the power set of A contain?", setOutput(set))
	theAnswer := fmt.Sprintf("%v", 1<<uint(len(set)))

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}

//choice question
