package cfssamplegomodule

import (
	"fmt"
	"time"
)

func SampleGoModuleFunction() string {
	t := time.Now()
	formatted := fmt.Sprintf("%d%02d%02d_%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	return fmt.Sprintf("This is from SampleGoModuleFunction [main] at %s", formatted)
}
