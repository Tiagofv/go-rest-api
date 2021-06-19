package main

import "fmt"
// Struct which contains pointers to database connections.
type App struct {

}

// Run - sets up ou app.
func (app *App) Run() error  {
	fmt.Println("Setting up our app!")
	return nil
}

func main(){
	fmt.Println("Go rest api !")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting server!")
		fmt.Println(err)
	}
}