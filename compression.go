package main

import (
	"fmt"
	"os/exec"
)

// some info can be found here:
// https://gist.github.com/firstdoit/6390547

func compress_pdf(inputPdf, outputPdf, pdfSettings string, c chan bool) {

	fmt.Println("Processing: ", inputPdf)

	if _, err :=
		exec.Command("gs", "-q",
			"-dNOPAUSE",
			"-dBATCH",
			"-dSAFER",
			"-sDEVICE=pdfwrite",
			"-dCompatibilityLevel=1.3",
			fmt.Sprintf("-dPDFSETTINGS=/%s", pdfSettings),
			"-dEmbedAllFonts=true",
			"-dSubsetFonts=true",
			"-dSubsetFonts=true",
			"-dColorImageDownsampleType=/Bicubic",
			"-dColorImageResolution=144",
			"-dGrayImageDownsampleType=/Bicubic",
			"-dGrayImageResolution=144",
			"-dMonoImageDownsampleType=/Bicubic",
			"-dMonoImageResolution=144",
			fmt.Sprintf("-sOutputFile=%s", outputPdf), inputPdf).Output(); err != nil {
		fmt.Print("Error executing compression:", err)
	}

	c <- true
}

/*
if _, err := exec.Command("gs", "-sDEVICE=pdfwrite",
	"-dNOPAUSE", "-dQUIET", "-dBATCH", fmt.Sprintf("-dPDFSETTINGS=/%s", pdfSettings),
	"-dCompatibilityLevel=1.6", fmt.Sprintf("-sOutputFile=%s", outputPdf), inputPdf).Output(); err != nil {
	fmt.Print("Error executing compression:", err)
}*/
