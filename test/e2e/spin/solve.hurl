# GET endpoint shows short information about endpoint
GET http://localhost:3000/solve
HTTP 200

# POST endpoint return solution for happy case
POST http://localhost:3000/solve
{
  "target": 93,
  "digits": [5, 7, 9, 10, 15, 25]
}
HTTP 200

[Asserts]
header "Content-Type" contains "json"
jsonpath "$.value" == 93
jsonpath "$.operations" count == 4
jsonpath "$.operations[0].prev_value" == 0
jsonpath "$.operations[0].digit" == 5
jsonpath "$.operations[0].value" == 5

# POST endpoint returns error when no valid request is sent
POST http://localhost:3000/solve
HTTP 500

# POST endpoint returns error when empty struct was sent
POST http://localhost:3000/solve
{}

HTTP 500
[Asserts]
body contains "target must be between"

# POST endpoint returns error when target is not a number
POST http://localhost:3000/solve
{
  "target": "a",
  "digits": [5, 7, 9, 10, 15, 25]
}

HTTP 500
[Asserts]
body contains "parse error"
