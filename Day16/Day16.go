package Day16

type packet struct{
	binaryRep string
	version int
	typeId int
	literalValue int
	operator int
	lengthType int
	lengthAmt int
	subPackets []*packet
}

func solvePt1(inputLines []string) {

	//for _,singleLine := range inputLines{
	//	thisPacket := createPacket(singleLine)
	//}

}

func createPacket(line string) packet {

	return packet{}
}

func solvePt2(inputLines []string) {

}

func Solve(inputLines []string) {
	solvePt1(inputLines)
	//solvePt2(inputLines)
}
