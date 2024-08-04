package filesystem

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"

	"github.com/fonipts/burpium/library/config"
)

func Getbranchstatus() config.Branchformat {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	dir_backup := filepath.Join(mydir, config.Localrepo, config.BranchConfigFile)

	b, err := os.ReadFile(dir_backup)
	if err != nil {
		fmt.Println(err)
	}
	var branch config.Branchformat
	Data := []byte(string(b))
	var err11 = json.Unmarshal(Data, &branch)
	if err11 != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}

	return branch
}

func Setbranchstatus(name string) {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	dir_backup := filepath.Join(mydir, config.Localrepo, config.BranchConfigFile)

	get_status := Getbranchstatus()
	get_status.Name = name

	var value_not_exist = false
	//	for k :=get_status.Branches {
	//		if k.Name == name {
	//
	//		}
	//	}

	for _, n := range get_status.Branches {
		if n.Name == name {
			value_not_exist = true
		}
	}
	if !value_not_exist {
		get_status.Branches = append(get_status.Branches, config.Subbranchformat{
			Name: name,
		})
	}

	branchconfig, err := json.MarshalIndent(get_status, "", " ")
	if err != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}

	//os.Stdout.Write(branchconfig)

	// You can also write it to a file as a whole.
	err = os.WriteFile(dir_backup, branchconfig, 0755)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("System updated")

}

func UpdateLocalCopy(name string, is_new_folder bool) {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	dir_root_branch_main := filepath.Join(mydir, config.Localrepo, config.Branchfolder, name)
	dir_root_local_repo := filepath.Join(mydir, config.Localrepo)
	var invalid_dir []string
	invalid_dir = append(invalid_dir, mydir)
	invalid_dir = append(invalid_dir, dir_root_local_repo)
	if is_new_folder {
		err_dir_root_branch_main := os.Mkdir(dir_root_branch_main, 0750)
		if err_dir_root_branch_main != nil {
			fmt.Println(err_dir_root_branch_main)
		}
	}
	Glob(mydir, dir_root_branch_main, func(s string) bool {
		r, _ := regexp.Compile("/.burpfold/")
		return !slices.Contains(invalid_dir, s) && !r.MatchString(s) //filepath.Ext(s) == ".txt"
	})
}
func UpdateWorkCopy(name string) {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	dir_root_branch_main := filepath.Join(mydir, config.Localrepo, config.Branchfolder, name)
	//var invalid_dir []string
	//invalid_dir = append(invalid_dir, dir_root_branch_main)
	Glob(dir_root_branch_main, mydir, func(s string) bool {
		//r, _ := regexp.Compile("/.burpfold/")
		return true //!slices.Contains(invalid_dir, s) && !r.MatchString(s) //filepath.Ext(s) == ".txt"
	})
}

func DeleteLocalCopy(name string) {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	get_status := Getbranchstatus()
	for k, n := range get_status.Branches {
		if n.Name == name {
			get_status.Branches = slices.Delete(get_status.Branches, k, k+1)
		}
	}

	branchconfig, err := json.MarshalIndent(get_status, "", " ")
	if err != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}

	//os.Stdout.Write(branchconfig)

	// You can also write it to a file as a whole.
	dir_root_branch_main := filepath.Join(mydir, config.Localrepo, config.Branchfolder, name)
	dir_backup := filepath.Join(mydir, config.Localrepo, config.BranchConfigFile)
	err = os.WriteFile(dir_backup, branchconfig, 0755)
	if err != nil {
		fmt.Println(err)
	}

	os.RemoveAll(dir_root_branch_main)

	//var invalid_dir []string
	//invalid_dir = append(invalid_dir, dir_root_branch_main)
	//Glob(dir_root_branch_main, mydir, func(s string) bool {
	//r, _ := regexp.Compile("/.burpfold/")
	//	return true //!slices.Contains(invalid_dir, s) && !r.MatchString(s) //filepath.Ext(s) == ".txt"
	//})
}

func DeleteLocalRootFile(name string) {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	dir_root_branch_main := filepath.Join(mydir, config.Localrepo, config.Branchfolder, name)
	dir_root_local_repo := filepath.Join(mydir, config.Localrepo)
	var invalid_dir []string
	invalid_dir = append(invalid_dir, mydir)
	invalid_dir = append(invalid_dir, dir_root_local_repo)
	//if is_new_folder {
	//	err_dir_root_branch_main := os.Mkdir(dir_root_branch_main, 0750)
	//	if err_dir_root_branch_main != nil {
	//		fmt.Println(err_dir_root_branch_main)
	//	}
	//}
	GlobDeleteFile(mydir, dir_root_branch_main, func(s string) bool {
		r, _ := regexp.Compile("/.burpfold/")
		return !slices.Contains(invalid_dir, s) && !r.MatchString(s) //filepath.Ext(s) == ".txt"
	})
}
