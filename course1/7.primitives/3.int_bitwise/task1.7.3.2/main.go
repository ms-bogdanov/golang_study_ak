package main

import "fmt"

var permissions = map[int]string{
	0: "-,-,-",
	1: "-,-,Execute",
	2: "-,Write,-",
	3: "-,Write,Execute",
	4: "Read,-,-",
	5: "Read,-,Execute",
	6: "Read,Write,-",
	7: "Read,Write,Execute",
}

func main() {
	fmt.Println(getFilePermissions(777))
}

func getFilePermissions(flag int) string {
	return fmt.Sprintf("Owner:%s Group:%s Other:%s", permissions[flag/100], permissions[flag%100/10], permissions[flag%10])
}
