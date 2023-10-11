package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var db *sqlx.DB

const jwtSecret = "secret secret"

func main() {

	var err error
	db, err = sqlx.Open("mysql", "root:123456@tcp(localhost:3306)/goprogramming?parseTime=true")
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use("/hello", jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(jwtSecret),
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.ErrUnauthorized
		},
	}))

	app.Post("/signup", SignUp)
	app.Post("/login", Login)
	app.Get("/hello", Hello)

	app.Listen("localhost:5000")
}

func SignUp(c *fiber.Ctx) error {
	req := SignUpRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	if req.Username == "" || req.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	query := "insert user (username,password) values (?,?)"
	result, err := db.Exec(query, req.Username, string(password))
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user := User{
		Id:       int(id),
		Username: req.Username,
		Password: string(password),
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx) error {
	req := LoginRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	if req.Username == "" || req.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	user := User{}
	query := "select id, username, password from user where username = ?"
	err = db.Get(&user, query, req.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect username or password")
	}

	expiresAt := time.Now().Add(time.Hour * 2)
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.Id),
		ExpiresAt: &jwt.Time{Time: expiresAt},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"jwtToken": token,
	})
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

type User struct {
	Id       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/*
.
.
.
Just to break.
.
.
.
*/
func Fiber() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	// Middleware all path
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("before")
		err := c.Next()
		fmt.Println("after")
		return err
	})

	app.Use(requestid.New())
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	// Middleware path
	app.Use("/hello", func(c *fiber.Ctx) error {
		fmt.Println("before")
		c.Next()
		fmt.Println("after")
		return nil
	})

	// Get
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Get : Hello, World!")
	})
	// Post
	app.Post("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Post : Hello, World!")
	})

	// Parameters
	app.Get("/hello/:name/:surname?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		surname := c.Params("surname")
		return c.SendString("Hello " + name + " " + surname)
	})

	// ParamsInt
	app.Get("/helloid/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.ErrBadRequest
		}
		return c.SendString(fmt.Sprintf("Get : Hello, %d", id))
	})

	// Query
	app.Get("/query", func(c *fiber.Ctx) error {
		name := c.Query("name")
		surname := c.Query("surname")
		return c.SendString("name : " + name + " surname : " + surname)
	})

	// Query 2
	app.Get("/query2", func(c *fiber.Ctx) error {
		person := Person{}
		c.QueryParser(&person)
		return c.JSON(person)
	})

	// Wildcards
	app.Get("/wildcards/*", func(c *fiber.Ctx) error {
		c.Params("*")
		wildcard := c.Params("*")
		return c.SendString(wildcard)
	})

	// Static File
	app.Static("/", "./wwwroot")

	// NewError
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "content not found")
	})

	// Group
	v1 := app.Group("v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})
	v1.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hell1 v1")
	})

	v2 := app.Group("v2", func(c *fiber.Ctx) error {
		c.Set("Version", "v2")
		return c.Next()
	})
	v2.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hell1 v2")
	})

	// Mounth
	userApp := fiber.New()
	userApp.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login!")
	})

	app.Mount("/user", userApp)

	// Server
	app.Server().MaxConnsPerIP = 1
	app.Get("/server", func(c *fiber.Ctx) error {
		time.Sleep(time.Second * 5)
		return c.SendString("Server!")
	})

	app.Get("/env", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"BaseURL":     c.BaseURL(),
			"Hostname":    c.Hostname(),
			"IP":          c.IP(),
			"IPs":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Subdomains":  c.Subdomains(),
		})
	})

	// Body
	app.Post("/body", func(c *fiber.Ctx) error {
		fmt.Printf("IsJson: %v\n", c.Is("json"))
		fmt.Println(string(c.Body()))

		person := Person{}
		err := c.BodyParser(&person)
		if err != nil {
			return err
		}

		fmt.Println(person)
		return nil
	})

	// Body 2
	app.Post("/body", func(c *fiber.Ctx) error {
		fmt.Printf("IsJson: %v\n", c.Is("json"))
		fmt.Println(string(c.Body()))

		data := map[string]interface{}{}

		err := c.BodyParser(&data)
		if err != nil {
			return err
		}

		fmt.Println(data)
		return nil
	})

	// app.Listen("127.0.0.1:5000")
	app.Listen("localhost:5000")
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
