package main

import "fmt"

func multiply(x, y float64) float64 {
var sum float = 0

for i := b; i > 0; i=i-0.0001{
sum = sum + a
}
return sum/10000
}


//2:
func fibbonachi1(n int) int{
var f1, f2, fsum int = 0, 1, 0
if n == 1 {
fsum = f1
}
if n == 2{
fsum = f2
}
for i := 2; i < n : i++{
fsum = fq+f2
f1 = f2
f2 = fsum
}
return fsum
}


//3:
func fibbonachi(n uint) uint{

if n == 0 {
return 0
}
if n == 1 {
return 1
}
return fibbonachi(n - 1) + fibbonachi(n - 2)
}
func main() {

fmt.Println(fibbonachi(0))

}



//4:
func bubble_sort (a[] float64){
var x float64
for i:=0; i<len(a); i++{
for j:= len(a)-1; j>i; j--{
if a[j-1] > a[j]{
x = a[j-1]
a[j-1] = a[j]
a[j] = x
}
}
}
}


func bubble_sort2 (a[] int){
var x int
for i:=0; i<len(a); i++{
for j:= len(a)-1; j>i; j--{
if a[j-1] > a[j]{
x = a[j-1]
a[j-1] = a[j]
a[j] = x
}
}
}
}


//5:
func unique_count (a[] int) int {
var sum_uniq int=1
bubble_sort2(a)
for i := 0; i < len(a); i++ {
if a[i-1] !=a[i] {
sum_uniq++
}
}
return sum_uniq
}
