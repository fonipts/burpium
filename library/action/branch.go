package commandaction

import (
	"fmt"
	"os"
	"path/filepath"

	branch "github.com/fonipts/burpium/library/action/branchfl"
	config "github.com/fonipts/burpium/library/config"
)

func CmdBranchlist() map[string]branch.BranchExec {
	commands := map[string]branch.BranchExec{
		"status":        &branch.Brnc_status{},
		"create":        &branch.Brnc_create{},
		"backup":        &branch.Brnc_backup{},
		"restore":       &branch.Brnc_restore{},
		"merge":         &branch.Brnc_merge{},
		"delete":        &branch.Brnc_delete{},
		"checkout_upd":  &branch.Brnc_chekcout_upd{},
		"checkout_ingr": &branch.Brnc_chekcout_ingr{},
	}

	return commands
}

func error_details() {
	fmt.Println("See all available command")

	commands := CmdBranchlist()
	for k, n := range commands {
		fmt.Println(k + " ... " + n.Aboutme())
	}

}
func Branching() {

	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	dir_backup := filepath.Join(mydir, config.Localrepo)
	__, err_dir_backup := os.Stat(dir_backup)
	os.IsNotExist(err_dir_backup)

	if __ == nil {
		fmt.Println("The folder` " + config.Localrepo + "` does not exist, please run `burpium init`")

	} else {

		args := os.Args
		if len(args) <= 2 {
			fmt.Println("No command found please check")
			error_details()
			os.Exit(0)
		}
		action_type := os.Args[2]
		commands := CmdBranchlist()

		if command := commands[action_type]; command == nil {
			fmt.Println("No such command found for `burpium branch " + action_type + "`")
			error_details()
			os.Exit(0)
		} else {
			command.Execute()
		}

	}

}
