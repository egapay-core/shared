package broker

const topicPrefix = "egapay."

type CollectionTopic string

const (
	// InitialCollectionRequest - topic for initial collection request
	InitialCollectionRequest CollectionTopic = topicPrefix + "collection_request"
	
	// InitialCollectionResponse - topic for initial collection response
	InitialCollectionResponse CollectionTopic = topicPrefix + "collection_response"
)
