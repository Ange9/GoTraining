package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}
type personP struct {
	first           string
	last            string
	iceCreamFlavors []string
}

func main() {
	//annonymous struct
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   37,
	}

	p2 := person{
		first: "Jenny",
		last:  "Penny",
		age:   37,
	}
	sa := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   38,
		},
		ltk: true,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(sa)
	fmt.Println(sa.first, sa.last, sa.age, sa.ltk)

	fmt.Println("-------Ice cream------")
	p3 := personP{
		first:           "Ange",
		last:            "Fallas",
		iceCreamFlavors: []string{"choco", "lime"},
	}

	//fmt.Println(p3)
	for i, v := range p3.iceCreamFlavors {
		fmt.Println(i, v)
	}

}
