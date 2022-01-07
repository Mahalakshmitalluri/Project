package calculator
import "fmt"

var a int8 =32
var b int8 = 76

var c int16 = 6700
var d int16 = 7890

var e int32 = 9086436
var f int32 = 7809542

var g int64 = 980754328
var h int64 = -78905432 


var  i uint = 987450
var  j uint = 9085564

var k uint16 = 7896
var l uint16 = 7863

var m uint32 = 897665
var n uint32 = 789533

var o uint64 = 8975325678900
var p uint64 = 9862167906578

var q float32 =780.9854
var r float32 =9034.98

var s float64 = 899654268.80
var t float64 = 67803.78

func Addint8(){
	fmt.Println(a+b)
}

func Addint16(){
	 fmt.Println(c+d)
}

func Addint32(){
	fmt.Println(e+f)
}


func Addint64(){
	fmt.Println(g+h)
}
func Adduint8(){
	fmt.Println(i+j)
}

func Adduint16(){
	fmt.Println(k+l)
}

func Adduint32(){
	fmt.Println(m+n)
}

func Adduint64(){
	fmt.Println(o+p)
}

func Addfloat32(){
	fmt.Println(q+r)
}


func Addfloat64(){
	fmt.Println(s+t)
}