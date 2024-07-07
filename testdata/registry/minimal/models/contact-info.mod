name: ContactInfo
fields:
  ID:
    type: AutoIncrement
    attributes:
      - mandatory
  Email:
    type: String
identifiers:
  primary: ID
  email: Email
related:
  Person:
    type: ForOne