package apperror

import "net/http"

//定義每個 Error Code 的 HTTP 狀態與公開訊息
type Definition struct {
	HTTPStatus int
	Message    string
}

//統一管理錯誤代碼對應的資訊
var definitions = map[Code]Definition{
	// Common
	ErrInvalidRequest: {
		HTTPStatus: http.StatusBadRequest,
		Message:    "請確認輸入資料是否正確",
	},
	ErrInvalidID: {
		HTTPStatus: http.StatusBadRequest,
		Message:    "無效的待辦事項 ID",
	},
	ErrInternalServer: {
		HTTPStatus: http.StatusInternalServerError,
		Message:    "伺服器發生錯誤，請稍後再試",
	},

	// Todo
	ErrTodoNotFound: {
		HTTPStatus: http.StatusNotFound,
		Message:    "找不到指定的待辦事項",
	},
	ErrTodoInvalidDate: {
		HTTPStatus: http.StatusBadRequest,
		Message:    "日期格式錯誤，請使用 YYYY-MM-DD",
	},
	ErrTodoCreate: {
		HTTPStatus: http.StatusInternalServerError,
		Message:    "建立待辦事項失敗",
	},
	ErrTodoUpdate: {
		HTTPStatus: http.StatusInternalServerError,
		Message:    "更新待辦事項失敗",
	},
	ErrTodoDelete: {
		HTTPStatus: http.StatusInternalServerError,
		Message:    "刪除待辦事項失敗",
	},
	ErrTodoFinish: {
		HTTPStatus: http.StatusInternalServerError,
		Message:    "完成待辦事項失敗",
	},
	ErrTodoList: {
		HTTPStatus: http.StatusInternalServerError,
		Message:    "取得待辦事項清單失敗",
	},

	// Database
	ErrDatabaseUnavailable: {
		HTTPStatus: http.StatusServiceUnavailable,
		Message:    "資料庫暫時無法使用",
	},
}
