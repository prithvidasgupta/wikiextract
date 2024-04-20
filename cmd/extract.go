/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/xml"
	"wikiextract/extract"
	"wikiextract/xsd"

	"github.com/spf13/cobra"
)

func convertXMLtoStruct(byteStr []byte) *xsd.Page {
	var page xsd.Page
	err := xml.Unmarshal(byteStr, &page)
	if err != nil {
		panic(err)
	}
	return &page
}

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract the file to the necessary format",
	Run: func(cmd *cobra.Command, args []string) {
		file := extract.File{InPath: filePath, OutPath: outPath}
		defer file.Close()

		var (
			isPage         bool
			pageBuffer     []byte
			openPageBytes  = []byte("  <page>")
			closePageBytes = []byte("  </page>")
		)

		scanner := file.Validate().Open().Decompress()
		file.WriteRow([]string{"id", "title", "text"})
		for scanner.Scan() {
			currLine := scanner.Bytes()
			if bytes.Equal(currLine, openPageBytes) {
				isPage = !isPage
			} else if bytes.Equal(currLine, closePageBytes) {
				pageBuffer = append(pageBuffer, currLine...)
				isPage = !isPage
				//TODO Need to figure out how to process this in goroutine
				// Only preserve articles for now
				page := convertXMLtoStruct(pageBuffer)
				if page.Namespace == 0 && page.Redirect == (xsd.RedirectType{}) {
					page = page.RemoveTemplates().
						RemoveTables().
						HandleInternalLinks().
						HandleExternalLinks().
						HandleHTMLTags()
					file.BatchWrite([]string{page.Id, page.Title, page.Revisions[0].Text})
				}
				pageBuffer = pageBuffer[:0]
			}
			if isPage {
				currLine = append(currLine, '\n')
				pageBuffer = append(pageBuffer, currLine...)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
}
