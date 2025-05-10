package main
import ("fmt")
type sender struct {
	user
	rateLimit int
}

type user struct {
	name   string
	number int
}

func main() {
	myUser := sender {
		rateLimit: 10,
		user: user{
			name:   "Wunna",
			number: 7,
		},
	}

	fmt.Println(myUser.name)   
	fmt.Println(myUser.number) 
	fmt.Println(myUser.rateLimit)
}
