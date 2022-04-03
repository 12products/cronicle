package update

import (
	"cronicle/utils"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update a daily file in vim",
		Long:  "update a daily file with number on ordered list as arg",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	files := utils.GetAllFiles("daily")

	n, err := strconv.Atoi(args[0])
	if err != nil || n == 0 || n > len(files) {
		fmt.Printf("Number is not valid")
		return
	}

	path := utils.GetPath([]string{"daily", files[n-1].Name()})

	utils.EditFile(path)
}
