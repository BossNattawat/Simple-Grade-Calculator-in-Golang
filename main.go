package main

import ("fmt")

func grade(score int) string{
	if score >= 80{
		return "A"
	}else if score >= 70{
		return "B"
	}else if score >= 60{
		return "C"
	}else if score >= 50{
		return "D"
	}else{
		return "F"
	}
}

func main(){

	var name string
	var score int

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your score: ")
	fmt.Scan(&score)

	fmt.Println()
	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Score: %v\n", score)
	fmt.Printf("Grade: %v", grade(score))

}