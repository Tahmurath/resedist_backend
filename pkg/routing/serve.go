package routing

import (
	"fmt"
	"log"

	// "net/http"
	"resedist/pkg/config"
	// "github.com/gin-gonic/autotls"
)

func Serve() {
	r := GetRouter()

	configs := config.Get()

	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// log.Fatal(autotls.Run(r, "localhost"))
	//log.Fatal(http.ListenAndServeTLS("localhost:4000", "localhost.crt", "localhost.key", r))

	if err != nil {
		log.Fatal("Error in routing")
		return
	}
}

//openssl req -new -subj "/C=US/ST=Utah/CN=localhost" -newkey rsa:2048 -nodes -keyout localhost.key -out localhost.csr
// openssl x509 -req -days 365 -in localhost.csr -signkey localhost.key -out localhost.crt
//log.Fatal(http.ListenAndServeTLS("localhost:4000", "localhost.crt", "localhost.key", r))
