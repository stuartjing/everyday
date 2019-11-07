package main

import "fmt"


func main(){
//	arr1 := [...]int {1,2,3,4,5,6,7}
//	var slice []int = arr1[0:7] 
	var arr []int = []int{1,2,3,4,5,6,7}
	rotate(arr,4)
}

func rotate(nums []int, k int)  {
    length := len(nums)
    
    var nums1 []int = make([]int,length) 
    var j int = 0
    var t int = 0
    
   if length != 1{
        if k>length{
            k = k-length
        }
        for i:=length-1;i>=0;i--{
            if (k -1) >=j {
                nums1[k-1-j]=nums[i]
            }else{
                nums1[length-1-t]=nums[i]
                t++
            }
            j++
        }
        for tt,v:=range nums1{
            fmt.Println(tt,v)
            nums[tt] = v
        }
    }
    fmt.Println(nums)
}

