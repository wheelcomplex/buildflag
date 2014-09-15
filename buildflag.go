//
// package buildfalg use for go build flag and comfirm
//
// ref:
// http://golang.org/cmd/ld/
/*
-X symbol value
	Set the value of an otherwise uninitialized string variable.
	The symbol name should be of the form importpath.name,
	as displayed in the symbol table printed by "go tool nm".
*/
// http://golang.org/pkg/go/build/#hdr-Build_Constraints
// http://stackoverflow.com/questions/15214459/how-to-properly-use-build-tags
// http://stackoverflow.com/questions/11354518/golang-application-auto-build-versioning
//
// USAGE:
// import this package by import _ "github.com/wheelcomplex/buildflag"
// use command: go build -ldflags '-X github.com/wheelcomplex/buildflag.BINFLAG MYFLAG2501' your.go to make binary
// and execute your program with -showbinflag to show the flag "MYFLAG2501"
//

package buildflag

import (
	"flag"
	"os"
)

var BINFLAG string
var showFlag *bool

func init() {
	showFlag = flag.Bool("showbinflag", false, "show binary build flag")
	flag.Parse()
	if *showFlag {
		println("Binary Build Flag", BINFLAG)
		os.Exit(0)
	}
}

//
func UNUSED() {

}
