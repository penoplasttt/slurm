package main

import (
	"fmt"
)

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

func main() {
	
	fmt.Println("Тут будет выведен идентификатор")

	result := GenerateCheck()

	for _, value := range result{
		if value.status == "pass"{
		fmt.Print(value.ServiceID)
		}
	}
}

type HealthCheck struct{
	ServiceID int
	status string
}

func GenerateCheck()( []HealthCheck){
	result := []HealthCheck{}
	st := ""
	for i:= 0; i<=5; i++ {
		if i%2 == 0{
			st = PassStatus
		} else {
			st = FailStatus
		}
		p := HealthCheck{i,st}
		result = append(result, p)
		
	}
	return result
}
