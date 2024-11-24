package main

import(
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
	"time"
	"github.com/gen2brain/beeep"
)

var lastModifiedDate string
var config Config

func main() {
	for {
		log.Println()

		if modifiedFile() {
			loadConfig()
		}

		for _, req := range config.Requests {
			quotation, err := NewQuotation(req.Url)
			if err != nil {
				continue
			}

			for _, rul := range req.Rules {
				alert, err := quotation.Alert(rul)
				if err != nil {
					continue
				}

				if alert {
					msg := fmt.Sprintf("The price of %s is %s %.2f\n", quotation.Symbol, rul.Operator, rul.Value)
					err := beeep.Notify("Crypto Monitoring Alert", msg, "")
					if err != nil {
						fmt.Println("Erro ao enviar a notificação:", err)
						fmt.Printf("This is a alert programmed into Crypto Monitoring to say that the price of %s is %s %.2f\n", quotation.Symbol, rul.Operator, rul.Value)
						continue
					}
				}
			}
		}

		time.Sleep(time.Duration(config.Interval) * time.Second)
	}
}

func modifiedFile() bool {
	fileInfo, err := os.Stat("./config.json");
	if err != nil {
		log.Fatal(err)
	}

	modifiedDateDate := fileInfo.ModTime().Format("20060102150405")

	if lastModifiedDate != modifiedDateDate {
		lastModifiedDate = modifiedDateDate
		return true
	}

	return false
}

func loadConfig() {
	content, _ := ioutil.ReadFile("./config.json")

	err := json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal(err)
	}
}