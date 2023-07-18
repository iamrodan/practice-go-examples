package switchcase

import "fmt"

type OsList struct {
	lists []string
}

func (osList *OsList) includes(osName string) bool {
	for _, os := range osList.lists {
		if os == osName {
			return true
		}
	}
	return false
}

func InterpretUserBasedOnOsUsed() {
	listOfOs := []string{"linux", "mac", "windows"}
	osList := OsList{lists: listOfOs}
	fmt.Println("Which os do you use ?")
	var answer string
	fmt.Scanln(&answer)
	if !osList.includes(answer) {
		fmt.Println("Please provide a valid answer")
		return
	}
	var message string
	switch answer {
	case listOfOs[0]:
		message = "You're a Nerd!"
	case listOfOs[1]:
		message = "You're Rich!"
	case listOfOs[2]:
		message = "You're a Gamer!"
	}
	fmt.Println(message)

}
