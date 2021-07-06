// CUBIC ROOT SOLVER
//
// Date Created   :    24.05.2017
// Created by     :    Shril Kumar [(shril.iitdhn@gmail.com),(github.com/shril)] &
//                     Devojoyti Halder [(devjyoti.itachi@gmail.com),(github.com/devojoyti)]
//
// Project        :    Classified
// Use Case       :    Instead of using standard numpy.roots() method for finding roots,
//                     we have implemented our own algorithm which is ~10x faster than
//                     in-built method.
//
// Algorithm Link :    www.1728.org/cubic2.htm
//
// This script (Cubic Equation Solver) is an independent program for computation of roots of Cubic Polynomials. This script, however,
// has no relation with original project code or calculations. It is to be also made clear that no knowledge of it's original project
// is included or used to device this script. This script is complete freeware developed by above signed users, and may further be
// used or modified for commercial or non-commercial purpose.
// Main Function takes in the coefficient of the Cubic Polynomial
// as parameters and it returns the roots in form of numpy array.
// Polynomial Structure -> ax^3 + bx^2 + cx + d = 0

package gocubicsolver

import (
	"errors"
	"fmt"
	"math"
)

func Solve(a, b, c, d float64) ([]float64, error) {
	roots := []float64{}
	if a+b == 0 { // Case for handling Liner Equation
		//fmt.Println("Linear Case")
		roots = append(roots, (-d*1.0)/c)
		return roots, nil
	}

	if a == 0 { // quadratic case
		fmt.Println("Cuadratic Case")
		D := c*c - 4*b*d
		if D >= 0 {
			fmt.Println("Cuadratic Case, D>0")
			D = math.Pow(D, 0.5)
			roots = append(roots, (-c+D)/(2.0*b))
			roots = append(roots, (-c-D)/(2.0*b))
			return roots, nil

		}

		if D < 0 {
			fmt.Println("Cuadratic Case, D<0")
			// TODO implement
			//(-c+D)/(2.0*b)
			//
			//            D = math.sqrt(-D)
			//            x1 = (-c + D * 1j) / (2.0 * b)
			//            x2 = (-c - D * 1j) / (2.0 * b)
			return roots, nil
		}
	}

	f := findF(a, b, c)    //  Helper Temporary Variable
	g := findG(a, b, c, d) // Helper Temporary Variable
	h := findH(g, f)       //  Helper Temporary Variable
	//fmt.Println("Cubic f,g,j", f, g, h)
	if f == 0 && g == 0 && h == 0 { // All 3 roots are real equals

		//fmt.Println("Cubic Case 3 equal roots,d,a", d, a)
		//fmt.Println("Cubic Case 3 equal roots")
		if (d / a) >= 0 {
			r := math.Pow(d/(1.0*a), 1.0/3.0) * -1 // TODO Check -1 fi multy the exp
			//fmt.Println("Cubic Case 3 equal roots,r", r)
			roots = append(roots, r, r, r)
			return roots, nil
		}
		if (d / a) < 0 {
			r := math.Pow(-d/(1.0*a), 1.0/3.0)
			//fmt.Println("Cubic Case 3 equal roots,r", r)
			roots = append(roots, r, r, r)
			return roots, nil
		}
	}
	if h <= 0 { // All 3 Roots are real but differents
		// Helpers
		//fmt.Println("Cubic Case 3 Real Roots") //
		i := math.Sqrt((g*g/4 - h))
		j := math.Pow(i, 1/3.0)
		k := math.Acos(-g / (2 * i))
		L := j * -1
		M := math.Cos(k / 3.0)
		N := math.Sqrt(3) * math.Sin(k/3.0)
		P := (b / (3.0 * a)) * -1
		roots = append(roots, 2*j*math.Cos(k/3.0)-(b/(3.0*a)))
		roots = append(roots, L*(M+N)+P)
		roots = append(roots, L*(M-N)+P)
		return roots, nil
	}

	if h > 0 { // One real Root and 2 complex, TODO implement complex
		R := -(g / 2.0) + math.Sqrt(h)
		var S float64
		var U float64
		if R >= 0 {
			S = math.Pow(R, 1/3.0)
		}
		if R < 0 {
			S = math.Pow(-R, 1/3.0) * -1
		}

		T := -(g / 2.0) - math.Sqrt(h)
		if T >= 0 {
			U = math.Pow(T, 1/3.0)
		}
		if T < 0 {

			U = math.Pow(-T, 1/3.0) * -1
		}
		fmt.Println("U, S", U, S)

		roots = append(roots, (S+U)-(b/(3.0*a)))
		// TODO Handling complex number
		//        x2 = -(S + U) / 2 - (b / (3.0 * a)) + (S - U) * math.sqrt(3) * 0.5j
		//        x3 = -(S + U) / 2 - (b / (3.0 * a)) - (S - U) * math.sqrt(3) * 0.5j

		return roots, nil
	}
	return roots, errors.New(fmt.Sprintf("Error no if entered pls check"))
}

//# Helper function to return float value of f.
func findF(a, b, c float64) float64 {
	return ((3.0 * c / a) - ((b * b) / (a * a))) / 3.0
}
func findG(a, b, c, d float64) float64 {
	return (((2.0 * (b * b * b)) / (a * a * a)) - ((9.0 * b * c) / (a * a)) + (27.0 * d / a)) / 27.0
}
func findH(g, f float64) float64 {
	return ((g*g)/4.0 + (f*f*f)/27.0)
}
