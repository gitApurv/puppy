package puppy

import "github.com/gitApurv/dog"

var PuppyName string = "Pilla"

func Family(s string) string {
	return dog.DogName + " son's name is " + PuppyName + " and his son's name is " + s
}
