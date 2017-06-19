In order to compare the different languages we will build an API in each with the following:

 - GET `/add/x/to/y` which returns a JSON response `{ "sum": <x + y> }`
 - POST `/add` which takes a list of numbers in XML `<numbers><value>x</value>...</numbers>` and returns `<sum>123</sum>`
