package main

import (
	"fmt"
	"go-tabungan-api/tabungan"
	"go-tabungan-api/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func main() {

	app := fiber.New()
	
	dsn := "host=postgres user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=myuser dbname=mydb password=mypassword sslmode=disable")
	if err != nil {
		fmt.Println("? Connected error to the Database", err)
	}
	fmt.Println("? Connected Successfully to the Database", db)

	//Migrate Tabel
	db.AutoMigrate(tabungan.Tabungan{}) 
	db.AutoMigrate(tabungan.Mutasi{}) 

	//Repo
	tabunganRepository := tabungan.NewRepository(db)

	//Service
	tabunganService := tabungan.NewService(tabunganRepository)

	//handler
	tabunganHandler := handler.NewTabunganHandler(tabunganService)

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Selamat datang di Tabungan API menggunakan Fiber")
	})
	app.Post("/daftar", tabunganHandler.CreateTabungan)
	app.Post("/tabung", tabunganHandler.TabungMutasi)
	app.Post("/tarik", tabunganHandler.TarikMutasi)
	app.Get("/saldo/:no_rekening", tabunganHandler.GetSaldo)
	app.Get("/mutasi/:no_rekening", tabunganHandler.GetTabungan)

	// Start server
	app.Listen(":8080")
}
