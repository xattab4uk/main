pakege main

import (
	"github.com/xattab4uk/main/"
	"log"
)


func main() {
	srv := new(serv.Server)
	if err := srv.Run( port: "8000"); err != nil {
		log.Fatalf (format "error occured whil running http server: %s", err.Error())
	}

}