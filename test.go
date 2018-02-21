package main
import (
	"strings"
	"log"
	"github.com/mbetel/go-diceware/diceware"
)

func main() {
	// Generate 6 words using the diceware algorithm.
	list, err := diceware.Generate(3)
	if err != nil  {
		log.Fatal(err)
	}
	log.Printf(strings.Join(list, ""))
}
