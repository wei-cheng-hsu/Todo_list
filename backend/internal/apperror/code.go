package apperror

type Code string

const (
	//Common error
	ErrInvalidRequest Code = "COMMON_001"
	ErrInvalidID      Code = "COMMON_002"
	ErrInternalServer Code = "COMMON_500"

	//Todo error
	ErrTodoNotFound    Code = "TODO_001"
	ErrTodoInvalidDate Code = "TODO_002"
	ErrTodoCreate      Code = "TODO_003"
	ErrTodoUpdate      Code = "TODO_004"
	ErrTodoDelete      Code = "TODO_005"
	ErrTodoFinish      Code = "TODO_006"
	ErrTodoList        Code = "TODO_007"

	//Database error
	ErrDatabaseUnavailable Code = "DB_001"
)
