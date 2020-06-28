package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	ip2location "github.com/ip2location/ip2location-go"
	"github.com/valyala/fasthttp"
)

func main() {
	port := os.Getenv("PORT")
	token := os.Getenv("API_TOKEN")

	log.Printf("starting up [port: %s] [token: %s]\n", port, token)

	fileURL := fmt.Sprintf("http://www.ip2location.com/download/?token=%s&file=DB1LITEBIN", token)
	err := downloadFile("ip2location.zip", fileURL)
	if err != nil {
		log.Fatalf("download err: %s", err)
	}

	_, err = unzip("ip2location.zip", "data")
	if err != nil {
		log.Fatalf("unzip err: %s", err)
	}

	log.Println("downloaded and unzipped")
	log.Println("start listening")

	h := requestHandler
	if err := fasthttp.ListenAndServe(fmt.Sprintf(":%s", port), h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func ip4ToCountry(ip string) (string, error) {
	db, err := ip2location.OpenDB("./data/IP2LOCATION-LITE-DB1.BIN")
	if err != nil {
		return "", err
	}

	defer db.Close()

	res, err := db.Get_country_short(ip)
	if err == nil {
		return res.Country_short, nil
	}

	return "", errors.New("country not found")
}

func requestHandler(ctx *fasthttp.RequestCtx) {

	// fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())

	path := ctx.Path()
	country := ""

	if len(path) > 1 {
		path = path[1 : len(path)-1]
		country, _ = ip4ToCountry(string(path))

		if len(country) > 2 {
			country = ""
		}
	}

	fmt.Fprintf(ctx, country)
}

//source: https://golangcode.com/download-a-file-from-a-url/
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
// source: https://golangcode.com/unzip-files-in-go/
func unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
