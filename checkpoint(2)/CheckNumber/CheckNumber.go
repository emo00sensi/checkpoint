package piscine

func CheckNumber(arg string) bool {
	for i:= 0 ; i < len (arg) ;i++{
		if arg[i]>= '0' && arg[i]<= '9'{
			return false
		}
	}
	return true
}