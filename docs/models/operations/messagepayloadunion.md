# MessagePayloadUnion

The message payload for the "OrganizationInvitation" template to use when sending the invitation via email. If it is `false`, the invitation will not be sent via email.


## Supported Types

### MessagePayload

```go
messagePayloadUnion := operations.CreateMessagePayloadUnionMessagePayload(operations.MessagePayload{/* values here */})
```

### 

```go
messagePayloadUnion := operations.CreateMessagePayloadUnionBoolean(bool{/* values here */})
```

