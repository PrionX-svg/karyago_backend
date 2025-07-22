package response

type CompanyDetailShiftWithShiftResponse struct {
	UUID      string `json:"uuid"`
	ShiftUUID string `json:"shift_uuid"`
	ShiftName string `json:"shift_name"`
	Day1      int    `json:"day_1"`
	Day2      int    `json:"day_2"`
	Day3      int    `json:"day_3"`
	Day4      int    `json:"day_4"`
	Day5      int    `json:"day_5"`
	Day6      int    `json:"day_6"`
	Day7      int    `json:"day_7"`
}
