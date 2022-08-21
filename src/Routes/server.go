// define api endpoint routes in here
// Routes should call an appropriate function 'Controller'

// The whole process:
/*
1. Routes called -> Controller function
2. Controller called -> Service function
3. Service does the required things and also calls -> DB operations like fetch,update,delete,insert
*/

// This file should NOT be in the Routes folder. Pls move it to the src folder or create a new folder and move it there.
package Routes

import (
	"github.com/gin-gonic/gin"
)

func StartServerProcess() {
	router := gin.Default()

	enableProductRoutes(router)
	//enableUserRoutes(app)
	// enableCartRoutes(app)
	// enableProductRoutes(app)
	// enableOrderRoutes(app)
	// Add the rest of the routes and stuff here ...

	router.Run(":8080")

}
