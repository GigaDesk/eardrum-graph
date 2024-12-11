# eardrum-graph
This Go module provides a robust and efficient interface for interacting with Neo4j graph databases. It offers a comprehensive set of graph algorithms and leverages the powerful neo4j-go-driver to optimize query execution, asynchronous operations, and robust error handling to ensure high performance and reliability.   

## Key Functionalities 

- Initializing a connection to a neo4j database
- Creating a school node in a neo4j database
- Creating a student node with a `STUDENT_AT` relationship with a school node in a neo4j database
- Retrieving student nodes with relation to the school they have a `STUDENT_AT` relationship with.
- Retrieving school nodes with relation to the student they have a `STUDENT_AT` relationship with.    

## Technologies 

- Neo4j
- Golang

## Frameworks and Libraries

- neo4j-go-driver 

## Tests

Tests were carried out to ensure all the key functionalities above could be performed successfully. The testing data is in packages mockschool and mockstudent for school and student data respectively.