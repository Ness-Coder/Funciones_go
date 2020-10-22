package main
 
import "fmt"
 
//fibonacci recursivo 
func fiboR (n uint) uint {
	if n < 2 {
	 return n
	}
	return fiboR (n-2) + fiboR (n-1)
   }
/*func fib(n uint) uint {
    if n == 0 {
        return 0
    } else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}*/

func sum(args... int) int{
	numeroMayor := args[0] //Asignar la variable el primer numero
	for _, v := range args { 
		if v > numeroMayor{
			numeroMayor = v
		}
		
	}

	return numeroMayor
}

func generadorPares() func() uint {
	i := uint(1)
	return func() uint {
		var impar = i
		if i%2 != 0{
			i += 2
		}
		
		return impar
	}
}

func intercambio(a *int, b *int){
	//SE aplica el cambio debido a cada uno de los enteros 
	numeroACambiar := a
 	a = b
	b = numeroACambiar
	fmt.Println("Primer numero: ",*a,", segundo Numero:",*b)
}

func main() {
	n := uint(5)
	fmt.Println("Fibonacci recursivo de: ",n)
	fmt.Println("Es igual a = ",fiboR(n))
	fmt.Println()
	
	fmt.Println("La suma total es ")
	a := []int{18, 28, 21, 96, 97}
	fmt.Println("El numero mas grande es: ", sum(a...))
	fmt.Println()
	
	fmt.Println("Numero impares")
	nextImpar := generadorPares()
	fmt.Println(nextImpar())
	fmt.Println(nextImpar())
	fmt.Println(nextImpar())
	fmt.Println(nextImpar())
	fmt.Println()

	var primerNumero int  = 64
	var segundoNumero int = 32
	fmt.Println("Cambio de numeros,  Primer numero: ",primerNumero,"Segundo Numero: ",segundoNumero)
	intercambio(&primerNumero , &segundoNumero)
	
}