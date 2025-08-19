package fileio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	errordemo "github.com/law-lee/go-essentials/error-demo"
)

const (
	EXAMPLE_FILE = "./fileio/example.txt"
	FOO_FILE     = "./fileio/foo.txt"
	ROOT_DIR = "."
)

// ReadLines reads all lines from a file
func ReadLines(filePath string) ([]string, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

// IsPathxists returns true if a given path exists, false if it doesn't.
// It might return an error if e.g. file exists but you don't have
// access
func IsPathExists(path string) (bool, error) {
	_, err := os.Lstat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	// error other than not existing e.g. permission denied
	return false, err
}

// CopyFile copies a src file to dst
func CopyFile(dst, src string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	_, err = io.Copy(dstFile, srcFile)
	err2 := dstFile.Close()
	if err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		// delete the destination if copy failed
		os.Remove(dst)
	}
	return err
}
func Run() {
	lines, err := ReadLines(EXAMPLE_FILE)
	errordemo.PanicIfErr(err)
	fmt.Printf("read lines from file with os.OpenFile: %v\n", lines)

	// open file for writing
	// 	If file doesn’t exist, it’ll be created.
	// If file does exist, it’ll be truncated.
	f, err := os.Create(FOO_FILE)
	errordemo.PanicIfErr(err)

	_, err = f.WriteString("this is written by os.Create\n")
	errordemo.PanicIfErr(err)
	f.Close()

	// open file for appending
	// When you open for reading, use os.Open.
	// When you open for writing to a new file, use os.Create.
	// When you open for appending to existing file, use os.OpenFile with the following flags: * os.O_WRONLY
	// Could also be os.RDWR if we also to both read and write
	// os.O_APPEND means that if file exists, we’ll append
	// os.O_CREATE means that if file doesn’t exist, we’ll create it. Without this flag opening non-existing file would fail
	f, err = os.OpenFile(FOO_FILE, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	errordemo.PanicIfErr(err)

	_, err = f.WriteString("this line is written by os.OpenFile\n")
	errordemo.PanicIfErr(err)
	f.Close()

	st, err := os.Stat(EXAMPLE_FILE)
	errordemo.PanicIfErr(err)
	fmt.Printf(`Name: %s
	Size: %d
	IsDir: %v
	Mode: %x
	ModTime: %s
	OS info: %#v
		`, st.Name(), st.Size(), st.IsDir(), st.Mode(), st.ModTime(), st.Sys())

	exist, err := IsPathExists(FOO_FILE)
	errordemo.PanicIfErr(err)
	fmt.Printf("%s file exists: %t\n", FOO_FILE, exist)

	err = os.Remove(FOO_FILE)
	errordemo.PanicIfErr(err)
	fmt.Printf("删除 %s 成功\n", FOO_FILE)

	// list files in directory
	entrys, err := os.ReadDir(ROOT_DIR)
	errordemo.PanicIfErr(err)

	for _, entry := range entrys {
		fi, err := entry.Info()
		errordemo.PanicIfErr(err)
		fmt.Printf("Path: %s, is dir: %v, size: %d bytes\n", fi.Name(), fi.IsDir(),fi.Size())
	}

	// list files recursively
	err = filepath.Walk(ROOT_DIR, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("Visited: %s, is dir: %v, fileName: %s, fileSize: %d\n", path, info.IsDir(),info.Name(), info.Size())
		return nil
	})
	errordemo.PanicIfErr(err)

	// file path operation
	path := filepath.Join("a","b","c","foo.txt")
	file := filepath.Base(path)
	fmt.Printf("get file name only from filepath.Base: %s\n", file)
	parts := filepath.SplitList("/usr/bin:/bin")
	fmt.Printf("parts from SplitList: %#v\n",parts)
	dir, file := filepath.Split(path)
	fmt.Printf("get dir name: %s, get file name: %s with Split\n", dir, file)
	dir = filepath.Dir(path)
	fmt.Printf("get dir name: %s with Dir\n", dir)
	ext := filepath.Ext(path)
	fmt.Printf("get file extention: %s with Ext\n", ext)
}
