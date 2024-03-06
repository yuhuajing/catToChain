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
	app.Post("/chainservice/logischain", logischain)
	app.Post("/chainservice/salechain", salechain)
	app.Post("/chainservice/repairchain", repairchain)
	log.Fatal(app.Listen(":3005"))
}

type ProdInfo struct {
	Number           string `json:"number"`
	Workamount       string `json:"workamount"`
	Persion          string `json:"persion"`
	Workmethod       string `json:"workmethod"`
	Worktime         string `json:"worktime"`
	Remarks          string `json:"remarks"`
	Productbatch     string `json:"productbatch"`
	Componentfactory string `json:"componentfactory"`
	Assemblyteam     string `json:"assemblyteam"`
}
type LogisInfo struct {
	Number            string `json:"number"`
	Workamount        string `json:"workamount"`
	Persion           string `json:"persion"`
	Workmethod        string `json:"workmethod"`
	Worktime          string `json:"worktime"`
	Remarks           string `json:"remarks"`
	Warehouselocation string `json:"warehouselocation"`
}
type SaleInfo struct {
	Number     string `json:"number"`
	Workamount string `json:"workamount"`
	Persion    string `json:"persion"`
	Workmethod string `json:"workmethod"`
	Worktime   string `json:"worktime"`
	Remarks    string `json:"remarks"`
	Saleprice  string `json:"saleprice"`
}
type RepairInfo struct {
	Number              string `json:"number"`
	Persion             string `json:"persion"`
	Workmethod          string `json:"workmethod"`
	Worktime            string `json:"worktime"`
	Remarks             string `json:"remarks"`
	Maintenancelocation string `json:"maintenancelocation"`
}
type ErrorResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

func prodchain(c *fiber.Ctx) error {
	fmt.Println(string(c.Body()))
	payload := &ProdInfo{}
	if err := c.BodyParser(payload); err != nil {
		fmt.Println("Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	description := fmt.Sprintf("workamount: %s, person: %s, workmethod: %s, worktime: %s, remarks: %s, batchnum: %s, assemblyteam: %s, componentfactory: %s", payload.Workamount, payload.Persion, payload.Workmethod, payload.Worktime, payload.Remarks, payload.Productbatch, payload.Assemblyteam, payload.Componentfactory)
	hash := sendtx.AddOrUpdateData(payload.Number, description, 0)
	fmt.Println(hash)
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: hash})
}
func logischain(c *fiber.Ctx) error {
	fmt.Println(string(c.Body()))
	payload := &LogisInfo{}
	if err := c.BodyParser(payload); err != nil {
		fmt.Println("Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	description := fmt.Sprintf("workamount: %s, person: %s, workmethod: %s, worktime: %s, remarks: %s, warehouselocation: %s", payload.Workamount, payload.Persion, payload.Workmethod, payload.Worktime, payload.Remarks, payload.Warehouselocation)
	hash := sendtx.AddOrUpdateData(payload.Number, description, 1)
	fmt.Println(hash)
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: hash})
}
func salechain(c *fiber.Ctx) error {
	fmt.Println(string(c.Body()))
	payload := &SaleInfo{}
	if err := c.BodyParser(payload); err != nil {
		fmt.Println("Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	description := fmt.Sprintf("workamount: %s, person: %s, workmethod: %s, worktime: %s, remarks: %s, saleprice: %s", payload.Workamount, payload.Persion, payload.Workmethod, payload.Worktime, payload.Remarks, payload.Saleprice)
	hash := sendtx.AddOrUpdateData(payload.Number, description, 2)
	fmt.Println(hash)
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: hash})
}
func repairchain(c *fiber.Ctx) error {
	fmt.Println(string(c.Body()))
	payload := &RepairInfo{}
	if err := c.BodyParser(payload); err != nil {
		fmt.Println("Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	description := fmt.Sprintf("person: %s, workmethod: %s, worktime: %s, remarks: %s, maintenancelocation: %s", payload.Persion, payload.Workmethod, payload.Worktime, payload.Remarks, payload.Maintenancelocation)
	hash := sendtx.AddOrUpdateData(payload.Number, description, 3)
	fmt.Println(hash)
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: hash})
}
