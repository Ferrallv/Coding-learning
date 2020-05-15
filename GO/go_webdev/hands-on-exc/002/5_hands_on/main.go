// I started with this answer
// But the answer provided helps

// package main

// import (
// 	"log"
// 	"os"
// 	// "text/template"
// 	"encoding/csv"
// 	"fmt"
// 	// "strings"
// 	"io"
// )

// func main() {
// 	f, err := os.Open("table.csv")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	the_file := csv.NewReader(f)

// 	for i := 0; i < 20; i++ {
// 		record, err := the_file.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(record[0], record[1], record[4])
// 	}
// }

package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	// parse csv
	records := prs("table.csv")

	// parse template
	tpl, err := template.ParseFiles("hw.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute template
	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}

func prs(filePath string) []Record {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]Record, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}

	return records

}