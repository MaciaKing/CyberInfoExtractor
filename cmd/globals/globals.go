package globals

var LinesReads = make(chan string, 500)

// Rabbitmq definitions
var DataExtractedQueue = "DataToExtract"

// Globals for tests
var QueueTest1 = "QueueTest1"
