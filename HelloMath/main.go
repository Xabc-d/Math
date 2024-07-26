package main

import (
	"HelloMath/src/component"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	// Components routing:

	app.Route("/", &component.Hello{Name: "123"})
	app.Route("/hello", &component.Hello{})
	app.Route("/week01", &component.Terms{})
	app.Route("/week02", &component.Logic{})
	app.Route("/week03", &component.Logic2{})
	app.Route("/week04", &component.Proofs{})
	app.Route("/week05", &component.Proofs2{})
	app.Route("/week06", &component.Induction{})
	app.Route("/week07", &component.Predicates{})
	app.Route("/week08", &component.Set{})
	app.Route("/week09", &component.Functions{})
	app.Route("/week10", &component.Relation{})
	app.Route("/all", &component.All{})

	app.RunWhenOnBrowser()
	_ = app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Resources:   app.GitHubPages("Math"),
		Icon: app.Icon{
			Default:    "/web/logo.png",       // Specify default favicon.
			AppleTouch: "/web/logo-apple.png", // Specify icon on IOS devices.
		},
		Styles: []string{
			"./web/bootstrap/dist/css/bootstrap.css", // Loads hello.css file.
		},
		Scripts: []string{
			"/web/bootstrap/dist/bootstrap.bundle.js", // Loads hello.js file.
		},
	})

	// http.Handle("/", &app.Handler{
	//  Name:        "Hello",
	//  Description: "An Hello World! example",
	//  Icon: app.Icon{
	//   Default:    "/web/logo.png",       // Specify default favicon.
	//   AppleTouch: "/web/logo-apple.png", // Specify icon on IOS devices.
	//  },
	//  Styles: []string{
	//   "./web/bootstrap/dist/css/bootstrap.css", // Loads hello.css file.
	//  },
	//  Scripts: []string{
	//   "/web/bootstrap/dist/bootstrap.bundle.js", // Loads hello.js file.
	//  },
	// })
	//
	// if err := http.ListenAndServe(":8000", nil); err != nil {
	//  log.Fatal(err)
	// }

}
