package set

// SetUnion 集合求并集
func SetUnion[T any](set1, set2 []T) (list []T) {
	//var AnyMap = make(map[any]T)
	//for _, t := range set1 {
	//	AnyMap[t] = t
	//}
	//for _, t := range set2 {
	//	AnyMap[t] = t
	//}
	//for _, t := range AnyMap {
	//	list = append(list, t)
	//}
	var AnyMap = make(map[any]T)
	for _, t := range set1 {
		_, ok := AnyMap[t]
		if ok {
			continue
		}
		AnyMap[t] = t
		list = append(list, t)
	}
	for _, t := range set2 {
		_, ok := AnyMap[t]
		if ok {
			continue
		}
		AnyMap[t] = t
		list = append(list, t)
	}
	return
}