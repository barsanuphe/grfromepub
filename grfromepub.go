package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"

	b "github.com/barsanuphe/endive/book"
	h "github.com/barsanuphe/endive/helpers"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	// test parameters
	if len(os.Args) == 1 {
		fmt.Println("Please provide an epub filename.")
		return
	}
	epubs := os.Args[1:]
	validEpubs := []string{}
	for i := range epubs {
		if !strings.HasSuffix(strings.ToLower(epubs[i]), ".epub") {
			fmt.Println(epubs[i] + " is not an epub, ignoring.")
			continue
		}
		absPath, err := h.FileExists(epubs[i])
		if err != nil {
			fmt.Println(epubs[i] + ": " + err.Error())
		} else {
			validEpubs = append(validEpubs, absPath)
		}
	}
	// parse epubs and open the goodreads page
	for _, epub := range validEpubs {
		e := b.Epub{Filename:epub}
		metadata, err := e.ReadMetadata()
		if err != nil {
			fmt.Println("Error reading metadata from " + epub + ": " + err.Error())
		}
		if metadata.ISBN != "" {
			fmt.Println(filepath.Base(epub) + ": found ISBN " + metadata.ISBN)
			err = open.Start("https://www.goodreads.com/search?q=" + metadata.ISBN)
			if err != nil {
				fmt.Println("Error: could not open goodreads page for " + filepath.Base(epub))
			}
		} else {
			fmt.Println(filepath.Base(epub) + ": could not find ISBN!")
		}
	}
}
