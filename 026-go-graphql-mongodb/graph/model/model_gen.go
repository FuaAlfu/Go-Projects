package model

type(
	CreateJobListingInput struct{
		Title       string `json: "title"`
		Description string `json: "description"`
		Company     string `json: "company"`
		URL         string `json: "url`
	}

	DeleteJobResponse struct{
		DeleteJobID string `json: "deletedJobId"`
	}
)