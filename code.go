package main

// takes an asm dest and converts it to 3 bits in a string format

var dMap = map[string]string{
	"":    "000",
	"M":   "001",
	"D":   "010",
	"DM":  "011",
	"A":   "100",
	"AM":  "101",
	"AD":  "110",
	"ADM": "111",
}

func dest(in string) string {
	if dMap[in] != "" {
		return dMap[in]
	}
	return "" // add error handling
}

var cMap = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0001100",
	"!D":  "0001101",
	"!A":  "0110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"D+A": "0000010",
	"D-A": "0010011",
	"A-D": "0000111",
	"D&A": "0000000",
	"D|A": "0010101",
	"M":   "1110000",
	"!M":  "1110001",
	"-M":  "1110011",
	"M+1": "1110111",
	"M-1": "1110010",
	"D+M": "1000010",
	"D-M": "1010011",
	"M-D": "1000111",
	"D&M": "1000000",
	"D|M": "1010101",
}

// takes an asm dest and converts it to 3 bits in a string format
func comp(in string) string {
	if cMap[in] != "" {
		return cMap[in]
	}
	return "" // add error handling
}

var jMap = map[string]string{
	"null": "000",
	"":     "000",
	"JGT":  "001",
	"JEQ":  "010",
	"JGE":  "011",
	"JLT":  "100",
	"JNE":  "101",
	"JLE":  "110",
	"JMP":  "111",
}

// takes an asm dest and converts it to 3 bits in a string format
func jump(in string) string {
	if jMap[in] != "" {
		return jMap[in]
	}
	return "" // Handle errors
}
