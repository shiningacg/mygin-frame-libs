package log

import (
	"testing"
)

func TestTimePostFix(t *testing.T) {
	OpenLog(&Config{},
		DefaultLogger,
		LoggerOutput(LoggerPostFixTime))
	_Logger.Log("Hello world")
	_Logger.Fatal("Bye")
	_Logger.Close()
}
