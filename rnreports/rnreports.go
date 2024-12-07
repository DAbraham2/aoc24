package rnreports

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

func SafeDampener(report *[]int) bool {
	_inc := true
	_diff := diff((*report)[0], (*report)[1])
	_error_count := 0
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
			_error_count++
		}

		if _error_count > 1 {
			return false
		}
	}

	return true
}
