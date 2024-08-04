package branch

import (
	"fmt"
	"os"

	filesystem "github.com/fonipts/burpium/library/core/filesystem"
)

type BranchExec interface {
	Execute()
	Aboutme() string
}

type Brnc_status struct {
}

func (c *Brnc_status) Aboutme() string {
	return "Initialize the project"
}
func (c *Brnc_status) Execute() {

	branch := filesystem.Getbranchstatus()
	fmt.Println("Current branch:" + branch.Name)
}

type Brnc_chekcout_upd struct {
}

func (c *Brnc_chekcout_upd) Aboutme() string {
	return "Initialize the project"
}
func (c *Brnc_chekcout_upd) Execute() {
	require_branch_name()
	action_type := os.Args[3]
	branch := filesystem.Getbranchstatus()
	filesystem.Setbranchstatus(action_type)
	filesystem.UpdateLocalCopy(branch.Name, false)

	filesystem.DeleteLocalRootFile(action_type)
	filesystem.UpdateWorkCopy(action_type)
}

type Brnc_chekcout_ingr struct {
}

func (c *Brnc_chekcout_ingr) Aboutme() string {
	return "Initialize the project"
}
func (c *Brnc_chekcout_ingr) Execute() {
	require_branch_name()
	action_type := os.Args[3]
	filesystem.Setbranchstatus(action_type)
	filesystem.DeleteLocalRootFile(action_type)
	filesystem.UpdateWorkCopy(action_type)
}

type Brnc_create struct {
}

func (c *Brnc_create) Aboutme() string {
	return "Initialize the project"
}
func (c *Brnc_create) Execute() {
	require_branch_name()
	action_type := os.Args[3]

	filesystem.Setbranchstatus(action_type)
	filesystem.UpdateLocalCopy(action_type, true)
}

type Brnc_backup struct {
}

func (c *Brnc_backup) Aboutme() string {
	return "Initialize the project"
}
func (c *Brnc_backup) Execute() {
	branch := filesystem.Getbranchstatus()
	filesystem.UpdateLocalCopy(branch.Name, false)
}

type Brnc_restore struct {
}

func (c *Brnc_restore) Aboutme() string {
	return "Initialize the project"
}
func (c *Brnc_restore) Execute() {
	branch := filesystem.Getbranchstatus()
	filesystem.UpdateWorkCopy(branch.Name)
}

type Brnc_merge struct {
}

func (c *Brnc_merge) Aboutme() string {
	return "Initialize the project"
}
func (c *Brnc_merge) Execute() {
	require_branch_name()
	action_type := os.Args[3]
	branch := filesystem.Getbranchstatus()
	if action_type == branch.Name {
		fmt.Println("Your current branch name can not be MERGE on your prefered branch")
		os.Exit(0)
	}
	filesystem.UpdateWorkCopy(action_type)
}

type Brnc_delete struct {
}

func (c *Brnc_delete) Aboutme() string {
	return "Initialize the project"
}
func (c *Brnc_delete) Execute() {
	require_branch_name()
	action_type := os.Args[3]
	branch := filesystem.Getbranchstatus()
	if action_type == branch.Name {
		fmt.Println("Your current branch name can not be deleted on your prefered branch")
		os.Exit(0)
	}
	filesystem.DeleteLocalCopy(action_type)
}

func require_branch_name() {
	args := os.Args
	if len(args) <= 3 {
		fmt.Println("Please specify your prefered branch name")
		os.Exit(0)
	}

}
