# Packages

Each package represents an individual layer of functionality found in most
modern applications.

### Package Descriptions

* Model - Extensions related to the protobuf definitions. Such as creators and utility methods.
* Data Access Layer (DAO) - The data access objects used by the service layer for persisting data.
* Service - The business logic layer.
* Module - Is the unifying unit responsible for the creation and life-cycle of services, DAOs, and the model.
* Server - Is the protocol listener and broker layer.
* Client - Is an embeddable client library for using the services over a protocol with little effort.
* Command - Is a wrapper for the server, allowing it to be natively constructed for running on operating systems.

## Dependencies

```mermaid
 flowchart LR
  i0[Server]
  i1[Module]
  i2[Service]
  i3[Data Access Layer]
  i3a[(Big Query)]
  i4[Model]
  i5[Client]
  i6[SDK]
  i7[Test]
  i8(Protocol Buffer)
  i9(GRPC)
  i10(Event Model)
  i11(Event Service)
  i0 --> i1
  i1 --> i2
  i2 --> i3
  i3 --> i4
  i3 --> i3a
  i2 --> i4
  i4 --> i6
  i5 --> i6
  i5 --> i4
  i7 --> i0
  i7 --> i5
  i6 --> i8
  i6 --> i9
  i8 --> i10
  i9 --> i11
  
```