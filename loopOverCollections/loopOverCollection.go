package main

func main() {
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for _, k := range wellKnownPorts {
		println(k)
	}
}
