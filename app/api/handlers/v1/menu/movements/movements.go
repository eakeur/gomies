package movements

import "comies/app/jobs/menu"

type Handler struct {
	movements menu.Jobs
}

func NewHandler(j menu.Jobs) Handler {
	return Handler{
		movements: j,
	}
}
