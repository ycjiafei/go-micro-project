package errcode

const (
	OK = iota
	Bad
	MissArgument
	UserNotFound
)

var ErrMsg = map[int]string {
	OK: "请求成功",
	Bad: "请求失败",
	MissArgument: "缺少必要参数",
	UserNotFound: "用户不存在",
}
