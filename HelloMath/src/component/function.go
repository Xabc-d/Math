package component

//
//import (
//	"fmt"
//	"github.com/maxence-charriere/go-app/v9/pkg/app"
//	"math/rand"
//	"strconv"
//	"time"
//)
//
////1. Given two sets A={1,2} B={x,y}, list all possible functions from set A to set B.
////solution:
////
////f1: 1→x, 2→x
////f2: 1→x, 2→y
////f3: 1→y, 2→x
////f4: 1→y, 2→y
//
//type Element string
//
//type function struct {
//	app.Compo
//	rand          *rand.Rand
//	inputValue    string
//	Q1input       string
//	question1     Question
//	question2     Question
//	question3     Question
//	question4     Question
//	question5     Question
//	question6     Question
//	theLevel      int64
//	seedClicked   bool
//	answerClicked bool
//	currentValue  string
//	a             string
//	b             string
//	c             string
//	Answer1       string
//	Answer2       string
//	Answer3       string
//	theInput      string
//	list          []string
//	solutionQ1    bool
//}
//
//// A := []Element{"a", "b"}
//// B := []Element{"1", "2", "3"}
//// theMapping := setMapping(A, B, []map[Element]Element{}, 0, make(map[Element]Element))
//// printMapping(theMapping)
//func setMapping(A, B []Element, theMapping []map[Element]Element, index int, currentMapping map[Element]Element) []map[Element]Element {
//	if index == len(A) {
//		allMapping := make(map[Element]Element)
//		for k, v := range currentMapping {
//			allMapping[k] = v
//		}
//		theMapping = append(theMapping, allMapping)
//		return theMapping
//	}
//
//	for _, b := range B {
//		currentMapping[A[index]] = b
//		theMapping = setMapping(A, B, theMapping, index+1, currentMapping)
//		delete(currentMapping, A[index])
//	}
//	return theMapping
//}
//
//func (p *function) printMapping(mapping []map[Element]Element) {
//	for i, theMapping := range mapping {
//		print("f", i+1, ":")
//		for a, b := range theMapping {
//			fmt.Printf("%s -> %s ", a, b)
//		}
//		fmt.Println()
//	}
//}
//
//func (p *function) handleMapping(ctx app.Context, e app.Event) {
//	A := []Element{"a", "b"}
//	B := []Element{"1", "2", "3"}
//	theMapping := setMapping(A, B, []map[Element]Element{}, 0, make(map[Element]Element))
//	p.printMapping(theMapping)
//	p.Update()
//}
//
//func injective(f func(int) int, domain []int) bool {
//	theMap := make(map[int]bool)
//	for _, value := range domain {
//		result := f(value)
//		if theMap[result] {
//			return false
//		}
//		theMap[result] = true
//	}
//	return true
//}
//
//func surjective(f func(int) int, domain []int, codomain []int) bool {
//	theMap := make(map[int]bool)
//	for _, value := range codomain {
//		theMap[value] = false
//	}
//	for _, value := range domain {
//		result := f(value)
//		theMap[result] = true
//	}
//	for _, value := range theMap {
//		if !value {
//			return false
//		}
//	}
//	return true
//}
//
//// y=kx+b
//func inverseLinear(k, b, y int) int {
//	return (y - b) / k
//}
//
////
//
//func (p *function) Render() app.UI {
//	return app.Div().Class("container-fluid").Body(
//		app.Div().Body(
//			app.Input().
//				ID("seedInput").
//				Type("text").
//				Value(p.inputValue).
//				OnChange(p.OnInputChange),
//			app.Button().
//				Text("Generate Seed").
//				OnClick(p.GenerateSeed),
//		),
//		app.H5().Text("Seed value: "+p.currentValue),
//
//		app.If(p.seedClicked,
//			app.Div().Class("row").Body(
//				app.Div().Class("col-auto").Body(
//					app.H6().Text("Q1 :"),
//				),
//				app.Div().Class("col").Body(
//					app.P().Text(p.question1.statement1),
//				),
//			),
//			app.Div().Class("row").Body(
//				app.Div().Class("col-auto").Body(
//					app.Input().
//						ID("seedInput").
//						Type("text").
//						Value(p.Q1input).
//						Style("width", "200px").
//						Style("height", "30px").
//						Style("margin-left", "50px").
//						Class("form-control mb-3").
//						Placeholder("").
//						AutoFocus(true),
//				),
//				app.If(p.answerClicked,
//					app.If(p.question1.answer == p.Q1input,
//						app.Div().Class("col").Body(
//							app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
//						),
//					),
//				),
//				app.If(p.answerClicked,
//					app.P().Text("Answer: "+p.question1.answer).Style("font-size", "20px").Style("margin-left", "50px"),
//				),
//			),
//
//			//app.H1().Text("question : "+p.question.statement1),
//
//			//app.Div().Class("row").Body(
//			//	app.Div().Class("col-12").Body(
//			//		app.Div().Class("row").Body(
//			//			app.H5().Class("col-auto").Text("Definition"),
//			//app.Button().
//			//	Class("col-auto btn btn-outline-primary").
//			//	Style("width", "60px").
//			//	Style("height", "20px").
//			//	Style("padding", "0").
//			//	Style("display", "flex").
//			//	Style("justify-content", "center").
//			//	Style("margin-top", "3px").
//			//	Style("font-size", "0.7rem").
//			//	Text("variation"),
//			//),
//			//app.P().Text("Given two sets A={1,2} B={x,y}, list all possible functions from set A to set B"),
//			//app.Input().
//			//	Type("text").
//			//	Style("width", "200px").
//			//	Style("height", "30px").
//			//	Class("form-control mb-3").
//			//	Placeholder("Enter your answer").
//			//	Value(p.Answer1).
//			//	AutoFocus(true).
//			//	OnChange(p.InputChangeQ1),
//			//app.Button().
//			//	Class("btn btn-outline-primary btn-sm").
//			//	Text("Submit Answer").
//			//	OnClick(p.SubmitQ1),
//			//),
//			//),
//			//app.Div().Class("row").Body(
//			//	app.Div().Class("col-12").Style("padding", "0").Style("margin", "0").Body(
//			//		app.Button().Class("btn btn-link").Text("Show Answer").OnClick(p.SolutionQ1),
//			//	),
//			//	p.renderSolutionQ1(),
//			//),
//
//			//app.Div().Class("row").Body(
//			//	app.Div().Class("col-12").Body(
//			//		app.H5().Class("mb-3").Text("Question"),
//			//		app.P().Text("Consider the function 𝑓:{0, 1, … ,9} → {} f(x) = 3x+1, calculate the preimage and image"),
//			//		app.Input().
//			//			Type("text").
//			//			Style("width", "200px").
//			//			Style("height", "30px").
//			//			Class("form-control mb-3").
//			//			Placeholder("Enter your answer").
//			//			Value(p.Answer2).
//			//			AutoFocus(true).
//			//			OnChange(p.InputChangeQ2),
//			//		app.Input().
//			//			Type("text").
//			//			Style("width", "200px").
//			//			Style("height", "30px").
//			//			Class("form-control mb-3").
//			//			Placeholder("Enter your answer").
//			//			Value(p.Answer2).
//			//			AutoFocus(true).
//			//			OnChange(p.InputChangeQ2),
//			//app.Button().
//			//	Class("btn btn-outline-primary btn-sm").
//			//	Text("Submit Answer").
//			//	OnClick(p.SubmitQ2),
//		),
//		//app.Div().Class("row").Body(
//		//	app.Div().Class().Style("padding", "0").Style("margin", "0").Body(
//		//		app.Button().Class("btn btn-link").Text("Show Answer").OnClick(p.SolutionQ2),
//		//	),
//		//),
//		//
//		//
//		//
//
//		//app.Div().Class("row").Body(
//		//	app.Div().Class("col-12").Body(
//		//		app.Div().Class("row").Body(
//		//			app.H5().Class("col-auto").Text("Question"),
//		//			app.Button().
//		//				Class("col-auto btn btn-outline-primary").
//		//				Style("width", "60px").
//		//				Style("height", "20px").
//		//				Style("padding", "0").
//		//				Style("display", "flex").
//		//				Style("justify-content", "center").
//		//				Style("margin-top", "3px").
//		//				Style("font-size", "0.7rem").
//		//				Text("variation"),
//		//		),
//		//		app.P().Text("For the following function, determine whether they are injective, surjective, or both(bijective)"),
//		//		app.Div().Class("form-check").Body(
//		//			app.Input().Type("radio").Class("form-check-input").ID("option1").Name("question").Value("A"),
//		//			app.Label().Class("form-check-label").For("option1").Text("injective only"),
//		//		),
//		//		app.Div().Class("form-check").Body(
//		//			app.Input().Type("radio").Class("form-check-input").ID("option2").Name("question").Value("B"),
//		//			app.Label().Class("form-check-label").For("option2").Text("surjective only"),
//		//		),
//		//		app.Div().Class("form-check").Body(
//		//			app.Input().Type("radio").Class("form-check-input").ID("option3").Name("question").Value("C"),
//		//			app.Label().Class("form-check-label").For("option3").Text("bijective (both)"),
//		//		),
//		//app.Input().
//		//	Type("text").
//		//	Style("width", "200px").
//		//	Style("height", "30px").
//		//	Class("form-control mb-3").
//		//	Placeholder("Enter your answer").
//		//	AutoFocus(true),
//
//		//app.Div().Class("row").Body(
//		//	app.Div().Class("col-12").Style("padding", "0").Style("margin", "0").Body(
//		//		app.Button().Class("btn btn-link").Text("Show Answer"),
//		//	),
//		//),
//
//		//app.Div().Class("row").Body(
//		//	app.Div().Class("col-12").Body(
//		//		app.H5().Class("mb-3").Text("Question"),
//		//		app.P().Text("Given th bijective function f: A→B f(x), determine its inverse function f^-1"),
//		//		app.Input().
//		//			Type("text").
//		//			Style("width", "200px").
//		//			Style("height", "30px").
//		//			Class("form-control mb-3").
//		//			Placeholder("Enter your answer").
//		//			Value(p.Answer2).
//		//			AutoFocus(true).
//		//			OnChange(p.InputChangeQ2),
//		//app.Button().
//		//	Class("btn btn-outline-primary btn-sm").
//		//	Text("Submit Answer").
//		//	OnClick(p.SubmitQ2),
//
//		//),
//		//app.Div().Class("row").Body(
//		//	app.Div().Class().Style("padding", "0").Style("margin", "0").Body(
//		//		app.Button().Class("btn btn-link").Text("Show Answer").OnClick(p.SolutionQ2),
//		//	),
//		//),
//	)
//}
//
//type functionQuestion struct {
//	statement1 string
//	answer     string
//}
//
//func (p *function) q1() functionQuestion {
//
//	return functionQuestion{
//		statement1: "Given two sets A={1,2} B={x,y}, list all possible functions from set A to set B",
//		answer:     "",
//	}
//}
//
//func (p *function) OnInputChange(ctx app.Context, e app.Event) {
//	p.inputValue = ctx.JSSrc().Get("value").String()
//	p.Update()
//}
//
//func (p *function) GenerateSeed(ctx app.Context, e app.Event) {
//	inputValue := app.Window().GetElementByID("seedInput").Get("value").String()
//	fmt.Println("input value:", inputValue)
//	seed, err := strconv.ParseInt(inputValue, 10, 64)
//	if err != nil {
//		fmt.Println("Error converting string to integer:", err)
//		seed = time.Now().UnixNano()
//	}
//	source := rand.NewSource(seed)
//	p.rand = rand.New(source)
//	//p.question = p.createRelation()
//	p.Update()
//}
