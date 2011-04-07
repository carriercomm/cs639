package loggerTester

import(
	"log"
	"../logger"
	"time"
)

func main () {
	err := logger.Init("test-log.txt")
	if err != nil {
		log.Println("Error in logger.Init: " + err.String())
	}
	id := logger.Start("2 seconds")
	time.Sleep(2000000000)
	errString := logger.End(id)
	if errString != "" {
		log.Println("Error in logger.End: " + errString)
	}
	id = logger.Start("4 seconds")
	time.Sleep(4000000000)
	errString = logger.End(id)
	if errString != "" {
		log.Println("Error in logger.End: " + errString)
	}
}
