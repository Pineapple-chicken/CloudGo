package main

import flag "github.com/spf13/pflag"
import "github.com/pineapple-chicken/cloudgo/service"
import "os"

const (
	PORT string = "8000"
)

func main() {
	var port string
	//Get predefined env port
	port = os.Getenv("PORT")
	if (len(port) <= 0) {
		port = PORT
	}
	//Define server port
	flag.StringVarP(&port, "port", "p", "8000", "define server port")
	flag.Parse()
	server := service.NewServer()
	server.Run(":" + port)
}
