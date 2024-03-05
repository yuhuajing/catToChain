package explorer

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	"log"
	"main/core/sendtx"
)

func Explorer() {
	app := fiber.New()
	app.Post("/chainservice/prodchain", prodchain)
	app.Post("/chainservice/logischain", uploadchain)
	app.Post("/chainservice/salechain", uploadchain)
	app.Post("/chainservice/repairchain", uploadchain)
	log.Fatal(app.Listen(":3005"))
}

type CarInfo struct {
	Number     string `json:"number"`
	Workamount string `json:"workamount"`
	Persion    string `json:"persion"`
	Workmethod string `json:"workmethod"`
	Worktime   string `json:"worktime"`
	Remarks    string `json:"remarks"`
}
type ErrorResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

func prodchain(c *fiber.Ctx) error {
	fmt.Println(string(c.Body()))
	payload := &CarInfo{}
	if err := c.BodyParser(payload); err != nil {
		fmt.Println("Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	hash := sendtx.AddOrUpdateProdData(payload.Number, payload.Workamount, payload.Persion, payload.Workmethod, payload.Worktime, payload.Remarks)
	fmt.Println(hash)
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: hash})
}
