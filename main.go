package main

import (
	"flag"
	"fmt"
	"log"
	"my-app/config"
	"my-app/school"
	"os"
)

func main() {
	// go run main.go --config config.yaml
	// go run main.go -h
	cFile := flag.String("config", "config.yaml", "Configuration File")
	flag.Parse()
	if *cFile == "" {
		log.Println("missing configuration file")
		os.Exit(1)
	}

	cfg, err := config.ReadConfig(*cFile)
	if err != nil {
		log.Printf("reading config file failed: %v\n", err)
		os.Exit(1)
	}

	stu, err := school.GetStudent("student:1001", cfg.RDBClient)
	if err != nil {
		log.Printf("cannot get student: %v", err)
	}
	fmt.Printf("%+v\n", stu)

}
