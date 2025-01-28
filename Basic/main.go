package main

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"

	// "crypto/rand"
	"fmt"
	// "math/big"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const LoginToken string = "sajiblogintoken" // UpperCase for Public

func main() {
	fmt.Println("Welcome to Golang Start Learning")

	// variables()
	// userInput()
	// conversion()
	// myMaths()
	// myTime()
	// myPointers()
	// myArray()
	// mySlices()
	// myMaps()
	// myStructs()
	// ifElse()
	// switchCase()
	// loops()
	// methods()
	// myDefer()
	// files()
	// encodeJson()
	decodeJson()

}

func variables() {
	var username string = "Sajib"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}

func userInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Your name is ", input)

}

func conversion() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our pizza between 1 to 5")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}

// func myMaths() {

// 	//var mynumberOne int = 2
// 	//var mynumberTwo float64 = 4.5

// 	// fmt.Println("The sum is: ", mynumberOne+int(mynumberTwo))

// 	//random number
// 	// rand.Seed(time.Now().UnixNano())
// 	// fmt.Println(rand.Intn(5) + 1)

// 	//random from crypto

// 	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
// 	fmt.Println(myRandomNum)
// }

func myTime() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

func myPointers() {
	var x = 10
	x = 30

	var ptr = &x
	var y = x

	*ptr = 50
	fmt.Println(*ptr)
	fmt.Println(x)
	fmt.Println(y)

	var addOne = func(x *int) {
		*x++
	}

	addOne(&x)
	fmt.Println(x)
}

func myArray() {
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", vegList)
	fmt.Println("Vegy list is: ", len(vegList))
}

func mySlices() {
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

func myMaps() {
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}

func myStructs() {
	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v.", hitesh.Name, hitesh.Email)
}

func ifElse() {
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}
}

func switchCase() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
		fallthrough
	case 4:
		fmt.Println("you can move to 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("What was that!")
	}
}

func loops() {
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("index is  and value is %v\n", day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		// if rougueValue == 3 {
		// 	break
		// }

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		if rougueValue == 7 {
			goto lco
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeonline.in")

}

// *** Methods Start ***//
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) UpdateAge(newAge int) {
	u.Age = newAge
	fmt.Println(u.Age)
}

func methods() {
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v.\n", hitesh.Name, hitesh.Email)

	hitesh.UpdateAge(60)
	fmt.Printf("Age is %v.\n", hitesh.Age)
}

//*** Methods End ***//

func myDefer() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
}

//*** Working with Files Start***//

func files() {
	content := "Thanks you very much from Sajib Bhai"

	file, err := os.Create("./testFile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println(length)
	defer file.Close()
	readFile("testFile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println(string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

//*** Working with Files End***//

// *** Json Start***//
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func encodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}

}

//*** Json End***//
