// Code generated by "stringer -type=day"; DO NOT EDIT.

package main

import "strconv"

const _day_name = "MondayTuesday"

var _day_index = [...]uint8{0, 6, 13}

func (i day) String() string {
	if i < 0 || i >= day(len(_day_index)-1) {
		return "day(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _day_name[_day_index[i]:_day_index[i+1]]
}
