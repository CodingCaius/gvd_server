package set

// SetSub2 求两个集合的并集在分别对两个集合做差
func SetSub2[T any](set1, set2 []T) (delSet, addSet []T) {
	allSet := SetUnion(set1, set2)
	delSet = SetSub(allSet, set1)
	addSet = SetSub(allSet, set2)
	return
}