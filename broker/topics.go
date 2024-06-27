package broker

const topicPrefix = "egapay."

type CollectionTopic string

const (
	PayPartnerRouterCollectionRequest CollectionTopic = topicPrefix + "pay_partner.router.collection_request"
	
	InitialMtnGhMadApiCollectionRequest  CollectionTopic = topicPrefix + "mtngh.mad_api.collection_request"
	InitialMtnGhMadApiCollectionResponse CollectionTopic = topicPrefix + "mtngh.mad_api.collection_response"
	
	InitialMtnGhOpenApiCollectionRequest  CollectionTopic = topicPrefix + "mtngh.open_api.collection_request"
	InitialMtnGhOpenApiCollectionResponse CollectionTopic = topicPrefix + "mtngh.open_api.collection_response"
)
