package main

import (
	"github.com/jung-kurt/gofpdf"
	"fmt"
)

//generate pdf from html tags

func main()  {
	HTMLBasicNew()
}

func HTMLBasicNew() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	// First page: manual local link
	pdf.SetFont("Helvetica", "", 20)
	_, lineHt := pdf.GetFontSize()
	pdf.AddPage()
	pdf.Image("/tmp/golang-logo.png", 10, 12, 30, 0, false, "", 0, "https://blog.golang.org/gopher")
	pdf.SetLeftMargin(45)
	pdf.SetFontSize(14)
	_, lineHt = pdf.GetFontSize()
	htmlStr := `You are about creating pdf file from golang: <b>GO IS AWESOME</b>, ` +
		`<i>GO AWESOME</i>, <u>beyond your imagination</u>, or <b><i><u>push forward your creativity</u></i></b>!<br><br>` +
		`<center>Thanks to creator and strong lovers.</center>` +
		`<right>We can just coding with brave and fun.</right>` +
		`for more samples and working code ` +
		`<a href="http://emrizki.com">www.emrizki.com</a>`
	html := pdf.HTMLBasicNew()
	html.Write(lineHt, htmlStr)
	fileStr := "/tmp/HTMLBasicNew.pdf"
	err := pdf.OutputFileAndClose(fileStr)
	fmt.Println(err)
}
