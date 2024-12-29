package workers

import (
	"CyberInfoExtractor/cmd/globals"
	"CyberInfoExtractor/database"
	"CyberInfoExtractor/models"
	"fmt"
)

func ReadDataToExtract(fileToRead string) error {
	fr := models.FileReader{}
	totalLinesToRead, err := fr.TotalLineCounter(fileToRead)
	domainsReaded := make(chan string, totalLinesToRead)
	linesRead, err := fr.ReadFileFromChan(fileToRead, 0, totalLinesToRead, domainsReaded)
	if err != nil && linesRead != 0 {
		return err
	}

	rb := models.Rabbitmq{}
	rb.InitRabbitMQ()

	for i := 0; i < totalLinesToRead; i++ {
		read := <-domainsReaded
		rb.PushDataToQueue(globals.DataExtractedQueue, read)
	}
	return nil
}

func ExtractAllQueue() {
	var readed = make(chan string)
	rb := models.Rabbitmq{}
	rb.InitRabbitMQ()
	rb.ReadDataFromQueue(globals.DataExtractedQueue, readed)
	stop := make(chan struct{})
	go func() {
		for element := range readed {
			fmt.Println("Extracted: ", element)
			dt := models.DataToExtract{}
			for _, queue := range globals.WorkersQueue {
				if models.DetectIp(element) == 0 {
					dt.Domain = element
				} else {
					dt.Ip = element
				}
				rb.PushDataToQueue(queue, element)
			}
			dt.VirusTotalId = createDefaultStructs()
			database.DB.Create(&dt)
		}
		close(stop)
	}()
	<-stop
	rb.CloseRabbitMQ()
}

func VirusTotalWorker() {
	// rb := models.Rabbitmq{}
	// rb.InitRabbitMQ()

	// var readed = make(chan string)
	// rb.ReadDataFromQueue(globals.VirusTotalQueue, readed)

	// stop := make(chan struct{})

	// go func() {
	// 	for element := range readed {
	// 		fmt.Println("[VT] Extracted: ", element)

	// 	}
	// 	close(stop)
	// }()
	// <-stop
}

func createDefaultStructs() int {
	vt := models.VirusTotal{}
	database.DB.Create(&vt)

	return vt.Id
}
