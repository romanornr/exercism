package main

import "fmt"

type Car struct {
    Wheels  int
    Brand   string
}

// met de * heb je een referentie naar het object, je past het direct aan
// zonder de * doe je niks met het object maar pas je de waarde alleen in de scope hieronder aan
func (car *Car) removeWheel() {
    car.Wheels -= 1
}


func main() {

    mycar := &Car{}

    mycar.Wheels = 6
    mycar.Brand = "lambo"

    mycar.removeWheel()
    mycar.removeWheel()


    fmt.Printf("%d", mycar.Wheels)



    //mycar := &Car{} //kan ook


}
