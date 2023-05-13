package static

import "fmt"

func HelpMsg() {
	helpText := `Usage: nup [OPTION]

Options:
    -a, --all   Log all requests
    -t          Show latency
    -s          Show status code
    -v          Show verbose output
    -h, --help  Show this help message
		
Examples:
    nup -t -s`

	fmt.Println("\n" + helpText) // "\n" is a newline hack
}
