package main

// "log"
// "os"

// "github.com/gofiber/fiber/v2"
// "encoding/csv"
// "fmt"
// "log"
// "os"

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	// f, err := excelize.OpenFile("excel.xlsx")
	// open file
	f, err := os.Open("alldata3.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
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

	//  fmt.Printf("%+v\n", resultSells)
	fmt.Println("Total vendido", resultSells)
	fmt.Println("Total comprado", resultPurchases)
	fmt.Println("Ganancias", resultSells-resultPurchases)

	// records := readCsvFile("./alldata.csv")
	// fmt.Println(records)

	// app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	// app.Get("/User", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	// app.Listen(":3000")
}
