package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"hris_backend/internal/models"
	"hris_backend/internal/services"
	"hris_backend/pkg"
)

type AttendanceEditHandler struct {
	svc services.AttendanceEditService
}

func NewAttendanceEditHandler(svc services.AttendanceEditService) *AttendanceEditHandler {
	return &AttendanceEditHandler{svc}
}

func (h *AttendanceEditHandler) Create(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return pkg.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	var req struct {
		WorkDate             string  `json:"work_date"`
		RequestType          string  `json:"request_type"`
		ProposedClockInAt    *string `json:"proposed_clock_in_at"`
		ProposedClockOutAt   *string `json:"proposed_clock_out_at"`
		ProposedIsHomeOffice *bool   `json:"proposed_is_home_office"`
		Reason               string  `json:"reason"`
	}
	if err := c.BodyParser(&req); err != nil || req.WorkDate == "" || req.RequestType == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "invalid payload")
	}

	wd, err := time.Parse("2006-01-02", req.WorkDate)
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, "invalid work_date")
	}

	var in, out *time.Time
	if req.ProposedClockInAt != nil && *req.ProposedClockInAt != "" {
		if t, e := time.Parse(time.RFC3339, *req.ProposedClockInAt); e == nil {
			in = &t
		}
	}
	if req.ProposedClockOutAt != nil && *req.ProposedClockOutAt != "" {
		if t, e := time.Parse(time.RFC3339, *req.ProposedClockOutAt); e == nil {
			out = &t
		}
	}

	id, err := h.svc.Create(userID, services.CreateEditReq{
		WorkDate:             wd,
		RequestType:          models.EditRequestType(req.RequestType),
		ProposedClockInAt:    in,
		ProposedClockOutAt:   out,
		ProposedIsHomeOffice: req.ProposedIsHomeOffice,
		Reason:               req.Reason,
	})
	if err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return pkg.Success(c, fiber.Map{"id": id}, "created")
}

func (h *AttendanceEditHandler) ListMine(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return pkg.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	var status *models.EditRequestStatus
	if v := c.Query("status"); v != "" {
		s := models.EditRequestStatus(v)
		status = &s
	}
	var from, to *time.Time
	if v := c.Query("from"); v != "" {
		t, _ := time.Parse("2006-01-02", v)
		from = &t
	}
	if v := c.Query("to"); v != "" {
		t, _ := time.Parse("2006-01-02", v)
		to = &t
	}

	rows, err := h.svc.ListMine(userID, status, from, to)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.Success(c, rows, "ok")
}

func (h *AttendanceEditHandler) ListForSupervisor(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return pkg.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	var status *models.EditRequestStatus
	if v := c.Query("status"); v != "" {
		s := models.EditRequestStatus(v)
		status = &s
	}
	var from, to *time.Time
	if v := c.Query("from"); v != "" {
		t, _ := time.Parse("2006-01-02", v)
		from = &t
	}
	if v := c.Query("to"); v != "" {
		t, _ := time.Parse("2006-01-02", v)
		to = &t
	}
	var q *string
	if v := c.Query("q"); v != "" {
		q = &v
	}
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	rows, total, err := h.svc.ListForSupervisor(userID, status, from, to, q, limit, offset)
	if err != nil {
		return pkg.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkg.Success(c, fiber.Map{"total": total, "items": rows}, "ok")
}

func (h *AttendanceEditHandler) Approve(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return pkg.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	idStr := c.Params("id")
	reqID, err := pkg.ParseUintE(idStr)
	if err != nil || reqID == 0 {
		return pkg.Error(c, fiber.StatusBadRequest, "invalid id")
	}

	var body struct {
		Note *string `json:"note"`
	}
	_ = c.BodyParser(&body)

	if err := h.svc.Approve(userID, reqID, body.Note); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return pkg.Success(c, nil, "approved")
}

func (h *AttendanceEditHandler) Reject(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok || userID == 0 {
		return pkg.Error(c, fiber.StatusUnauthorized, "unauthorized")
	}

	reqID := pkg.ParseUint(c.Params("id"))
	var body struct {
		Note string `json:"note"`
	}
	if err := c.BodyParser(&body); err != nil || body.Note == "" {
		return pkg.Error(c, fiber.StatusBadRequest, "note is required")
	}
	if err := h.svc.Reject(userID, reqID, body.Note); err != nil {
		return pkg.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return pkg.Success(c, nil, "rejected")
}
