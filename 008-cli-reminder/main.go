package main

import(
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const(
	markName = "GOLANG_CLI_REMINDER"
	markValue = "1"
)

func main(){
	if len(os.Args) > 3{
		fmt.Println("Usage: %s <hh:mm> <text message\n>",os.Args[0])
		os.Exit(1)
	}

	now := time.Now()

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t, err := w.Parse(os.Args[1],now)
	if err != nil{
		fmt.Println(err)
		os.Exit(2)
	}
	if t == nil{
		fmt.Println("Unable to parse time..")
		os.Exit(2)
	}
	if now.After(t.Time){
		fmt.Println("set a future time")
		os.Exit(3)
	}
}