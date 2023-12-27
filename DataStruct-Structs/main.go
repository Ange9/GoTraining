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
type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
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

	p4 := personP{
		first:           "Rox",
		last:            "Qeuesa",
		iceCreamFlavors: []string{"Mint", "Vanilla"},
	}

	//fmt.Println(p3)
	for i, v := range p3.iceCreamFlavors {
		fmt.Println(i, v)
	}

	fmt.Println("-------Map of person------")
	m := map[string]personP{
		p3.last: p3,
		p4.last: p4,
	}

	for k, v := range m {
		for _, f := range v.iceCreamFlavors {
			fmt.Println(k, f)
		}
	}

	fmt.Println("-------Embed struct------")
	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Toyota",
		model: "Rav4",
		color: "Blue",
		doors: 4,
	}

	fmt.Println(v1.electric, v1.make, v1.model, v1.color, v1.doors)

	fmt.Println("-------Anonym struct------")
	f := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "Vero",
		friends:   map[string]int{"Raque": 8, "Chela": 9},
		favDrinks: []string{"water", "Coffee"},
	}

	//fmt.Println(f)

	for k, v := range f.friends {
		fmt.Println(k, v)
	}
	for k, v := range f.favDrinks {
		fmt.Println(k, v)
	}

}
