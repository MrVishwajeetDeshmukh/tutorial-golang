package main

import "fmt"

func main() {
	var answer, name string

askYesOrNo:
	fmt.Print("Do you have a favorite celebrity or singer? (yes/no): ")
	fmt.Scanf("%s\n", &answer)

	if answer == "no" {
		goto askProgramExit
	}

	fmt.Print("What's their name?: ")
	fmt.Scanf("%s\n", &name)

askProgramExit:
	if name == "" || name == "no" {
		var programExitAnswer string
		fmt.Print("Shall we end the program? (yes/no): ")
		fmt.Scanf("%s\n", &programExitAnswer)

		if programExitAnswer == "yes" {
			goto exit
		}
		goto askYesOrNo
	}

	fmt.Printf("Ah, %s! Thanks for sharing your interest!\n", name)

exit:
}
