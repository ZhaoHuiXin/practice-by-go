package _array01

import (
	"errors"
	"fmt"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 *
 * Practice: zhaohuixin
 */

type Array struct{
 	data []int
 	length uint
 }

// initialize memory
func NewArray(capacity uint) *Array{
	if capacity == 0{
		return nil
	}
	return &Array{
		data: make([]int, capacity, capacity),
		length: 0,
	}
}

// return Array length
func (this *Array) Len() uint{
	return this.length
}

// Determines whether the index is out of range
func (this *Array) isIndexOutOfRange(index uint) bool{
	if index >= uint(cap(this.data)){
		return true
	}
	return false
}

// find element by index
func (this *Array) Find(index uint) (int, error){
	if this.isIndexOutOfRange(index){
		return 0, errors.New("out of index range")
	}
	return this.data[index], nil
}

// insert operation
func (this *Array) Insert(index uint, v int) error{
	if this.Len() == uint(cap(this.data)){
		return errors.New("full array")
	}
	// if only one element location left and it's the last one, the index can
	// be the arrary.Len(), and when add it, the length of array equal capacity.
	// so must determines index != this.Len() and whether it's out of range.
	if index != this.Len() && this.isIndexOutOfRange(index) {
		return errors.New("index out of range")
	}

	for i := this.Len(); i>index; i--{
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}

// insert v to tail
func (this *Array) InsertToTail(v int) error{
	return this.Insert(this.Len(), v)
}

// delete element by index
func (this *Array) Delete(index uint) (int, error){
	if this.isIndexOutOfRange(index){
		return 0, errors.New("index out of range")
	}
	v := this.data[index]
	for i := index; i < this.Len()-1; i++{
		this.data[i] = this.data[i+1]
	}
	this.length --
	return v, nil
}

// print array
func (this *Array) Print(){
	var format string
	for i := uint(0); i < this.Len(); i++{
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}