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
	fileName   = flag.String("file", "output", "name of output file")
	tplPath    = flag.String("tpl", "daily", "template name")
	startDate  = flag.String("start", "", "date of start (yyyy-mm-dd)")
	courseDays = flag.String("courseDays", "Monday,Wednesday,Other", "days for the class template")
	holidays   = flag.String("holidays", "", "dates (yyyy-mm-dd) comma separated of holidays")
	course     = flag.String("course", "Cxxx", "course for lesson template")
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
	case "lesson":
		lesson(tmpfile)
	default:
		log.Fatal("Unknown template")
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("xelatex", "-halt-on-error", "-jobname="+*fileName, tmpfile.Name())
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stderr
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

func lesson(file *os.File) {
	head := template.Must(template.ParseFiles("templates/lesson/head.tpl"))
	body := template.Must(template.ParseFiles("templates/lesson/body.tpl"))
	foot := template.Must(template.ParseFiles("templates/lesson/foot.tpl"))

	hs := strings.Split(*holidays, ",")
	holdayDates := map[time.Time]bool{}
	for _, h := range hs {
		if h == "" {
			continue
		}
		day, err := time.Parse("2006-01-02", h)
		check(err)
		holdayDates[day] = true
	}

	if *startDate == "" {
		*startDate = time.Now().Format("2006-01-02")
	}
	day, err := time.Parse("2006-01-02", *startDate)
	check(err)

	check(head.Execute(file, nil))
	for i := 0; i < *pages*2; {
		info := map[string]interface{}{
			"Course": *course,
			"Date":   fmt.Sprintf("%s", day.Format("2006/01/02")),
			"Day":    fmt.Sprintf("Day %d", i+1),
		}
		if !holdayDates[day] {
			check(body.Execute(file, info))
			i++
		}
		day2 := day.Add(2 * 24 * time.Hour)
		info["Date"] = fmt.Sprintf("%s", day2.Format("2006/01/02"))
		info["Day"] = fmt.Sprintf("Day %d", i+1)
		if !holdayDates[day2] {
			check(body.Execute(file, info))
			i++
		}
		day = day.Add(24 * 7 * time.Hour)
	}
	check(foot.Execute(file, nil))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
