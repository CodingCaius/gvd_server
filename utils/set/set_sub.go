package set

// SetSub 集合求差集 set1是大集合
func SetSub[T any](set1, set2 []T) (list []T) {
	var AnyMap = make(map[any]T)
	// 把小的放前面，这样大的才能在小的里面找不到
	for _, t := range set2 {
		AnyMap[t] = t
	}
	for _, t := range set1 {
		_, ok := AnyMap[t]
		if !ok {
			list = append(list, t)
		}
	}
	return
}