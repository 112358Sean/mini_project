package routes

import (
	"mini_project/configs"
	c "mini_project/controllers"
	m "mini_project/middlewares"
	r "mini_project/repositories"
	s "mini_project/services"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	DB = configs.InitDB()

	JWT = m.NewJWTS()

	userR = r.NewUserRepository(DB)
	userS = s.NewUserService(userR)
	userC = c.NewUserController(userS, JWT)

	bookR = r.NewBookRepository(DB)
	bookS = s.NewBookService(bookR)
	bookC = c.NewBookController(bookS)

	cartR = r.NewCartRepository(DB)
	cartS = s.NewCartService(cartR)
	cartC = c.NewCartController(cartS)

	transactionR = r.NewTransactionRepository(DB)
	transactionS = s.NewTransactionService(transactionR)
	transactionC = c.NewTransactionController(transactionS)
)

func New() *echo.Echo {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()

	m.LoggerMiddleware(e)

	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))
	auth.GET("/users", userC.GetUsersController)
	auth.GET("/users/:id", userC.GetUserController)
	e.POST("/users", userC.CreateController)
	auth.DELETE("/users/:id", userC.DeleteController)
	auth.PUT("/users/:id", userC.UpdateController)

	e.GET("/books", bookC.GetBooksController)
	e.GET("/books/:id", bookC.GetBookController)
	auth.POST("/books", bookC.CreateController)
	auth.DELETE("/books/:id", bookC.DeleteController)
	auth.PUT("/books/:id", bookC.UpdateController)

	auth.GET("/carts", cartC.GetCartsController)
	auth.GET("/carts/:id", cartC.GetCartController)
	auth.POST("/carts", cartC.CreateController)
	auth.DELETE("/carts/:id", cartC.DeleteController)
	auth.PUT("/carts/:id", cartC.UpdateController)

	auth.GET("/transactions", transactionC.GetTransactionsController)
	auth.GET("/transactions/:id", transactionC.GetTransactionController)
	auth.POST("/transactions", transactionC.CreateController)
	auth.DELETE("/transactions/:id", transactionC.DeleteController)
	auth.PUT("/transactions/:id", transactionC.UpdateController)

	e.POST("/login", userC.LoginController)

	return e
}
