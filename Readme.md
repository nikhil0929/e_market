# About
This is a backend server written in Go for an online market place. 

It is has API endpoints for different functionalities:
1. Product:
  - Allows users to get a list of products returned in JSON. 
  - Product data can be used to display products to users on frontend
2. Cart
  - Allows users to add products to their cart which can later be ordered through checkout
  - Cart is saved through a session cookie saved on the users browser
3. Order
  - Users can create an order by checking out the items in their cart. 
  - Use the Stripe API for checkout processing
4. Login 
  - Users can login through username and password which will return a JSON web token (kind of like a cookie) that allows for SSO without re-logging in
  - Done through the jwt-go package
  
All data is stored on a Postgres Database
  
## Next Steps
- I need to do are create an actual frontend in React to interact with the API.
- I need to fix the database table architecture which might be a little messed up. 

### Steps ###

1. go get -u gorm.io/driver/postgres 
2. go get -u gorm.io/gorm
3. go get github.com/joho/godotenv
4. go get -u github.com/gin-gonic/gin
5. go get github.com/mitchellh/mapstructure
6. go get -u github.com/golang-jwt/jwt/v4
7. go get -u github.com/stripe/stripe-go
