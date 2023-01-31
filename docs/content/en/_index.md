---
title: "Google Retail Data Model"

resources:
  - name: placemat
    src: placemat.png
    title: "Retail Capabilities"
    params:
      credits: "[Ryan McGuinness](https://www.github.com/rrmcguinness)"
---

## Welcome

Google Retail Data Model is a collection of data and service definitions
describing the nouns and verbs related to the retail domain. It's broken down 
into several modules allowing a retailer or engineer to adopt as much
or as little as required.

> NOTE: Modules MAY NOT be independent, for example, Inventory requires
> Location and Merchandise.

## Contributing

Please see the [Contributing Guide]({{< relref "contributing" >}})


## Overview

{{< img name="placemat" size="large" lazy=false >}}

## Modules

* **Enums** - A set of common, standard, enumerated values used as adjectives to
  refine entity descriptions. Such as: country code, language code, etc.
* **Common** - A common set of definitions used by multiple modules. This includes
  constructs like Address, Phone, Country, etc.
* **Enterprise** - A set of definitions related to enterprise general ledger,
  assets, the human workforce, and hierarchies.
* **Location** - A set of definitions for describing physical, movable, and virtual
  locations. E.g. Site, Store, Warehouse, etc.
* **Party** - A set of definitions used for abstracting business and human
  relationships and understanding their change history over time. E.g.
  household, organization, membership, etc.
* **Customer** - A set of definitions used to describe a human or business customer
  and their journey to becoming a customer. E.g. anonymous customer,
  key customer, multi-channel customer etc.
* **Merchandise** - A set of definitions used to describe Product, Vendor, Pricing,
  Cost, Hierarchy, Hierarchy Pricing and Taxation Rules.
* **Inventory** - A set of definitions used to describe master inventory, assorted
  inventory, locational awareness, receiving, and stock.
* **Events** - A set of definitions used for describing events that MAY or MAY NOT
  support transactional flows. Events are used to enable performance indicators
  and represent near-real-time observability.
* **Transactions** - A set of definitions used to describe transactional events
  such as POS sales, Returns, Inventory Change Control, Cross Dock, etc.
  These definitions represent near-real-time observability and decision-making.
* **Supply Chain** - A set of definitions used for describing the distribution
  network used by a retailer
* **Payments** - A set of definitions used to describe payment account types such as
  gift cards, line of credit, etc. and the relationship to GL structures.

|              | Enums | Common | Enterprise | Location | Party | Customer | Merchandise | Inventory | Events | Transactions | Supply Chain | Payments |
|--------------|-------|--------|------------|----------|-------|----------|-------------|-----------|--------|--------------|--------------|----------|
| Enums        |       |        |            |          |       |          |             |           |        |              |              |          |   
| Common       | X     |        |            |          |       |          |             |           |        |              |              |          |
| Enterprise   | X     | X      |            | X        |       |          |             |           | X      | X            |              |          |
| Location     | X     | X      |            |          |       |          |             |           |        |              |              |          |
| Party        | X     | X      |            |          |       |          |             |           | X      | X            |              |          |
| Customer     | X     | X      | X          | X        | X     |          |             |           | X      | X            |              |          |
| Merchandise  | X     | X      | X          |          | X     |          |             |           | X      | X            |              |          |
| Inventory    | X     | X      | X          | X        |       |          | X           |           | X      | X            |              |          |
| Events       |       | X      |            |          |       |          |             |           |        |              |              |          |
| Transactions |       | X      |            |          | X     |          |             |           |        |              |              |          |
| Supply Chain |       |        |            | X        | X     |          | X           | X         | X      | X            |              |          |
| Payments     | X     | X      | X          |          | X     |          |             |           |        |              |              |          |

