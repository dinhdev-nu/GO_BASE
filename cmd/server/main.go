package main

import (
	"github.com/dinhdev-nu/GO_BASE/internal/routers"
)

func main() {
  
	r:= routers.NewRouter()
	

  r.Run(":3005") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}






