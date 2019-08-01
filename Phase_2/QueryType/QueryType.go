package QueryType

type QueryType uint16

var (
	UNKNOWN QueryType = 0
	A       QueryType = 1
)

func QueryTypeToInt(QT QueryType) uint16 {
	if QT == A {
		return 1
	}
	return uint16(UNKNOWN)
}

func IntToQueryType(Int uint16) QueryType {
	if Int == 1 {
		return A
	}
	return UNKNOWN
}
