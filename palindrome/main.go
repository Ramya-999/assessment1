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
	fmt.Println("testing")
	s:="RADAR SIS RADAR"
	x:=Reverse(s)
	fmt.Println(s)
	fmt.Println(x)
    if x == s{
        fmt.Println("Palindrome")
    }else{
        fmt.Println("Not Palindrome")
    }
}
