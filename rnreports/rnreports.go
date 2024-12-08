package rnreports

import "errors"

func diff(a, b int) int {
	return a - b
}

func valid_diff(d int) bool {
	if d < 0 {
		d = d * -1
	}
	if d <= 3 && d >= 1 {
		return true
	}
	return false
}

func is_safe(report *[]int) (bool, []uint) {
	_valid := true
	_inc := true
	_diff := diff((*report)[1], (*report)[0])
	if _diff < 0 {
		_inc = false
	}
	_error_location := make([]uint, 0, len(*report))
	if !valid_diff(_diff) {
		_error_location = append(_error_location, 0)
		_error_location = append(_error_location, 1)
		_valid = false
	}

	for i := 2; i < len(*report); i++ {
		_diff = diff((*report)[i], (*report)[i-1])
		__inc := (0 < _diff)
		if i == 2 && _inc != __inc {
			_error_location = append(_error_location, 0)
			_error_location = append(_error_location, 1)
			_valid = false
			_inc = __inc
		}
		if !(_inc == __inc) || !valid_diff(_diff) {
			_error_location = append(_error_location, uint(i-1))
			_error_location = append(_error_location, uint(i))
			_valid = false
		}
	}

	if !_valid {
		return _valid, _error_location
	}

	return _valid, nil
}

func Safe(report *[]int) bool {
	_inc := true
	_diff := diff((*report)[0], (*report)[1])
	if _diff < 0 {
		_inc = false
	}

	if !valid_diff(_diff) {
		return false
	}

	for i := 2; i < len(*report); i++ {
		_diff = diff((*report)[i-1], (*report)[i])
		if !(_inc == (0 < _diff)) {
			return false
		}

		if !valid_diff(_diff) {
			return false
		}
	}

	return true
}

func remove(slice []int, idx uint) ([]int, error) {
	if idx > uint(len(slice)) {
		return nil, errors.New("invalid operation")
	}
	ret := make([]int, 0, len(slice)-1)
	ret = append(ret, slice[:idx]...)
	ret = append(ret, slice[idx+1:]...)
	return ret, nil
}

func SafeDampener(report *[]int) bool {
	safe, err := is_safe(report)
	if safe {
		return true
	}

	for _, v := range err {
		corrected, _ := remove(*report, v)
		safe, _ := is_safe(&corrected)
		if safe {
			return true
		}
	}

	return false
}
