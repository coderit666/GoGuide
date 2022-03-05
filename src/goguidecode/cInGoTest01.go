package main

//#include <stdio.h>
//void say(){
// printf("Hello World\n");
//}
import "C"

func main() {
	C.say()
}
