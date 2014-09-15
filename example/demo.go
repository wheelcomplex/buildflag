//
// buildflag demo
//
// ref: http://golang.org/pkg/go/build/#hdr-Build_Constraints
//

package main

//import "flag"
import "os"
import _ "github.com/wheelcomplex/buildflag"

func main() {
	var cmd string
	for k, v := range os.Args {
		if k == 0 {
			cmd = v
		} else {
			cmd = cmd + " " + v
		}

	}
	println("buildflag demo, command args:", cmd)
}
