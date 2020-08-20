package errcode

const (
	OK = iota
	Bad
	MissArgument
	NewServiceFail
	UserNotFound
)

var ErrMsg = map[int]string {
	OK: "请求成功",
	Bad: "请求失败",
	MissArgument: "缺少必要参数",
	NewServiceFail: "初始化服务失败",
	UserNotFound: "用户不存在",
}
