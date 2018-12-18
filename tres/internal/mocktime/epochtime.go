package mocktime

import "fmt"
import "time"
import "bytes"
import "strconv"

func SimpleRange(start, end string) string {

	var buffer bytes.Buffer
	p := buffer.WriteString

	// Parse example for ISO time
	//t, err := time.Parse(time.RFC3339, "2013-06-05T14:10:43.678Z")
	// use above to get the NOW and THEN from passed arguements

	// We'll start by getting the current time.
	// You can build a `time` struct by providing the
	// year, month, day, etc. Times are always associated
	// with a `Location`, i.e. time zone.
	//now := time.Now()
	//then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	now, err := time.Parse(time.RFC3339, start)
	then, err := time.Parse(time.RFC3339, end)
	if err != nil {
		fmt.Println(err)
	}

	p(then.String() + "\n")
	p(now.String() + "\n")

	// You can extract the various components of the time
	// value as expected.
	p(strconv.Itoa(then.Year()) + "\n")
	p(then.Month().String() + "\n")
	p(strconv.Itoa(then.Day()) + "\n")
	p(strconv.Itoa(then.Hour()) + "\n")
	p(strconv.Itoa(then.Minute()) + "\n")
	p(strconv.Itoa(then.Second()) + "\n")
	p(strconv.Itoa(then.Nanosecond()) + "\n")
	p(then.Location().String() + "\n")

	// The Monday-Sunday `Weekday` is also available.
	p(then.Weekday().String() + "\n")

	// These methods compare two times, testing if the
	// first occurs before, after, or at the same time
	// as the second, respectively.
	p(strconv.FormatBool(then.Before(now)) + "\n")
	p(strconv.FormatBool(then.After(now)) + "\n")
	p(strconv.FormatBool(then.Equal(now)) + "\n")

	// The `Sub` methods returns a `Duration` representing
	// the interval between two times.
	diff := now.Sub(then)
	p(diff.String() + "\n")

	// We can compute the length of the duration in
	// various units.
	p(floatToString(diff.Hours()) + "\n")
	p(floatToString(diff.Minutes()) + "\n")
	p(floatToString(diff.Seconds()) + "\n")
	p(strconv.FormatInt(diff.Nanoseconds(), 10) + "\n")

	// You can use `Add` to advance a time by a given
	// duration, or with a `-` to move backwards by a
	// duration.
	p(then.Add(diff).String() + "\n")
	p(then.Add(-diff).String() + "\n")

	//fmt.Println(buffer.String())
	return buffer.String()

}

func floatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
