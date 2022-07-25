package read

import "fmt"

var (
	firstName, lastName, s string
	i int
	f float32
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)

func Exec() {
	fmt.Println("Please enter you name:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Printf("From the string we read: %d %f %s\n", i, f, s)
}

/**
read.Exec()

output:
Please enter you name:
John Snow
Hi John Snow!
From the string we read: 5212 56.119999 Go
 */