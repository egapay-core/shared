package broker

const topicPrefix = "egapay."

type CollectionTopic string

const (
	InitialMtnGhMadApiCollectionRequest  CollectionTopic = topicPrefix + "mtngh.mad_api.collection_request"
	InitialMtnGhMadApiCollectionResponse CollectionTopic = topicPrefix + "mtngh.mad_api.collection_response"

	InitialMtnGhOpenApiCollectionRequest  CollectionTopic = topicPrefix + "mtngh.open_api.collection_request"
	InitialMtnGhOpenApiCollectionResponse CollectionTopic = topicPrefix + "mtngh.open_api.collection_response"
)
