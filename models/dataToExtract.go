package models

import (
	"net"
)

type DataToExtract struct {
	Id           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	VirusTotalId int        `gorm:"not null" json:"virus_total_id"`                           // Clave foránea explícita
	VirusTotal   VirusTotal `gorm:"foreignKey:VirusTotalId;references:Id" json:"virus_total"` // Relación
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
