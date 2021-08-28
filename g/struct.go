package g

type Model struct {
	Id        uint   `json:"id"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
}

type ReqId struct {
	Id uint `json:"id" param:"id"`
}

type Paginate struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	OrderBy  string `json:"orderBy"`
	Sort     string `json:"sort"`
}

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type ErrMsg struct {
	Code int64  `json:"code"`
	Msg  string `json:"message"`
}

var (
	NoErr         = ErrMsg{2000, "请求成功"}
	AuthErr       = ErrMsg{4001, "认证错误"}
	AuthExpireErr = ErrMsg{4002, "token 过期，请刷新token"}
	AuthActionErr = ErrMsg{4003, "权限错误"}
	ParamErr      = ErrMsg{4004, "参数解析失败，请联系管理员"}
	SystemErr     = ErrMsg{5000, "系统错误，请联系管理员"}
	DataEmptyErr  = ErrMsg{5001, "数据为空"}
	TokenCacheErr = ErrMsg{5002, "TOKEN CACHE 错误"}
)
