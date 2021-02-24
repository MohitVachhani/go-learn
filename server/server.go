package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func playWithMap() {

	// maps are always pass by reference.

	// initialize map
	studentMarks := make(map[string]int)
	studentMarks = map[string]int{
		"mohit": 10,
	}

	// deleting element from map
	delete(studentMarks, "mohit")

	// to check whether such key exists or not.
	studentMark, isPresent := studentMarks["mohit"]

	fmt.Println(studentMark, isPresent)

	fmt.Println(studentMarks)
}

// pass by value, copy is created.
func playWithStruct() {
	type UserEmailSettings struct {
		emailId     string
		toSendEmail bool
	}

	type User struct {
		firstName   string `required max: 100`
		lastName    string
		phoneNumber string
		UserEmailSettings
	}

	user := User{
		firstName:         "Mohit",
		lastName:          "Vachhani",
		phoneNumber:       "9825336634",
		UserEmailSettings: UserEmailSettings{emailId: "mohitvachhani55@gmail.com", toSendEmail: true},
	}

	fmt.Println(user.firstName)
	fmt.Println(user.toSendEmail)
	fmt.Println(user.emailId)
	fmt.Println(user.UserEmailSettings)
}

func playWithCasingAndLopping() {
	var t bool = true
	fmt.Println(t)

	if t == true {
		fmt.Println("t value is true")
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// defer -> value variable same as above
func readWebsite() {
	res, err := http.Get("http://www.google.com/robots/txt")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello go"))
	})

	err := http.ListenAndServe(":4000", nil)

	if err != nil {
		panic(err.Error())
	}
}

func implementWg() {
	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	countThings("hello", 200)
	// 	wg.Done()
	// }()
	// wg.Wait()
}

// function execute manner: normal function, defer, panics, return values

// if panic then first defer than panic is called.

// create a function which generated fib numbers sequence.
func generateFinNumber(number int) int {
	if number <= 1 {
		return number
	}
	return generateFinNumber(number-1) + generateFinNumber(number-2)
}

func worker(n chan int, ans chan int) {

	for numbers := range n {
		ans <- generateFinNumber(numbers)
	}

}

func main() {
	n := make(chan int, 100)
	ans := make(chan int, 100)

	go worker(n, ans)

	for i := 0; i < 100; i++ {
		n <- i
	}

	close(n)

	for i := 0; i < 100; i++ {
		fmt.Println("The answer is: ", <-ans)
	}
}
