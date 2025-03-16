package solutions

import "fmt"

var charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\""


/*
 	link: https://www.codewars.com/kata/58a57c6bcebc069d7e0001fe/train/go
	neden ilk 13 bit 88den küçük mü diye kontrol ediyoruz?

	eğer ilk 13 bit 88den küçük bir değerse bu sayıyı tek bir karakterle de kodlayabiliriz
	ama 13 bit alırsak boşu boşuna ekstra bir karakter bu durumda 14 bit kullanıp sayıyı büyütmek ve 
	14 bit kullanmak daha mantıklı hale geliyor
*/
func Encode(d []byte) []byte {
	var count uint = 0
	var bits uint = 0
	var encoded []byte

	for _, byteVal := range d {
		bits |= uint(byteVal) << count 
		count += 8                       

		for count > 13 {
			extractedBits := bits & 8191 
			if extractedBits > 88 {
				bits >>= 13  
				count -= 13  
			} else {
				extractedBits = bits & 16383 
				bits >>= 14  
				count -= 14  
			}
			
			quotient := extractedBits / 91
			remainder := extractedBits % 91
			encoded = append(encoded, charset[remainder], charset[quotient]) 
		}
	}

	if count > 0 {
		remainder := bits % 91
		encoded = append(encoded, charset[remainder])
		if bits > 90 || count > 7 {
			quotient := bits / 91
			encoded = append(encoded, charset[quotient])
		}
	}
	
	fmt.Println(string(encoded))
	return encoded
}


func Decode(data []byte) []byte {
	var (
		bitPool uint = 0
		bitCount uint = 0
		decoded []byte
	)

	var value int = -1 

	for i := 0; i < len(data); i++ {
		index := indexOf(data[i]) 
		if index == -1 {
			continue 
		}

		if value == -1 {
			
			value = index
		} else {
			
			number := value + index*91 
			value = -1 

			var bitsToTake uint
			if (number & 8191) > 88 {
				bitsToTake = 13 
			} else {
				bitsToTake = 14 
			}

			
			bitPool |= uint(number) << bitCount
			bitCount += bitsToTake

			
			for bitCount >= 8 {
				decoded = append(decoded, byte(bitPool&255)) 
				bitPool >>= 8 
				bitCount -= 8
			}
		}
	}

	
	if value != -1 {
		bitPool |= uint(value) << bitCount
		bitCount += 7 

		
		if bitCount >= 8 {
			decoded = append(decoded, byte(bitPool&255))
		}
	}

	return decoded
}

func indexOf(char byte) int {
	for i := 0; i < len(charset); i++ {
		if charset[i] == char {
			return i
		}
	}
	return -1 
}

