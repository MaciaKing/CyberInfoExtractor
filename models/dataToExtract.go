package models

import (
	"fmt"
	"net"

	"gorm.io/gorm"
)

type DataToExtract struct {
	Id           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	VirusTotalId int        `gorm:"not null" json:"virus_total_id"` // Clave foránea explícita
	Malicious    bool       `json:malicious`
	VirusTotal   VirusTotal `gorm:"foreignKey:VirusTotalId;references:Id" json:"virus_total"` // Relación
	FromFile     string     `json:from_file`
	LineOfFile   int        `json:line_of_file`
}

func (dt *DataToExtract) ExtractData(db *gorm.DB, domain string, malicious bool) error {
	vt := VirusTotal{}
	vt.DomainReport(domain)

	if err := db.Create(&vt).Error; err != nil {
		return fmt.Errorf("********** failed to save VirusTotal record: %w", err)
	}

	dt.VirusTotalId = vt.Id
	dt.Malicious = malicious
	if err := db.Create(&dt).Error; err != nil {
		return fmt.Errorf("********** failed to save VirusTotal record: %w", err)
	}
	return nil
}

func (dt *DataToExtract) GetAllVirusTotal(db *gorm.DB) []VirusTotal {
	var virusTotal []VirusTotal
	db.Raw("SELECT * FROM data_to_extracts, virus_totals WHERE data_to_extracts.id= ? AND data_to_extracts.virus_total_id = virus_totals.id", dt.Id).Scan(&virusTotal)
	return virusTotal
}

// Return 0 if teDetect is a domain or URL.
// Retrun 1 if toDetect is an IP.
func DetectIp(toDetect string) int {
	ip := net.ParseIP(toDetect)
	if ip == nil {
		// Is not an IP
		return 0
	}
	// Is a IP
	return 1
}
