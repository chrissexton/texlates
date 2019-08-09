package main

import (
	"flag"
	"log"
	"os"
	"path"
	"text/template"
	"time"
)

var (
	fileName  = flag.String("file", "output.tex", "Name of output file")
	tplPath   = flag.String("tpl", "templates/daily-planner", "Path to templates")
	startDate = flag.String("start", "", "date of start (yyy-mm-dd)")
	days      = flag.Int("days", 7, "number of days")

	output = os.Stdout
)

func main() {
	flag.Parse()
	head := template.Must(template.ParseFiles(path.Join(*tplPath, "head.tpl")))
	body := template.Must(template.ParseFiles(path.Join(*tplPath, "body.tpl")))
	foot := template.Must(template.ParseFiles(path.Join(*tplPath, "foot.tpl")))

	if *startDate == "" {
		*startDate = time.Now().Format("2006-01-02")
	}
	day, err := time.Parse("2006-01-02", *startDate)
	check(err)

	check(head.Execute(output, nil))
	for i := 0; i < *days; i++ {
		info := map[string]interface{}{
			"Date":      day.Format("2 Jan"),
			"DayOfWeek": day.Format("Monday"),
		}
		check(body.Execute(output, info))
		day = day.Add(24 * time.Hour)
	}
	check(foot.Execute(output, nil))

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
