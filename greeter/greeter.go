package greeter


// Build with:
// gomobile bind -target ios -o greeter.framework 

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Printer types can print things.
type Printer interface {
	PrintSomething(s string)
}

// Greeter greets people.
type Greeter struct {
	printer Printer
}

// NewGreeter makes a new Greeter.
func NewGreeter(printer Printer) *Greeter {
	return &Greeter{
		printer: printer,
	}
}

func BytesToString(data []byte) string {
	return string(data[:])
}

// Greet greets someone.
func (g *Greeter) Greet(name string) {

	url := "https://google.com/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		fmt.Printf(err.Error())
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		fmt.Printf(err.Error())
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		fmt.Printf(err.Error())
		return
	}

	fmt.Printf("testing")

	fmt.Println(BytesToString(body))

	g.printer.PrintSomething("Hello hi  " + name + "\n")
	//g.printer.PrintSomething()
}
