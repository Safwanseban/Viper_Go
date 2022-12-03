package main

import (
	"fmt"

	"github.com/Safwanseban/Viper_Go/configs"
)

var Env *configs.EnvCOnfig

func init() {
	Env=configs.Loadenv()
}
func main() {
	configs.Loadenv()
	fmt.Printf("Port:%s\n", Env.Port)

	
}
