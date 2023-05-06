package statics

import "fmt"

func HelpMsg() {
	helpText := `\nUsage: nup [OPTION]

Options:
    -t          Show latency
    -s          Show status code
    -h, --help  Show this help message
		
Examples:
    nup -t -s`

	fmt.Println(helpText)
}
