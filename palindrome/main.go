package main
import "fmt"

func Reverse (s string) string{
    var r string
    for i:=len(s)-1;i>=0;i--{
        r += string(s[i])
    }
   return r
}

func main(){
	
    fmt.Println("enter a string")
    var s string
    fmt.Scanln(&s)
	x:=Reverse(s)
	
    if x == s{
        fmt.Println("Palindrome")
    }else{
        fmt.Println("Not Palindrome")
    }
}
