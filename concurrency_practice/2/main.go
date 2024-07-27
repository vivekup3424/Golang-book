package main
func captionElec(ninja chan string, message string){
	ninja <- message
}
func main(){
	ninja1 , ninja2 := make(chan string), make(chan string)
}