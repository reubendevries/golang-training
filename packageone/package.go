package packageone

import "fmt"

// public variable called PackageVar
var PackageVar = "This is a public package level variable."

// public function called PrintMe which takes the following parameters:
// myVar, blockVar, PackageVar

func PrintMe(myVar, blockVar, PackageVar string) {

	fmt.Println(myVar, blockVar, PackageVar)

}
