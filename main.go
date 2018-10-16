package main

import (
	"os"
	"fmt"
	"strconv"
	"encoding/csv"
	L "./lib"
	"github.com/jung-kurt/gofpdf"
)

const (
    colCount = 6
    colWd    = 30.0
    marginH  = 15.0
    lineHt   = 5.5
	cellGap  = 2.0
	title = "Time sheet"
)
type cellType struct {
    str  string
    list [][]byte
    ht   float64
}
var (
    cellList [colCount]cellType
    cell     cellType
)

// Date,In,Out,h:m,Time,Rate (by hour),Euro,Budget,Approved,Status,Billable,Customer,Project,Activity,Description,Comment,Location,Tracking Number,Username,cleared
type TimeEntry struct {
    Date string
	In string
	Out string
	Hm string
	Time /*int*/string
	Comment string
}

func main(){
	var fname string
	if len(os.Args)<2{
		fname = "./example.csv"
	} else{
		fname = os.Args[1]
	}

	fmt.Println("reading", fname)

	f, err := os.Open(fname)
    if err != nil {
        panic(err)
    }
	defer f.Close()
	
	lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        panic(err)
	}
	
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 13)
	_, lineHt := pdf.GetFontSize()

	pdf.Cell(40, 10, title)
	pdf.Ln(10)
	var r []TimeEntry

	for _, line := range lines {
        data := TimeEntry{
            Date: line[0],
			In: line[1],
			Out: line[2],
			Hm: line[3],
			Time: line[4],//time,
			Comment: line[15],
		}
		//fmt.Println(data)
		r = append(r, data)
	}
	spaces := "                "
	var timeCounter float64
	for _, e := range r {
		htmlStr := ""
		htmlStr = L.Join(htmlStr, e.Date, spaces)
		htmlStr = L.Join( htmlStr, e.In, spaces)
		htmlStr = L.Join(htmlStr,e.Out, spaces)
		htmlStr = L.Join( htmlStr, e.Hm,spaces)
		htmlStr = L.Join(htmlStr, e.Time, spaces)
		//htmlStr = L.Join(htmlStr, e.Comment,spaces)
		pdf.Write(lineHt, htmlStr)
		pdf.Ln(7)

		timeValue, iErr := strconv.ParseFloat(e.Time, 64)
		if iErr != nil {
			fmt.Println(iErr)
		}
		timeCounter = timeCounter + timeValue
		//fmt.Println(timeValue)
	}
	//fmt.Println(r)
	timeInfo := L.Join("time: ", strconv.FormatFloat(timeCounter, 'f', 6, 64))
	pdf.Write(lineHt, timeInfo)
	pdf.Ln(20)
	pdf.Write(lineHt, "		_____________")

	perr := pdf.OutputFileAndClose("out.pdf")
	if perr != nil{
		fmt.Println(perr)
	}

}