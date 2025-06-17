package main;
import "fmt";
import "os";
func main(){
 fd, err := os.Create("hello.txt");
 if err != nil {
   fmt.Print("errror");
 }
 defer fd.Close();

 fd.WriteString("hello");
 fmt.Print("hello");
}
