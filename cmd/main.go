package main

import "github.com/elem1092/crud/pkg/logging"

func main() {
	//entry point
	logger := logging.GetLogger()
	logger.Info("Some log")
}
