package main

import "fmt"

func main()  {
	/*a :=[]int{3,2,4}
	for _,v:=range a{
		fmt.Printf("%d\t",v)
	}
	fs:=twoSum(a,6)
	fmt.Println(fs)*/
    var n int32
    n=1<<31-1
	fmt.Println(n)
}

func twoSum(nums []int, target int) []int {
	ids:=make([]int,0)
	nums2:=nums
	floag:=false
	for i,v1:=range nums{
		for j,v2:=range nums2{
			if (v1+v2)==target&&i!=j{
				ids=append(ids,i)
				ids=append(ids,j)
				floag=true
				break
			}
		}
		if floag{
			break
		}
	}
	return ids
}
