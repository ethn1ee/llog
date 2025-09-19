package view

import "fmt"

func PrintGet(count int) {
	var msg string
	switch count {
	case 0:
		msg = "no entry retrieved"
	case 1:
		msg = "1 entry retrieved"
	default:
		msg = fmt.Sprintf("%d entries retrieved", count)
	}

	fmt.Println(msg)
}

func PrintAdd(count int) {
	var msg string
	switch count {
	case 0:
		msg = "no entry created"
	case 1:
		msg = "1 entry created"
	default:
		msg = fmt.Sprintf("%d entries created", count)
	}

	fmt.Println(msg)
}

func PrintDelete(count int) {
	var msg string
	switch count {
	case 0:
		msg = "no entry deleted"
	case 1:
		msg = "1 entry deleted"
	default:
		msg = fmt.Sprintf("%d entries deleted", count)
	}

	fmt.Println(msg)
}

func PrintNuke() {
	fmt.Println("database nuked")
}
