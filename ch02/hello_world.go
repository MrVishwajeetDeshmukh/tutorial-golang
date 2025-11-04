package main

// When importing multiple packages, you can group them using parentheses.
// Note: Do not separate them with commas — be careful to avoid typos.
import (
	"fmt"
	// For packages we’ve downloaded, use their exact import path.
	"github.com/fatih/color"
)

func main() {
	// The Printf function prints formatted text.
	// It displays the following arguments according to the format string specified in the first argument.
	// %v is a format specifier that prints the value according to its type.
	fmt.Printf("%v %v\n",
		// The color package belongs to the imported fatih/color module.
		// The following code adds color to the output text.
		color.RedString("Hello"),
		color.GreenString("World!"))
}
