package streamcode

import (
	"fmt"
	"strings"
	"io"
)

// MyStringData - simple struct to hold string data
type MyStringData struct {
	str string
	readIndex int // default: 0
}

// add `Read` method (to impletement io.Reader)
func (myStringData *MyStringData) Read(p []byte) (n int, err error) {

	// convert `str` string to a slice of bytes
	strBytes := []byte(myStringData.str)

	// if `readIndex` is GTE source length, return `EOF` error
	if myStringData.readIndex >= len(strBytes) {
		return 0, io.EOF // `0` bytes read
	}

	// get next readable limit (exclusive)
	nextReadLimit := myStringData.readIndex + len(p)

	// if `nextReadLimit` is GTE source length
	// set `nextReadLimit` to source length and `err` to `EOF`
	if( nextReadLimit >= len(strBytes) ) {
		nextReadLimit = len(strBytes)
		err = io.EOF
	}

	// get next bytes to copy and set `n` to its length
	nextBytes := strBytes[ myStringData.readIndex : nextReadLimit ]
	n = len(nextBytes)

	// copy all bytes of `nextBytes` into `p` slice
	copy(p, nextBytes)

	// increment `readIndex` to `nextReadLimit`
	myStringData.readIndex = nextReadLimit

	// return values
	return

}




func Streamcode() {
	
	// create data source
	src := MyStringData{
		str: "Hello Amazing World!",
	}

	// create a packet
	p := make([]byte, 3) // slice of length `3`

	// read `src` until an error is returned
	for {

		// read `p` bytes from `src`
		n, err := src.Read(p);
		fmt.Printf( "%d bytes read, data: %s\n", n, p[:n] )

		// handle error
		if err == io.EOF {
			fmt.Println( "--end-of-file--" )
			break;
		} else if err != nil {
			fmt.Println( "Oops! Some error occured!", err )
			break;
		}
	}

}


func LimitReaderFunc() {
	
	// create a main data source
	mainSrc := strings.NewReader("Hello Amazing World!") // 20 characters

	// create data source from `mainSrc` with cap of `10` bytes
	src := io.LimitReader(mainSrc, 10)

	// create a packet
	p := make([]byte, 4) // slice of length `3`

	// read `src` until an error is returned
	for {

		// read `p` bytes from `src`
		n, err := src.Read(p);
		fmt.Printf( "%d bytes read, data: %s\n", n, p[:n] )

		// handle error
		if err == io.EOF {
			fmt.Println( "--end-of-file--" )
			break;
		} else if err != nil {
			fmt.Println( "Oops! Some error occured!", err )
			break;
		}
	}

}