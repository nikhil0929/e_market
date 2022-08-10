// define api endpoint routes in here
// Routes should call an appropriate function 'Controller'

// The whole process:
/*
1. Routes called -> Controller function
2. Controller called -> Service function
3. Service does the required things and also calls -> DB operations like fetch,update,delete,insert
*/

//RENAME THIS FILE
package Routes

import (
	"log"

	"github.com/gofiber/fiber"
)

func StartServerProcess() {
	app := fiber.New()

	enableProductRoutes(app)
	enableCollectionRoutes(app)
	// Add the rest of the routes and stuff here ...

	log.Fatal(app.Listen(":8080"))

}
