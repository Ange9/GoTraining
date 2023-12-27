package main

import "fmt"

func main() {

	m := make(map[string][]string)
	m[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m[`money_penny`] = []string{`intellegence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sun sets`}
	m[`flemming`] = []string{`steaks`, `cigars`, `espionage`}

	delete(m, `no_dr`)
	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}

}
