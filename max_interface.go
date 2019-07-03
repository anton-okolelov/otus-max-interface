package max_interface

func getMax(customSlice []interface{}, comparator func(value1 interface{}, value2 interface{}) bool) interface{} {
	var max interface{}

	for _, v := range customSlice {
		if max != nil {
			if comparator(v, max) {
				max = v
			}
		} else {
			max = v
		}
	}
	return max
}
