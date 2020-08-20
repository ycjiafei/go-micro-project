package errcode

const (
	OK = iota
	Bad
	UserNotFound
)

var ErrMsg = map[int]string {
	OK: "请求成功",
	Bad: "请求失败",
	UserNotFound: "用户不存在",
}
