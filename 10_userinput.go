package main

import (

	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a)
}

/*


there are various other methods to take input from the user/console and read it by using 

in  = bufio.NewScanner(os.Stdin)
out = bufio.NewWriter(os.Stdout)


reader := bufio.NewReader(os.Stdin)


*/