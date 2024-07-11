# Kafka Configuration for Core Services

This file contains the configuration for Kafka topics and consumer groups for the core services.

| Topic                                     | Consumer Groups            | Schema                      |
|-------------------------------------------|----------------------------|-----------------------------|
| egapay.mtngh.mad_api.collection_request   | * pay-partner-mtn-mad-api  | Message: Content of message |
|                                           |                            | Type: Protobuf              |
| egapay.mtngh.open_api.collection_request  | * pay-partner-mtn-open-api | Message: Content of message |
|                                           |                            | Type: Protobuf              |
| egapay.mtngh.mad_api.collection_response  | *                          | Message: Content of message |
|                                           |                            | Type: Protobuf              |
| egapay.mtngh.open_api.collection_response | Group C                    | Message: Content of message |
