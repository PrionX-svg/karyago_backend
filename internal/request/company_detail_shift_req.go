package request

type CompanyDetailShiftRequest struct {
	ShiftUUID string `json:"shift_uuid" validate:"required"`
	Day1      int    `json:"day_1"`
	Day2      int    `json:"day_2"`
	Day3      int    `json:"day_3"`
	Day4      int    `json:"day_4"`
	Day5      int    `json:"day_5"`
	Day6      int    `json:"day_6"`
	Day7      int    `json:"day_7"`
}
