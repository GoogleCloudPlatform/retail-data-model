# Retail

The retail module is the base extensible module for all other Go Modules.
The Retail module implements the life-cycle and configuration objects 
required to start and stop all module types.

Life-Cycle

```mermaid
flowchart TB
  i1[Start]
  i2[Load Module]
  i3[Initialize Logging]
  i4[Initialize BigQuery]
  i5[Run]
  i10[Load Loggers]
  i10a[Initialize Logging Client]
  i10b[Logging Ready]
  i11[Load Default]
  i12[Load BQ Logger]
  i13[Load PubSub Logger]
  i20[Load BigQuery]
  i21[Initialize BigQuery Client]
  i22[Initialize Module Tables]
  i30[Error Check]
  i1 --> i2
  i2 --> Logging
  Logging --> BigQuery
  BigQuery --> i30
  i30 --> i5
  subgraph Logging
    direction LR
    i3 --> i10
    i10 --> i10a
    i10a --> i10b
    i10b --> i11
    i10b --> i12
    i10b --> i13
  end
  subgraph BigQuery
   direction LR
   i4 --> i20
   i20 --> i21
   i21 --> i22
  end
```