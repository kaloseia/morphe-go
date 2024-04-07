name: ContactInfo
fields:
  ID:
    type: AutoIncrement
    attributes:
      - mandatory
  Email:
    type: String
    attributes:
      - mandatory
  PhoneNumber:
    type: String
identifiers:
  primary: ID
related:
  Person:
    type: ForOne