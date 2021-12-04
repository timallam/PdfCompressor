package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {

	inputDirPtr := flag.String("input_dir", ".", "input directory")
	outputDirPtr := flag.String("output_dir", ".", "output directory")
	pdfSettingsPtr := flag.String("pdf_settings", "ebook", "settings for dPDFSETTINGS [default, screen, printer]")

	flag.Parse()

	// Get all input pdf files
	files, err := ioutil.ReadDir(*inputDirPtr)
	if err != nil {
		log.Fatal(err)
	}

	// check if there are files to process
	if len(files) == 0 {
		fmt.Println("Please make sure you have at least 1 pdf file to process")
		os.Exit(1)
	}

	// create output directory if it doesn't exist
	if _, err := os.Stat(*outputDirPtr); os.IsNotExist(err) {
		if err = os.Mkdir(*outputDirPtr, 0755); err != nil {
			fmt.Println("Error creating output directory: ", err)
		}
	}

	// create a channel to communicate with the main thread
	opComplete := make(chan bool)
	pdfFiles := 0

	for _, f := range files {
		inputPdfPath := filepath.Join(*inputDirPtr, f.Name())
		fileExtension := filepath.Ext(inputPdfPath)

		// check we are about to process a pdf file
		if f.IsDir() || fileExtension != ".pdf" {
			continue
		}

		pdfFiles++
		outputPdfPath := filepath.Join(*outputDirPtr, f.Name())

		go compress_pdf(inputPdfPath, outputPdfPath, *pdfSettingsPtr, opComplete)
	}

	for i := 0; i < pdfFiles; i++ {
		<-opComplete
	}
}
