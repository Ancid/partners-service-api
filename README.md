# Golang API service coding exercise
This is a coding exercise done by Anton Sidorov

### Task description
Your task is to write an API that offers the following functionality:
- Based on a customer-request, return a list of partners that offer the service. The list
should be sorted by “best match”. The quality of the match is determined first on
average rating and second by distance to the customer.
- For a specific partner, return the detailed partner data.

Matching a customer and partner should happen on the following criteria:
- The partner should be experienced with the materials the customer requests for the
  project.
- The customer is within the operating radius of the partner.

The data in the request from the customer is:
- Material for the floor (wood, carpet, tiles)
- Address (assume that this is the lat/long of the house)
- Square meters of the floor
- Phone number (for the partner to contact the customer)

The structure of the partner data is as follows:
- Experienced in flooring materials (wood, carpet, tiles or any combination)
- Address (assume that this is the lat/long of the office)
- Operating radius (consider the beeline from the address)
- Rating (for this assignment you can assume that you already have a rating for a
partner)

### The Partners service
The service provides several APIs for testing:

**Search partners endpoint**

Endpoint for searching partners with material filter and distance check. Results are ordered by rating and distance. Partnes who are out of radius are not shown.

example: http://localhost:8080/partners?latitude=49.671072&longitude=8.850669&material=wood

**Partner details endpoint**

Endpoint for detailed partner info

example: http://localhost:8080/partners/1

**All partners endpoint**

Endpoint for checking all the list of partners in database

example: http://localhost:8080/partners/list 

### Installing

Run docker and init project once:
```
 make build
```

run tests:
```
 make tests
```

for start and stop the project use:
```
 make up
 make down
```

to see service logs use:
```
 make logs
```

### Dev notes
What can be improved?
- tests. Lack of test in this project relates to not the best structure of api handler. 
  In current structure it is pretty hard to mock services
- caching can be added
- use more common framework for server implementation
- API Swagger documentation can be added
