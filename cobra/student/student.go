package student

import (
	"fmt"
	"os"
)
import "encoding/json"

type Student struct {
	Name    string
	Age     int
	Math    float32
	English float32
}

func (s *Student) ShowInfo() {
	fmt.Printf("name:%s\t", s.Name)
	fmt.Printf("age:%d\t", s.Age)
	fmt.Printf("Math score:%.2f\t", s.Math)
	fmt.Printf("English Score:%.2f\t", s.English)
}

func AddStudent(name string, age int, math float32, english float32) {
	jsonArray, err := json.Marshal(Student{name, age, math, english})
	if err != nil {
		return
	}
	_, err = os.Stdout.Write(jsonArray)
	if err != nil {
		return
	}
	fmt.Println("\nAddStudent called")
}
