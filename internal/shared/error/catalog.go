package error

import "net/http"

var (
    // --- 4xx Client Errors ---

    // 400 - Bad Request: Yêu cầu không hợp lệ (sai cú pháp, thiếu field)
    ErrBadRequest = New(http.StatusBadRequest, "common.client.bad_request")

    // 401 - Unauthorized: Chưa đăng nhập hoặc Token hết hạn
    ErrUnauthorized = New(http.StatusUnauthorized, "common.client.unauthorized")

    // 402 - Payment Required: Cần thanh toán để tiếp tục (ít dùng nhưng có sẵn)
    ErrPaymentRequired = New(http.StatusPaymentRequired, "common.client.payment_required")

    // 403 - Forbidden: Đã đăng nhập nhưng không có quyền truy cập tài nguyên này
    ErrForbidden = New(http.StatusForbidden, "common.client.forbidden")

    // 404 - Not Found: Tài nguyên không tồn tại
    ErrNotFound = New(http.StatusNotFound, "common.client.not_found")

    // 405 - Method Not Allowed: HTTP Method (GET/POST...) không được hỗ trợ tại route này
    ErrMethodNotAllowed = New(http.StatusMethodNotAllowed, "common.client.method_not_allowed")

    // 408 - Request Timeout: Request quá thời gian xử lý cho phép từ phía client
    ErrRequestTimeout = New(http.StatusRequestTimeout, "common.client.request_timeout")

    // 409 - Conflict: Xung đột dữ liệu (ví dụ: trùng Email khi đăng ký)
    ErrConflict = New(http.StatusConflict, "common.client.conflict")

    // 413 - Payload Too Large: Dữ liệu gửi lên quá lớn (ví dụ: file video quá dung lượng)
    ErrPayloadTooLarge = New(http.StatusRequestEntityTooLarge, "common.client.payload_too_large")

    // 422 - Unprocessable Entity: Dữ liệu đúng cú pháp nhưng sai nghiệp vụ (Validation lỗi)
    ErrUnprocessableEntity = New(http.StatusUnprocessableEntity, "common.client.validation_failed")

    // 429 - Too Many Requests: Client bị giới hạn vì gọi API quá nhiều (Rate limit)
    ErrTooManyRequests = New(http.StatusTooManyRequests, "common.client.too_many_requests")

    // --- 5xx Server Errors ---

    // 500 - Internal Server Error: Lỗi hệ thống không xác định
    ErrInternal = New(http.StatusInternalServerError, "common.server.internal_error")

    // 501 - Not Implemented: Tính năng chưa được phát triển
    ErrNotImplemented = New(http.StatusNotImplemented, "common.server.not_implemented")

    // 502 - Bad Gateway: Lỗi kết nối giữa các tầng server (Gateway lỗi)
    ErrBadGateway = New(http.StatusBadGateway, "common.server.bad_gateway")

    // 503 - Service Unavailable: Server đang quá tải hoặc bảo trì
    ErrServiceUnavailable = New(http.StatusServiceUnavailable, "common.server.service_unavailable")

    // 504 - Gateway Timeout: Server phía sau không phản hồi kịp cho Gateway
    ErrGatewayTimeout = New(http.StatusGatewayTimeout, "common.server.gateway_timeout")
)