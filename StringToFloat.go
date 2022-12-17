package StringToFloat

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func checknagtive(s string) (string, bool) {
	if strings.Contains(s, "-") {
		return s[1:], true
	} else {
		return s, false
	}
}

func GetFloat64(s string) float64 {
	var (
		IntAndFloat    []string
		Integer, Float int
		Result         float64
		err            error
		Negative       bool
	)
	// check has nagtive
	s, Negative = checknagtive(s)

	// Split Integer and Float Part
	IntAndFloat = strings.Split(s, ".")

	// convert string to integer
	Integer, err = strconv.Atoi(IntAndFloat[0])
	if err != nil {
		log.Println(err)
	}
	Float, err = strconv.Atoi(IntAndFloat[1])
	if err != nil {
		log.Println(err)
	}

	// add Integer and Float Part
	Result = float64(Integer) + float64(Float)/math.Pow(10, float64(len(IntAndFloat[1])))

	// check orgin is nagtive
	if Negative {
		Result *= -1
	}
	return Result
}
