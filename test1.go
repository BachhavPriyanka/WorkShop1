package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Education struct {
	Qualification string `json:"qualification"`
	Stream        string `json: "stream"`
	YearOfPassOut string `json:"Year Of PassOut"`
}
type Student struct {
	Name                string `json:"name"`
	EmailID             string `json:"emailID"`
	Education           Education
	LocationPreferences []string `json:"location Preferences"`
}

func main() {

	//Q1) Read the LearnerData.json file and print the data.
	content, err := ioutil.ReadFile("LearnerData.json")
	if err != nil {
		panic(err)
	}

	var students []Student

	err2 := json.Unmarshal(content, &students)
	if err2 != nil {
		panic(err2)
	}

	for _, x := range students {
		fmt.Printf("Name : %s | Email ID : %s | Qualification : %s | Stream : %s | Year of passout : %s \n", x.Name, x.EmailID, x.Education.Qualification, x.Education.Stream, x.Education.YearOfPassOut)
		fmt.Println(x.LocationPreferences)
	}
	//Q3) From the above data in json file create a map of name and year of passout.
	mapNameYear := make(map[string]string)
	for _, value := range students {
		mapNameYear[value.Name] = value.Education.YearOfPassOut
	}

	fmt.Println(mapNameYear)

	//Q4)From the above data, print the names and year of passout for persons who passed out 2020 or late

	for _, checkYOP := range students {
		storeYOP, err := strconv.Atoi(checkYOP.Education.YearOfPassOut)
		if err != nil {
			panic(err)
		}
		if storeYOP >= 2020 {
			fmt.Printf("Name :%s | Year Of Passout : %s\n", checkYOP.Name, checkYOP.Education.YearOfPassOut)
		}
	}

	//Q5)Using the above data, please print out the data (name, location preferences) who have location preferences as West.
	for _, strCompare := range students {
		for _, location := range strCompare.LocationPreferences {
			if location == "West" {
				fmt.Printf("Name : %s | Location : %s", strCompare.Name, strCompare.LocationPreferences)
			}
		}
	}
}
