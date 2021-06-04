package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

type Vertex struct {
	X, Y int
}

type Vertexx struct {
	Lat, Long float64
}

type Vertex1 struct {
	X, Y float64
}

var m1 = map[string]Vertexx{
	"Bell Labs": Vertexx{
		40.68433, -74.39967,
	},
	"Google": Vertexx{
		37.42202, -122.08408,
	},
}

type MyFloat float64

var m3 = map[string]Vertexx{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
var m map[string]Vertexx
var pow1 = []int{1, 2, 4, 8, 16, 32, 64, 128}
var c, python, java bool
var p, j int = 1, 2
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	v4 = &Vertex{1, 2} // has type *Vertex
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
	Pi    = 3.14
)

func main() {
	fmt.Println("Hello, 世界")

	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var p int
	fmt.Println(p, c, python, java)

	var c, python, java = true, false, "no!"
	fmt.Println(p, j, c, java, python)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var q int
	var f float64
	var v bool
	var s string
	fmt.Printf("%v %v %v %q\n", q, f, v, s)

	var x, y int = 3, 4
	var w float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(w)
	fmt.Println(x, y, z)

	v1 := 42 // change me!
	fmt.Printf("v is of type %T\n", v1)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	i1, j1 := 42, 2701

	o := &i1        // point to i
	fmt.Println(*o) // read i through the pointer
	*o = 21         // set i through the pointer
	fmt.Println(i1) // see the new value of i

	o = &j1        // point to j
	*o = *o / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(Vertex{1, 2})

	fmt.Println(v1, v4, v2, v3)

	var n [2]string
	n[0] = "Hello"
	n[1] = "World"
	fmt.Println(n[0], n[1])
	fmt.Println(n)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	primes1 := [6]int{2, 3, 5, 7, 11, 13}

	var s1 []int = primes1[1:4]
	fmt.Println(s1)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println(a1, b1)

	b1[0] = "XXX"
	fmt.Println(a1, b1)
	fmt.Println(names)

	q1 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q1)

	r2 := []bool{true, false, true, true, false, true}
	fmt.Println(r2)

	s2 := []struct {
		i2 int
		b2 bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)

	s3 := []int{2, 3, 5, 7, 11, 13}

	s3 = s3[1:4]
	fmt.Println(s3)

	s3 = s3[:2]
	fmt.Println(s3)

	s3 = s3[1:]
	fmt.Println(s3)

	s4 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s4)

	// Slice the slice to give it zero length.
	s4 = s4[:0]
	printSlice(s4)

	// Extend its length.
	s4 = s4[:4]
	printSlice(s4)

	// Drop its first two values.
	s4 = s4[2:]
	printSlice(s4)

	var s5 []int
	fmt.Println(s5, len(s5), cap(s5))
	if s5 == nil {
		fmt.Println("nil!")
	}

	a3 := make([]int, 5)
	printSlice1("a3", a3)

	b3 := make([]int, 0, 5)
	printSlice1("b3", b3)

	c3 := b3[:2]
	printSlice1("c3", c3)

	d3 := c3[2:5]
	printSlice1("d3", d3)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i3 := 0; i3 < len(board); i3++ {
		fmt.Printf("%s\n", strings.Join(board[i3], " "))
	}

	var s7 []int
	printSlice(s7)

	// append works on nil slices.
	s7 = append(s7, 0)
	printSlice(s7)

	// The slice grows as needed.
	s7 = append(s7, 1)
	printSlice(s7)

	// We can add more than one element at a time.
	s7 = append(s7, 2, 3, 4)
	printSlice(s7)

	for i4, v4 := range pow1 {
		fmt.Printf("2**%d = %d\n", i4, v4)
	}

	pow := make([]int, 10)
	for i4 := range pow {
		pow[i4] = 1 << uint(i4) // == 2**i4
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	m = make(map[string]Vertexx)
	m["Bell Labs"] = Vertexx{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(m1)

	fmt.Println(m3)

	m4 := make(map[string]int)

	m4["Answer"] = 42
	fmt.Println("The value:", m4["Answer"])

	m4["Answer"] = 48
	fmt.Println("The value:", m4["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m4["Answer"])

	v3, ok := m4["Answer"]
	fmt.Println("The value:", v3, "Present?", ok)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i5 := 0; i5 < 10; i5++ {
		fmt.Println(
			pos(i5),
			neg(-2*i5),
		)
	}

	v1 := Vertex1{3, 4}
	fmt.Println(v1.Abs())

	f1 := MyFloat(-math.Sqrt2)
	fmt.Println(f1.Abs())

	v1 := Vertex1{3, 4}
	v1.Scale(10)
	fmt.Println(v1.Abs())

	v1 := Vertex1{3, 4}
	Scale(&v1, 10)
	fmt.Println(Abs(v1))

	v1 := Vertex1{3, 4}
	v1.Scale(2)
	ScaleFunc(&v1, 10)

	p3 := &Vertex1{4, 3}
	p3.Scale(3)
	ScaleFunc(p3, 8)

	fmt.Println(v1, p3)

}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func printSlice(s4 []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s4), cap(s4), s4)
}

func printSlice1(s6 string, x6 []int) {
	fmt.Printf("%s6 len=%d cap=%d %v\n",
		s6, len(x6), cap(x6), x6)
}

func printSlice2(s7 []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s7), cap(s7), s7)
}

func WordCount(s8 string) map[string]int {
	return map[string]int{"x": 1}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f

}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
