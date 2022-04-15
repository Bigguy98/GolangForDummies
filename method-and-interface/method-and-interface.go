package main

import (
	"fmt"
)

// declaer interface
type Robot interface {
	PowerOn() error
	/** if you uncomment this line, there will be an error
	"cannot use &t850 (value of type *T850) as Robot value in argument to Boot: *T850 does not implement Robot (missing method PowerOff)"
	**/
	// PowerOff() error
}

type T850 struct {
	Name string
}

type T750 struct {
	Model string
}

// this is how youi implement PowerOn() method of Robot interface
func (t *T850) PowerOn() error {
	fmt.Println("Robot type: ", t.Name)
	return nil
}

// this is how youi implement PowerOn() method of Robot interface
func (t *T750) PowerOn() error {
	fmt.Println("Robot type: ", t.Model)
	return nil
}

// cause T850, T750 implement interface Robot, they are considered as a Robot
func Boot(r Robot) error {
	return r.PowerOn()
}

func main() {

	fmt.Println("*************** METHOD AND INTERFACE EXAMPLE *********************")
	t850 := T850{"God"}
	// that why we can pass T850 here
	if Boot(&t850) == nil {
		fmt.Println("Restart robot T850")
	}

	t750 := T750{"God Model"}
	if Boot(&t750) == nil {
		fmt.Println("Restart robot T750")
	}

}
