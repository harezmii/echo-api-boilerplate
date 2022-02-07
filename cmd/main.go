package main

import (
	"api/app/rest"
)

// @title Echo Template Api
// @version 1.0
// @description This is echo server.
// @termsOfService https://localhost/api/v1/

// @contact.name Api Support
// @contact.url https://localhost/support
// @contact.email suatcnby06@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4000
// @BasePath /api/v1
func main() {
	rest.SetupRest("localhost" + ":4000")
}
