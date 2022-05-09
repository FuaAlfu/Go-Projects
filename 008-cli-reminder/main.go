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
	}
}