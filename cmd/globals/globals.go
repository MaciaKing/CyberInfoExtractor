package globals

var LinesReads = make(chan string, 500)

// Rabbitmq definitions
var DataExtractedQueue = "DataToExtract"
var VirusTotalQueue = "VirusTotal"
var AlienVault = "Alienvault"

var WorkersQueue = [...]string{VirusTotalQueue, AlienVault}

// Globals for tests
var QueueTest1 = "QueueTest1"
