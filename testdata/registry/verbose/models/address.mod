name: Address
fields:
  ID:
    type: AutoIncrement
    attributes:
      - mandatory
  Street:
    type: String
  HouseNumber:
    type: String
identifiers:
  primary: ID
  street:
    - Street
    - HouseNumber
related:
  Company:
    type: ForOne