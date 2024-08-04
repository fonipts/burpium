package commandaction

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	config "github.com/fonipts/burpium/library/config"
)

func Initialize() {

	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	//invalid_name := ".ignoreburp"
	main_branch := "master"

	dir_backup := filepath.Join(mydir, config.Localrepo)
	__, err_dir_backup := os.Stat(dir_backup)
	os.IsNotExist(err_dir_backup)
	if __ == nil {
		err_dir_backup_mkdir := os.Mkdir(dir_backup, 0750)
		if err_dir_backup_mkdir != nil {
			fmt.Println(err_dir_backup_mkdir)
		}

		dir_root_branch := filepath.Join(dir_backup, config.Branchfolder)
		err_dir_root_branch := os.Mkdir(dir_root_branch, 0750)
		if err_dir_root_branch != nil {
			fmt.Println(err_dir_root_branch)
		}

		dir_root_branch_main := filepath.Join(dir_root_branch, main_branch)
		err_dir_root_branch_main := os.Mkdir(dir_root_branch_main, 0750)
		if err_dir_root_branch_main != nil {
			fmt.Println(err_dir_root_branch_main)
		}

		file_branch_config := filepath.Join(dir_backup, config.BranchConfigFile)
		var branchconfigdetails = config.Branchformat{Name: main_branch, Branches: []config.Subbranchformat{
			{Name: main_branch},
		}}

		branchconfig, err := json.MarshalIndent(branchconfigdetails, "", " ")
		if err != nil {

			fmt.Println(err)
		}

		err = os.WriteFile(file_branch_config, branchconfig, 0755)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("This folder was successfully generated the burpium details")

	} else {
		fmt.Println("The folder` " + config.Localrepo + "` does exist in the system")

	}

}
