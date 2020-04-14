package main

import "fmt"

//Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop


func main() {
	//map[string]int
	x := map[string][]string{}
	fmt.Println(len(x))
	x["bond_james"] = []string{`"Shaken, not stirred"`, `"Martinis"`, `"Women"`}
	x["moneypenny_miss"] = []string{`"James Bond"`, `"Literature"`, `"Computer Science"`}
	x["no_dr"] = []string{`"Being evil"`, `"Ice cream"`, `"Sunsets"`}
	x["nobody"] = []string{`"nothing1"`,`"nothing2"`, `"nothing3"`}

	//m := map[string][]string{
	//	`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
	//	`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
	//	`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	//}

	for k, v := range x {
		fmt.Println("whoi = ", k)
		for k1, v1 := range v {
			fmt.Printf("\t %v = %v \n", k1, v1)
		}
	}
}


