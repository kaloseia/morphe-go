name: Company
fields:
  UUID:
    type: UUID
    attributes:
      - immutable
      - mandatory
  ID:
    type: AutoIncrement
    attributes:
      - mandatory
  FoundedAt:
    type: Time
  Name:
    type: String
identifiers:
  primary: ID
  entity: UUID
related:
  Address:
    type: HasOne
  Person:
    type: HasMany