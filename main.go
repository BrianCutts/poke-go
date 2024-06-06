package main

import ("fmt")

func main() {
	var readIn string;
	var do_not_exit bool = true;
	for do_not_exit {
		fmt.Print("pokedex >");
		fmt.Scanln(&readIn);
		if readIn == "exit" {
			do_not_exit = false;
		}
		fmt.Print(readIn + "\n");
	}
}