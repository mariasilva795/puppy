package puppy

import "github.com/mariasilva795/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func Woofsss() string {
	return "Woof Woof" + dog.WhenGroup("Woof")
}
