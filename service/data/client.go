package data

type ConferenceStorageModel struct {
	UniqueName  string `bson:"uniqueName"`
	DisplayName string `bson:"displayName"`
	Description string `bson:"description"`
}

type DatabaseClient interface {
	GetConference(uniqueName string) (*ConferenceStorageModel, error)
	GetConferences() ([]*ConferenceStorageModel, error)
	UpdateConference(conference *ConferenceStorageModel) error
	CreateConference(conference *ConferenceStorageModel) error
}
