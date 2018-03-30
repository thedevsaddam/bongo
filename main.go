package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/thedevsaddam/ponjika"
)

func main() {
	now := time.Now()
	p := ponjika.New(now)
	mnthTotalDys := p.TotalDays
	crntDate := p.Date
	crntDateIndx := int(now.Weekday())
	frstDateIndx := firstDayIndex(crntDate, crntDateIndx)

	totalRow := int(math.Ceil(float64(mnthTotalDys) / 7.0))
	calender := [6][7]int{}

	fstIndex := frstDateIndx
	tmpDate := 1
	for i := 0; i < totalRow; i++ {
		for j := 0; j < 7; j++ {
			if j >= fstIndex && tmpDate <= mnthTotalDys {
				calender[i][j] = tmpDate
				tmpDate++
			}
		}
		fstIndex = 0
	}

	isPh := flag.Bool("p", false, "show bengali calendar in english phonetic")
	flag.Parse()

	if *isPh {
		heading := fmt.Sprintf("   %s %s", p.BengaliMonth.Phonetic, p.BengaliYear.Phonetic)
		fmt.Println(bold(heading))
		phoneticBengaliCalendar(calender, crntDate)
		os.Exit(1)
	}
	heading := fmt.Sprintf("\t     %s %s", p.BengaliMonth.Bengali, p.BengaliYear.Bengali)
	fmt.Println(bold(heading))
	bengaliCalender(calender, crntDate)
}

func bengaliCalender(calender [6][7]int, today int) {
	fmt.Println("রবি   সোম   মঙ্গল  বুধ   বৃহ   শুক্র   শনি")
	for i := 0; i < len(calender); i++ {
		for j := 0; j < 7; j++ {
			dt := calender[i][j]
			if dt != 0 {
				if dt > 9 {
					if dt == today {
						fmt.Print(highlight(enToBnNumber(dt)), "   ")
					} else {
						fmt.Print(enToBnNumber(dt), "   ")
					}
				} else {
					fmt.Print(enToBnNumber(dt), "    ")
				}
			} else {
				fmt.Print("     ")
			}
		}
		fmt.Println()
	}
}

func phoneticBengaliCalendar(calender [6][7]int, today int) {
	fmt.Println("S  M  T  W  T  F  S")
	for i := 0; i < len(calender); i++ {
		for j := 0; j < 7; j++ {
			dt := calender[i][j]
			if dt != 0 {
				if dt > 9 {
					if dt == today {
						fmt.Print(highlight(fmt.Sprintf("%d", dt)), " ")
					} else {
						fmt.Print(dt, " ")
					}
				} else {
					fmt.Print(dt, "  ")
				}
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

// firstDayIndex return the months first day index
func firstDayIndex(date, dateIndex int) int {
	a := date - (dateIndex + 1)
	r := a % 7
	result := (6 - r) + 1
	if result >= 7 {
		return result - 7
	}
	return result
}

// enToBnNumber covert english number to bengali number String
func enToBnNumber(d int) string {
	var o string
	ds := fmt.Sprintf("%v", d)
	bdmap := map[string]string{
		"1": "১",
		"2": "২",
		"3": "৩",
		"4": "৪",
		"5": "৫",
		"6": "৬",
		"7": "৭",
		"8": "৮",
		"9": "৯",
		"0": "০",
	}
	for i := 0; i < len(ds); i++ {
		o += bdmap[string(ds[i])]
	}
	return o
}

func highlight(in string) string {
	lg := "\033[1;32m" // light green
	nc := "\033[0m"
	return fmt.Sprintf("%s%s%s", lg, in, nc)
}

func bold(in string) string {
	bld := "\033[1m" // bold
	nb := "\033[0m"
	return fmt.Sprintf("%s%s%s", bld, in, nb)
}
