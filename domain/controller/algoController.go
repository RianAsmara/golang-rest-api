package controller

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func sockMerchant(n int, arr []int) int {
	counts := make(map[int]int)
	pairs := 0

	for _, sock := range arr {
		counts[sock]++
		if counts[sock] == 2 {
			pairs++
			delete(counts, sock) // Remove pairs to avoid double-counting
		}
	}

	return pairs
}

func PairColor(c echo.Context) error {
	nStr := c.FormValue("n")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid n"})
	}

	arrStr := c.FormValue("arr")
	arrStrSlice := strings.Split(arrStr, " ")
	arrInt := make([]int, len(arrStrSlice))
	for i, numStr := range arrStrSlice {
		arrInt[i], err = strconv.Atoi(numStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid arr"})
		}
	}

	result := sockMerchant(n, arrInt)
	return c.JSON(http.StatusOK, map[string]int{"pairs": result})
}

func countingValleys(steps int, path string) int {
	valleys := 0
	curLevel := 0

	for _, step := range path {
		if step == 'U' {
			curLevel++
			if curLevel == 0 {
				valleys++
			}
		} else if step == 'D' {
			curLevel--
		}
	}

	return valleys
}

func CountValley(c echo.Context) error {
	stepsStr := c.FormValue("steps")
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid steps"})
	}

	path := c.FormValue("path")

	result := countingValleys(steps, path)
	return c.JSON(http.StatusOK, map[string]int{"valleys": result})
}

func ProcessNumber(c echo.Context) error {
	inputNumberStr := c.QueryParam("number")
	inputNumber, err := strconv.ParseFloat(inputNumberStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid number")
	}

	multiplier := 1000000

	for i := 7; i >= 1; i-- {
		currentValue := int(inputNumber * math.Pow10(i-1))

		c.JSON(http.StatusOK, currentValue)

		multiplier /= 10
	}

	finalDigit := int(inputNumber*10) % 10
	c.JSON(http.StatusOK, finalDigit)
	return nil
}

func CountLamps(c echo.Context) error {
	const numSwitches = 100
	lamps := make([]bool, numSwitches)

	// Andrew's trips
	for i := 1; i <= numSwitches; i++ {
		for j := 1; j <= numSwitches; j++ {
			if j%i == 0 {
				lamps[j-1] = !lamps[j-1]
			}
		}
	}

	onLamps := 0
	for _, isOn := range lamps {
		if isOn {
			onLamps++
		}
	}

	return c.JSON(http.StatusOK, map[string]int{"onLamps": onLamps})
}
