package main


import (
	"time"
	"fmt"
	"strings"
)

//time formatting
func main() {

	today := time.Now()
	fmt.Print(today.Weekday(), ", ", today.Month(), today.Day(), today.Year(), " at ", today.Hour(), ":", today.Minute(), ":", today.Second())
	fmt.Println()

	fmt.Println("Go Date Format(Today - 'yyyy-mmmm-dd HH:mm:ss Z'): ", GetTodayIndonesia(ConvertFormatNow("dddd, dd mmmm yyyy, HH:mm:ss Z")))
	fmt.Println("Go Date Format(Today - 'yyyy-MM-dd HH:mm:ss Z'): ", GetToday(ConvertFormatNow("yyyy-mm-dd HH:mm:ss Z")))
	fmt.Println("Go Date Format(Today - 'yyyy-MMM-dd'): ", GetToday(ConvertFormatNow("yyyy-MMM-dd")))
	fmt.Println("Go Time Format(NOW - 'HH:MM:SS'): ", GetToday(ConvertFormatNow("HH:MM:SS")))
	fmt.Println("Go Time Format(NOW - 'HH:MM:SS tt'): ", GetToday(ConvertFormatNow("HH:MM:SS tt")))
	
}


func GetToday(format string) (todayString string){
	today := time.Now()
	todayString = today.Format(format);
	return
}

func GetTodayIndonesia(format string) (todayString string){
	today := time.Now()
	todayString = today.Format(format)
	var replacer = strings.NewReplacer("January", "Januari",
		"February", "Februari",
		"March", "Maret",
		"April", "April",
		"May", "Mei",
		"June", "Juni",
		"July", "Juli",
		"August", "Agustus",
		"September", "September",
		"October", "Oktober",
		"November", "November",
		"December", "Desember",
		"Monday", "Senin",
		"Tuesday", "Selasa",
		"Wednesday", "Rabu",
		"Thursday", "Kamis",
		"Friday", "Jumat",
		"Saturday", "Sabtu",
		"Sunday", "Minggu")
	str := todayString
	str = replacer.Replace(str)
	return str
}

func ConvertFormatNow(format string) (string){

	const  (
		yyyy = "2006"
		yy = "06"
		mmmm = "January"
		mmm = "Jan"
		mm = "01"
		dddd = "Monday"
		ddd = "Mon"
		dd = "02"

		HHT = "03"
		HH = "15"
		MM = "04"
		SS = "05"
		ss = "05"
		tt = "PM"
		Z = "MST"
		ZZZ = "MST"
	)

	result := format
	var replacer = strings.NewReplacer("YYYY", yyyy,
		"yyyy", yyyy,
		"YY", yy,
		"yy", yy,
		"MMMM", mmmm,
		"mmmm", mmmm,
		"MMM", mmm,
		"mmm", mmm,
		"mm", mm,
		"dddd", dddd,
		"ddd", ddd,
		"dd", dd,
		"HH", HHT,
		"hh", HHT,
		"tt", tt,
		"MM", MM,
		"SS", SS,
		"ss", ss,
		"ZZZ", ZZZ,
		"zzz", ZZZ,
		"Z", Z,
		"z", Z)
	result = replacer.Replace(result)
	return (result)
}
