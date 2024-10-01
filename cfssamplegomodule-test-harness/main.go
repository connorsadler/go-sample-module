package main

import (
	"fmt"

	cfssamplegomodule "github.com/connorsadler/go-sample-module/cfssamplegomodule"
)

func main() {
	fmt.Println("cfssamplegomodule-test-harness - main running")

	resultFromModule := cfssamplegomodule.SampleGoModuleFunction()
	fmt.Printf("resultFromModule: %s\n", resultFromModule)

	fmt.Println("cfssamplegomodule-test-harness - main done")
}
