package main

var buffer []string

var wp = 0
var rp = 0

func NewBuffer() {
	buffer = make([]string, 5)
	wp, rp = 0, 0
}

func IsEmpty() bool {
	return wp == rp
}

func Insert(d string){
	buffer[wp] = d
	wp++
}

func Remove() string {
	defer func() {
		rp++
	}()
	return buffer[rp]
}
