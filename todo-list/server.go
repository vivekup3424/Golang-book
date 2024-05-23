package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type task struct {
	Item string
}

func main() {
	//before creating an instance of the fiber app,
	//add this code
	//connStr := "postgresql://<username>:<password>@<database_ip>/<database-name>?sslmode=disable"
	connStr := "postgresq://postgres:postgres@127.0.0.1:5432/testdb?sslmode=disable"

	//Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	/** we will add the code to create a database connection and also update our routes to pass the database connection to our handlers so we can use it to execute database queries:**/

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})
	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})
	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})
	port := os.Getenv("PORT")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))

	if port == "" {
		port = "3000"
	}
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var tasks []string
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&res)
		tasks = append(tasks, res)
	}
	return c.Render("index", fiber.Map{
		"Todos": tasks,
	})
}
func postHandler(c *fiber.Ctx, db *sql.DB) error {
	newTask := task{}
	if err := c.BodyParser(&newTask); err != nil {
		log.Printf("An error occured in post request: %v",
			err)
	}
	fmt.Printf("Task = %v", newTask)

	//check if non-empty request is being send
	if newTask.Item != "" {
		_, err := db.Exec(`
			INSERT INTO tasks(id,name)
			VALUES ($1)
			`, newTask.Item)

		if err != nil {
			log.Fatalf(`An error occured
						while executing queries: %v`, err)
		}
	}
	return c.Redirect("/")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	olditem := c.Query("olditem")
	newitem := c.Query("newitem")
	db.Exec("UPDATE tasks set name = $1, WHERE item = $2",
		newitem, olditem)
	return c.Redirect("/")
}
func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	taskDelete := c.Query("name")
	db.Exec("DELETE from tasks WHERE item=$1", taskDelete)
	c.SendString("deleted")
	c.Redirect("/")
}
