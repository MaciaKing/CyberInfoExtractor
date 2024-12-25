package workers

import (
	"CyberInfoExtractor/cmd/globals"
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
	// globals.Rb

	for i := 0; i < totalLinesToRead; i++ {
		read := <-domainsReaded
		rb.PushDataToQueue(globals.DataExtractedQueue, read)
	}

	return nil
}

func ExtractAllQueue() {
	rb := models.Rabbitmq{}
	rb.InitRabbitMQ()

	var readed = make(chan string)
	rb.ReadDataFromQueue(globals.DataExtractedQueue, readed)

	stop := make(chan struct{})

	go func() {
		for element := range readed {
			fmt.Println("Extracted: ", element)
			for _, queue := range globals.WorkersQueue {
				rb.PushDataToQueue(queue, element)
			}
		}
		close(stop)
	}()
	<-stop
}

func VirusTotalWorker() {
	rb := models.Rabbitmq{}
	rb.InitRabbitMQ()

	var readed = make(chan string)
	rb.ReadDataFromQueue(globals.VirusTotalQueue, readed)

	stop := make(chan struct{})

	go func() {
		for element := range readed {
			fmt.Println("[VT] Extracted: ", element)

		}
		close(stop)
	}()
	<-stop
}
