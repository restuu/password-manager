package app

import (
	"fmt"

	"github.com/spf13/viper"
)

// App represent base app structure
type App struct {
	Port      int
	JWTSecret string
}

// Init App
func Init() *App {
	return &App{}
}

// Start App
func (a *App) Start() {
	fmt.Println("starting app...")
	port := viper.GetInt64("port")
	fmt.Println("PORT:", port)
}
