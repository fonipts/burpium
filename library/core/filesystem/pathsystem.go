package filesystem

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Glob(root string, dir_replace string, fn func(string) bool) []string {
	var files []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if fn(s) {
			files = append(files, s)
			new_dir := strings.Replace(s, root, dir_replace, 1)

			if stat, err := os.Stat(s); err == nil && stat.IsDir() {

				_ = os.Mkdir(new_dir, 0750)
				//if err_dir_backup_mkdir != nil {
				//	//fmt.Println(err_dir_backup_mkdir)
				//}
			} else {
				b, err := os.ReadFile(s)
				if err != nil {
					log.Fatal(err)
				}

				// `b` contains everything your file has.
				// This writes it to the Standard Out.
				//os.Stdout.Write(b)

				// You can also write it to a file as a whole.
				_ = os.WriteFile(new_dir, b, 0755)
				//if err != nil {
				//	log.Fatal(err)
				//}

			}
		}
		return nil
	})
	return files
}

func GlobDeleteFile(root string, dir_replace string, fn func(string) bool) []string {
	var files []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if fn(s) {
			files = append(files, s)
			//new_dir := strings.Replace(s, root, dir_replace, 1)

			if stat, err := os.Stat(s); err == nil && stat.IsDir() {
				//fmt.Println(new_dir + "r1::" + s)
				_ = os.RemoveAll(s)
				//if err_dir_backup_mkdir != nil {
				//	//fmt.Println(err_dir_backup_mkdir)
				//}
			} else {

				// `b` contains everything your file has.
				// This writes it to the Standard Out.
				//os.Stdout.Write(b)

				// You can also write it to a file as a whole.
				//	fmt.Println(new_dir + "r2::" + s)
				_ = os.Remove(s)
				//if err != nil {
				//	log.Fatal(err)
				//}

			}
		}
		return nil
	})
	return files
}
