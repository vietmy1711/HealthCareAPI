package banana

import "errors"
var (
	UserConflict   = errors.New("Người dùng đã tồn tại")
	UserNotFound   = errors.New("Không tìm thấy người dùng")
	UserNotUpdated = errors.New("Cập nhật thông tin người dùng thất bại"))
