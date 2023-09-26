package model

type Petition struct {
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Image       string `json:"path_to_image"`
}

type PetitionID struct {
	PID uint32 `json:"pid"`
}

type PetitionStatus struct {
	Status string `json:"status"`
}
