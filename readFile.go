package main
import (
    "fmt"
    "os"
    "bufio"
    //"io/ioutil"
    //"log"
)/*
func main() {

    words, err := ioutil.ReadFile("smallDictionary.txt")

     if err != nil {
          log.Fatal(err)
     }

    fmt.Println(string(words))
}
*/

func main() {

    //The Open function opens the file for reading.
    word, err := os.Open("smallDictionary.txt")
    //if error, prints the error
    if err != nil {
        fmt.Println(err)
     }
     //The file descriptor is closed at the end of the main function.
    defer word.Close()
    // A new scanner is created.
    scanner := bufio.NewScanner(word)
    //split the content by words.
    scanner.Split(bufio.ScanWords)
    //The Scan advances the Scanner to the next token, 
    //which will then be available through the Text function.
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    //Checks for Error in scanning
	scanErr := scanner.Err()
    if err != nil {
        fmt.Println(scanErr)
    }
}
