package main

import "fmt"

func main() {
	var (
		name string 
		courseNumber int64
		total int64 = 0
	)

	fmt.Println("Please insert your name:")
	fmt.Scanln(&name)

	fmt.Println("Please insert the number of courses you take:")
	fmt.Scanln(&courseNumber)

	fmt.Println(courseNumber, name)
	courses := make(map[string]int64)


	for i := int64(0); i < courseNumber; i++{
		var 
		(
			subject string
			score int64
		)
		fmt.Println("Please insert the subject")
		fmt.Scanln(&subject)

		fmt.Printf("Please insert your %s score\n", subject)
		fmt.Scanln(&score)

		courses[subject] = score
		total += score
	}

	fmt.Println( " --------------------------- ---------------- ---------------------")
	fmt.Println("Name :", name)
	fmt.Println("Courses and scores:")
	for subject, score := range courses {
			fmt.Printf("%s: %d\n", subject, score)
	}
	fmt.Println(" ___________________________________________________________________")
	fmt.Println("Average :", total/courseNumber)
}