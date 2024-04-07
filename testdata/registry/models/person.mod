name: Person
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
  FirstName:
    type: String
  LastName:
    type: String
identifiers:
  primary: ID
  entity: UUID
  name:
    - FirstName
    - LastName
related:
  Company:
    type: ForOne
  ContactInfo:
    type: HasOne