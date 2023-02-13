package request

type AircraftCreateRequest struct {
	Name         string `json:"name" binding:"required"`
	CoverPicture string `json:"cover_picture"  binding:"required"`
	Tag          string `json:"tag"  binding:"required"`
	Type         string `json:"type" binding:"required"`
	Content      string `json:"content" binding:"required"`
}

type AircraftEditRequest struct {
	Id           uint   `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	CoverPicture string `json:"cover_picture"  binding:"required"`
	Tag          string `json:"tag"  binding:"required"`
	Type         string `json:"type" binding:"required"`
	Content      string `json:"content" binding:"required"`
}

type QueryPageAircraftR struct {
}
