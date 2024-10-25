package e

const (
	SUCCESS = 200
	ERROR   = 400
)

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "something went wrong",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
