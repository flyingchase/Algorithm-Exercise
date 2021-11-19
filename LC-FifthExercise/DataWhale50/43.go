package datawhale50

func multiply(nums1 string, nums2 string) string {
	data := make([]byte, len(nums1)+len(nums2))
	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			num := (nums1[i]-'0')*(nums2[j]-'0') + data[i+j+1]
			// 进位的处理,从后往前
			data[i+j], data[i+j+1] = data[i+j]+num/10, num%10
		}
	}
	// 便于 strings 转化为对应的数字
	for i := 0; i < len(data); i++ {
		data[i] += '0'
	}
	for index, num := range data {
		// 去除前导 0
		// len(n1)+len(n2) 多出两位数，最多进位一位
		if num != '0' {
			return string(data[index:])
		}
	}
	return "0"
}
func multiply2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var res []byte
	temp, carry, size := 0, 0, len(num1)+len(num2)-2
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			temp = int(num2[j]-'0')*int(num1[i]-'0') + carry
			// 存储位置是否越界
			if len(res) > size-i-j {
				// 逆存放位置
				temp += int(res[size-i-j] - '0')
				carry = temp / 10
				res[size-i-j] = byte(temp%10 + int('0'))
			} else {
				carry = temp / 10
				res = append(res, byte(temp%10+int('0')))
			}
		}
		// 需要进位
		if carry != 0 {
			res = append(res, byte(carry+int('0')))
		}
		carry = 0
	}
	// 存储位置对调
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
