package handlers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jossep11/config"
	"github.com/jossep11/entities"
)

// Data
func GetDatas(c *fiber.Ctx) error {

	var datas []entities.Data

	config.Database.Find(&datas)

	// return c.Status(200).JSON(dogs)
	// return ;
	// fmt.Print(dogs, dogs)
	return c.Status(200).JSON(datas)
}

func GetData(c *fiber.Ctx) error {
	id := c.Params("id")
	var data entities.Data

	result := config.Database.Find(&data, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	f, err := os.Open("public/upload/" + data.Upload)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	var resultSells float64 = 0
	var resultPurchases float64 = 0
	var i int = 0
	for {
		rec, err := csvReader.Read()
		//   result += float64(rec[5])
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if i > 0 {
			var marksStr string = rec[7]
			marks, err := strconv.ParseFloat(marksStr, 32)
			if err != nil {
				log.Fatal(err)
			}
			// do something with read line
			if rec[13] == "Completed" && rec[2] == "Sell" && rec[3] == "USDT" && rec[4] == "VES" {
				resultSells += (marks)
			}

			if rec[13] == "Completed" && rec[2] == "Buy" && rec[3] == "USDT" && rec[4] == "VES" {
				resultPurchases += (marks)
			}
		}
		i++
	}

	fmt.Println("Total vendido", resultSells)
	fmt.Println("Total comprado", resultPurchases)
	fmt.Println("Ganancias", resultSells-resultPurchases)

	return c.Status(200).JSON(&data)
}

func AddData(c *fiber.Ctx) error {
	dog := new(entities.Data)
	if err := c.BodyParser(dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	file, err := c.FormFile("Upload")
	if err != nil {
		return err
	}
	if _, err := os.Stat("public/upload"); os.IsNotExist(err) {
		// fmt.Print("it doesnt exist")
		if err := os.Mkdir("public/upload", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	c.SaveFile(file, "public/upload/"+file.Filename)
	dog.Upload = file.Filename

	f, err := os.Open("public/upload/" + file.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	var resultSells float64 = 0
	var resultPurchases float64 = 0
	i := 0
	for {
		rec, err := csvReader.Read()
		//   result += float64(rec[5])
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if i > 0 {
			var marksStr string = rec[7]
			marks, err := strconv.ParseFloat(marksStr, 32)
			if err != nil {
				log.Fatal(err)
			}
			// do something with read line
			if rec[13] == "Completed" && rec[2] == "Sell" && rec[3] == "USDT" && rec[4] == "VES" {
				resultSells += (marks)
			}

			if rec[13] == "Completed" && rec[2] == "Buy" && rec[3] == "USDT" && rec[4] == "VES" {
				resultPurchases += (marks)
			}
		}
		i++
	}

	fmt.Println("Total vendido", resultSells)
	fmt.Println("Total comprado", resultPurchases)
	fmt.Println("Ganancias", resultSells-resultPurchases)

	config.Database.Create(&dog)

	return c.Status(201).JSON(dog)
}

func UpdateData(c *fiber.Ctx) error {
	dog := new(entities.Data)
	id := c.Params("id")

	if err := c.BodyParser(dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveData(c *fiber.Ctx) error {
	id := c.Params("id")
	var dog entities.Data

	result := config.Database.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	config.Database.Find(&dog)
	var dogs []entities.Data

	config.Database.Find(&dogs)

	return c.Next()
}

// Users
func GetUsers(c *fiber.Ctx) error {

	var users []entities.Users

	config.Database.Find(&users)

	return c.Status(200).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entities.Users

	result := config.Database.Find(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&user)
}

func AddUser(c *fiber.Ctx) error {
	user := new(entities.Users)
	if err1 := c.BodyParser(user); err1 != nil {
		return c.Status(503).JSON("gohome" + err1.Error())
	}

	// hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return c.Status(500).JSON("Error hashing password")
	// }
	// compare := bcrypt.CompareHashAndPassword(hash, []byte(user.Password))
	// if compare == nil {
	// 	log.Println("success")
	// }

	// user.Password = string(hash)
	config.Database.Create(&user)

	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	user := new(entities.Users)
	id := c.Params("id")

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&user)
	return c.Status(200).JSON(user)
}

func RemoveUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entities.Users

	result := config.Database.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Next()
}
