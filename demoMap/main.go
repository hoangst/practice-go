package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Person struct {
	id string
	name string
	age int
}



func main()  {
	//demoMap()
	m := WordCount("So Tuan Tuan Hoang Hoang Hoang1 Hoang1 27 08n1 0_1")
	fmt.Println(m)
}

func demoMap()  {
	var m map[string]Person
	m = make(map[string]Person)
	var person = Person{
		id:   "p01",
		name: "Person01",
		age:  10,
	}
	m["p01"] = person
	fmt.Println(m["p01"].name)

	m["p02"] = Person{
		id:   "p02",
		name: "Person02",
	}
	fmt.Println(m)

	m["p01"] = Person{
		id:   "p01",
		name: "Person03",
		age:  30,
	}
	fmt.Println(m)

	v, ok := m["p01"]
	fmt.Println("The value = ", v, "Present?", ok)

	delete(m, "p01")
	v, ok = m["p01"]
	fmt.Println("The value = ", v, "Present?", ok)
}

func WordCount(s string) map[string] int {
	mapWord := make(map[string]int)

	words := fieldFunc2(s)
	for _,word := range words{
		//fmt.Println("Word = ", word)
		v,ok := mapWord[word]
		if ok {
			v++
		} else {
			v = 1
		}
		mapWord[word] = v
	}

	return mapWord
}

//with condition
func fieldFunc(s string) []string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	return strings.FieldsFunc(s, f)
}

//simple
func fieldFunc2(s string) []string {
	return strings.Fields(s)
}