Simple Todo api rest service

# Stack
- Golang with chi as web server
- MongoDB
- Docker

# To run the all thing
'''
make run
'''

# Endpoint
'''
GET http://localhost:3000/tasks/
POST http://localhost:3000/tasks/
PUT http://localhost:3000/tasks/{id}
DELETE http://localhost:3000/tasks/{id}
'''

## Request body

'''
{
 "name": "string"
 "description": "string"
}
'''
