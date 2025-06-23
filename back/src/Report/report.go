package report

import "time"

type ReportCreateDTO struct {
	DeliveryID    uint   `json:"delivery_id" binding:"required"`
	Description   string `json:"description" binding:"required"`
	TypeID        uint   `json:"type_id" binding:"required"`
	ReportStateID uint   `json:"report_state_id" binding:"required"`
}

type ReportTypeInResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ReportStateInResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type DeliveryInReportResponseDTO struct {
	ID      uint `json:"id"`
	OrderID uint `json:"order_id"`
}

type ReportResponseDTO struct {
	ID          uint                        `json:"id"`
	Description string                      `json:"description"`
	Date        time.Time                   `json:"date"`
	Delivery    DeliveryInReportResponseDTO `json:"delivery"`
	ReportType  ReportTypeInResponseDTO     `json:"report_type"`
	ReportState ReportStateInResponseDTO    `json:"report_state"`
}
