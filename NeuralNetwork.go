package main

import (
	"fmt"
	"math"
)

func sigMoid (z float64) float64{
	return 1/ (1 + math.Exp(-z))
}

func X1(i int, n int) float64{
	var sum float64
	sum = math.Sin((2 * math.Pi * (float64(i) - 1)) / float64(n))
	return sum
}

func X2 (i int, n int) float64{
	var sum float64
	sum = math.Cos((2 * math.Pi * (float64(i) -1)) / float64(n) )
	return sum
}

func Z(w1 float64, w2 float64, x1 float64, a1 float64, x2 float64) float64{
	var v,z float64
	v = w1 + w2*x1 + a1*x2
	z = sigMoid(v)
	return z
}

func T (b1 float64, b2 float64, z1 float64, b3 float64, z2 float64, b4 float64, z3 float64) float64{
	var v,z float64
	v = b1 + b2*z1 + b3*z2 + b4*z3
	z = sigMoid(v)
	return z
}

func main (){
	num := sigMoid(2)
	fmt.Println(num)

	//Obtain input from user
	fmt.Print("Enter the n number: ")
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil{
		fmt.Print("error")
	}
	for i := 1; i <= input ; i++{
		//calculate X1 and X2
		x1 := X1(i ,input)
		x2 := X2(i, input)
		//declare weights and a
		var w10, w11, a12, w20, w21, a22,w30,w31,a32, b10,b11,b12,b13 float64
		w10, w11, a12, w20, w21, a22,w30,w31,a32, b10,b11,b12,b13 = 0.1, 0.3, 0.4,0.5,0.8,0.3,0.7,0.6,0.6,0.5,0.3,0.7,0.1

		var z1,z2,z3 float64
		z1 = Z(w10, w11, x1, a12 , x2)
		z2 = Z(w20, w21, x1, a22, x2)
		z3 = Z(w30, w31, x1, a32, x2)

		var t float64
		t = T(b10, b11, z1, b12, z2, b13, z3)
		fmt.Print("The result ")
		fmt.Print(i)
		fmt.Print("is: ")
		fmt.Println(t)
	}

}