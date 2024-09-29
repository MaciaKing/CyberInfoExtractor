package models

import (
	"io"
	"net/http"
	"os"
)

// https://docs.virustotal.com/reference/public-vs-premium-api

var (
	MaxRequestPerDay = 500
	RequestPerMinute = 4
	apiKey           = os.Getenv("VIRUS_TOTAL_API_KEY")
	baseUrl          = "https://www.virustotal.com/api/v3/"
)

type VirusTotal struct {
	Id                   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Ip                   string `json:"ip"`
	Domain               string `json:"domain"`
	InformationExtracted string `json:information_extracted`
	Maliciuos            bool   `json:malicious`
}

func DomainReport(domainToRequest string) string {
	endpoint := "domains/" + domainToRequest

	req, _ := http.NewRequest("GET", baseUrl+endpoint, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return string(body)
}
