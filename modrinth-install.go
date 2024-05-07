package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	url := "https://launcher-files.modrinth.com/versions/0.7.1/linux/modrinth-app_0.7.1_amd64.AppImage"
	fileName := "modrinth-app_0.7.1_amd64.AppImage"
	downloadDir := "Modrinth"

	// Create the Modrinth directory if it doesn't exist
	if _, err := os.Stat(downloadDir); os.IsNotExist(err) {
		err := os.Mkdir(downloadDir, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	filePath := filepath.Join(downloadDir, fileName)

	out, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer out.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Modrinth downloaded successfully")

	// Move the downloaded file to the Modrinth directory
	err = os.Rename(fileName, filepath.Join(downloadDir, fileName))
	if err != nil {
		fmt.Println("Error moving file:", err)
		return
	}

	fmt.Println("File moved to Modrinth directory successfully")
}

