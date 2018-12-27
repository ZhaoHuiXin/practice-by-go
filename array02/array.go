package _array02

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
	return nil
}

// return Array length
func (this *Array) Len() uint{
	return 0
}

// Determines whether the index is out of range
func (this *Array) isIndexOutOfRange(index uint) bool{
	return false
}

// find element by index
func (this *Array) Find(index uint) (int, error){
	return 0, nil
}

// insert operation
func (this *Array) Insert(index uint, v int) error{
	return nil
}

// insert v to tail
func (this *Array) InsertToTail(v int) error{
	return nil
}

// delete element by index
func (this *Array) Delete(index uint) (int, error){
	return 0, nil
}

// print array
func (this *Array) Print(){

}