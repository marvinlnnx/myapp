package main

import (
	"flag"
	"fmt"
	"log"
	"my-app/config"
	"my-app/school"
	"my-app/shape"
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

	// fmt.Printf("\n<<< Query Student From Redis >>>")
	// stu, err := school.GetStudent("student:1001", cfg.RDBClient)
	// if err != nil {
	// 	log.Printf("cannot get student: %v", err)
	// } else {
	// 	fmt.Printf("%+v\n\n", stu)
	// }

	// fmt.Printf("\n<<< Query Student From Postgres >>>")
	// stu2, err := school.GetStudentP(1, cfg.PDBClient)
	// if err != nil {
	// 	log.Printf("cannot get student (postgres): %v", err)
	// } else {
	// 	fmt.Printf("%+v\n\n", stu2)
	// }

	school := school.NewSchool(cfg.RDBClient, cfg.PDBClient)
	GetStudentFromRedis(school)
	GetStudentFromPostgres(school)

	var s shape.Shape = shape.NewShape(2, 3, "rectangle")
	var s2 shape.Area = shape.NewArea(3)
	var myCalc []shape.IShape = []shape.IShape{s, s2}

	fmt.Printf("Shape Calculated => %v\n\n", myCalc[0].Calculate())
	fmt.Printf("Area Calculated => %v\n\n", myCalc[1].CalculateCircleArea())

}

// GetStudentFromRedis is reteriving student data from redis
func GetStudentFromRedis(school *school.School) {
	student, err := school.GetStudent()
	if err != nil {
		log.Printf("get student from redis failed: %v\n", err)
		return
	}
	fmt.Printf("\n\n Redis => %+v\n\n", student)
}

// GetStudentFromPostgres is reteriving student data from postgres
func GetStudentFromPostgres(school *school.School) {
	student, err := school.GetStudentP()
	if err != nil {
		log.Printf("get student from postgres failed: %v\n", err)
		return
	}
	fmt.Printf("\n\n Postgres => %+v\n\n", student)

}
