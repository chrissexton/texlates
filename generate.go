package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

var (
	fileName   = flag.String("file", "output", "Name of output file")
	tplPath    = flag.String("tpl", "daily", "Path to templates")
	startDate  = flag.String("start", "", "date of start (yyy-mm-dd)")
	courseDays = flag.String("courseDays", "Monday,Wednesday,Other", "days for the class template")
	courses    = flag.String("courses", "Cxxx,Cyyy,Czzz", "courses for the class template")
	pages      = flag.Int("pages", 7, "number of pages")
	debug      = flag.Bool("debug", false, "leave LaTeX log files")
)

func main() {
	flag.Parse()
	tmpfile, err := ioutil.TempFile("", "texgen")
	defer os.Remove(tmpfile.Name()) // clean up
	if err != nil {
		log.Fatal(err)
	}
	switch *tplPath {
	case "daily":
		daily(tmpfile)
	case "class":
		class(tmpfile)
	default:
		log.Fatal("Unknown template")
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("xelatex", "-halt-on-error", "-jobname="+*fileName, tmpfile.Name())
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	if !*debug {
		os.Remove(*fileName + ".aux")
		os.Remove(*fileName + ".log")
	}

}

func class(file *os.File) {
	head := template.Must(template.ParseFiles("templates/class/head.tpl"))
	body := template.Must(template.ParseFiles("templates/class/body.tpl"))
	foot := template.Must(template.ParseFiles("templates/class/foot.tpl"))

	if *startDate == "" {
		*startDate = time.Now().Format("2006-01-02")
	}
	day, err := time.Parse("2006-01-02", *startDate)
	check(err)

	check(head.Execute(file, nil))
	for i := 0; i < *pages; i++ {
		info := map[string]interface{}{
			"Courses":   strings.Split(*courses, ","),
			"Days":      strings.Split(*courseDays, ","),
			"DateRange": fmt.Sprintf("%s-%s", day.Format("2006/01/02"), day.Add(48*time.Hour).Format("02")),
			"Week":      fmt.Sprintf("Week %d", i+1),
		}
		check(body.Execute(file, info))
		day = day.Add(24 * 7 * time.Hour)
	}
	check(foot.Execute(file, nil))
}

func daily(file *os.File) {
	head := template.Must(template.ParseFiles("templates/daily/head.tpl"))
	body := template.Must(template.ParseFiles("templates/daily/body.tpl"))
	foot := template.Must(template.ParseFiles("templates/daily/foot.tpl"))

	if *startDate == "" {
		*startDate = time.Now().Format("2006-01-02")
	}
	day, err := time.Parse("2006-01-02", *startDate)
	check(err)

	check(head.Execute(file, nil))
	for i := 0; i < *pages; i++ {
		info := map[string]interface{}{
			"Date":      day.Format("2 Jan"),
			"DayOfWeek": day.Format("Monday"),
		}
		check(body.Execute(file, info))
		day = day.Add(24 * time.Hour)
	}
	check(foot.Execute(file, nil))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
