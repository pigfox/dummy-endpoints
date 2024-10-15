package structs

import "math/rand"

const ResponseRowsPerServer = 500

type Ports struct {
	Min      int   `json:"min"`
	Max      int   `json:"max"`
	Excluded []int `json:"excluded"`
}

type Response struct {
	Message   string `json:"message"`
	TimeStamp string `json:"timestamp"`
	Price     int    `json:"random_int"`
	Address   string `json:"address"`
}

func GetPorts() Ports {
	/*
		return Ports{
			Min:      10001,
			Max:      10999,
			Excluded: []int{10002, 10003, 10064, 10088, 10089},
		}
	*/
	return Ports{
		Min:      10001,
		Max:      10011,
		Excluded: []int{10002, 10003, 10028},
	}
}

func Contains(arr []int, num int) bool {
	for _, value := range arr {
		if value == num {
			return true
		}
	}
	return false
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
