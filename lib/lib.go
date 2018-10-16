package lib

import (
	"os"
	"fmt"
	"strings"
	"time"
	"unicode"
	"strconv"
)

const (
	debug = false
)

type LibError struct {
	When time.Time
	What string
}
func (e LibError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}
func oops() error {
	return LibError{
		time.Now(),
		"error in lib.go",
	}
}
func GetTimeframeHours(t1 string, t2 string) float64{
	// 10:30, 17:30
	// returns timeframe in hours  e.g.: 2.8

	hhmm1 := strings.Split(t1, ":")
	hhmm2 := strings.Split(t2, ":")
	hh1, err := strconv.Atoi(hhmm1[0])
	mm1, err := strconv.Atoi(hhmm1[1])
	hh2, err := strconv.Atoi(hhmm2[0])
	mm2, err := strconv.Atoi(hhmm2[1])
		
	if err /*:= oops(); err*/ != nil {
		fmt.Println(err)
	}

	t := time.Now()
	berlinTime, err := time.LoadLocation("Europe/Berlin")
	s1 := time.Date(t.Year(), t.Month(), t.Day(), hh1, mm1, 0, 0, berlinTime)
	s2 := time.Date(t.Year(), t.Month(), t.Day(), hh2, mm2, 0, 0, berlinTime)

	//d := s2.Since(s1)
	d := s2.Sub(s1)
	return d.Hours()
}

func GetTimePercentage(s string, e string) string{
	t1 := GetLastLineOfFile(s)
	t2 := GetLastLineOfFile(e)
	now := time.Now()
	
	t := now.Format( "15:04" )

	timeOver := GetTimeframeHours(t1, t) 
	timeFrame := GetTimeframeHours(t1, t2)

	p := FloatToString(timeOver / timeFrame)
	if(debug){	
		fmt.Printf("start: %s\n", t1)
		fmt.Printf("end: %s\n",t2)
		fmt.Printf("now: %s\n",t)
		fmt.Printf("percentage: %s\n",p)
	}
	return p
}

func IsLetter(s string) bool {
    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func FloatToString(input_num float64) string {
    return strconv.FormatFloat(input_num, 'f', 6, 64)
}


func Join(strs ...string) string {
	b := strings.Builder{}
	for _, str := range strs {
		b.WriteString(str)
	}
	r := b.String()
	return r
}
func WriteToFile(file string, content string){
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(content); err != nil {
		panic(err)
	}
	// new line
	if _, err = f.WriteString("\n"); err != nil {
		panic(err)
	}
}

func GetLastLineOfFile(fname string) string {
	file, err := os.Open(fname)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    buf := make([]byte, 62)
    stat, err := os.Stat(fname)
    start := stat.Size() - 62
    _, err = file.ReadAt(buf, start)
	
	lines := string(buf)

	if(debug){
		fmt.Printf("%s\n", lines)
	}
	l := strings.Split(lines, "\n")
	ll := l[len(l)-2]
	return ll
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}