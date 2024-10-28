package app

import "log"

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run() {
}

func (app *App) Logo() {
	banner := `                                                                                                                                                 
     .oooooo.              .o8          o8o       .   
    d8P'  'Y8b             "888         '"'     .o8   
   888      888  oooo d8b   888oooo.   oooo   .o888oo 
   888      888  '888""8P   d88' '88b  '888     888   
   888      888   888       888   888   888     888   
   '88b    d88'   888       888   888   888     888 . 
    'Y8bood8P'   d888b      'Y8bod8P'  o888o    "888" 																																									
	`
	log.Println(banner)
}
