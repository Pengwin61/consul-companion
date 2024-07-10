package core

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func gracefulShutdown() {
	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	log.Println("stopping application:", sign)
	os.Exit(0)

}

func getFileChecksum(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func DiffChecksum(file1 string, file2 string) bool {

	if _, err := os.Stat(file1); os.IsNotExist(err) {
		return false

	}

	if _, err := os.Stat(file2); os.IsNotExist(err) {
		return false
	}

	checksum1, err := getFileChecksum(file1)
	if err != nil {
		fmt.Println("Error calculating checksum for file1:", err)
		return false
	}

	checksum2, err := getFileChecksum(file2)
	if err != nil {
		fmt.Println("Error calculating checksum for file2:", err)
		return false
	}

	if checksum1 == checksum2 {
		deleteFile(file2)
		return true
	} else {
		return false
	}
}

func deleteFile(file string) {

	err := os.Remove(file)
	if err != nil {
		fmt.Println("Error deleting file:", err)
	}
}

func mkDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			fmt.Println("Error creating directory:", err)
		}
	}
}
