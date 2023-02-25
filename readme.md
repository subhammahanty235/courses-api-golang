<h1>API Development in Golang</h1>

This is an API development code written in Golang. It provides endpoints for CRUD (Create, Read, Update, Delete) operations on the Course struct. It uses Gorilla Mux as a router for HTTP requests and responses.

<b>Course Struct </b>
The Course struct represents a course in the system. It contains the following fields:

CourseId (string): The unique identifier of the course.
CourseName (string): The name of the course.
CoursePrice (int): The price of the course.
Author (*Author): A pointer to an Author struct that contains the name and website of the author of the course.

<b>Author Struct</b>
The Author struct represents the author of a course. It contains the following fields:

Fullname (string): The name of the author.
Website (string): The website of the author.
Middleware
The IsEmpty() function is a middleware that checks whether the CourseName field of a Course struct is empty or not.

<b>Endpoints</b>
The API provides the following endpoints:

GET / - Returns a simple HTML response.
GET /getcourses - Returns a list of all courses in the system.
GET /course/{id} - Returns the details of a specific course.
POST /course - Creates a new course in the system.
PUT /updatecourse/{id} - Updates the details of a specific course in the system.
<b>Running the API</b>
To run the API, execute the following command:

go
Copy code
go run main.go
The API will be available at http://localhost:5000.

<b>Dependencies</b>
This code uses the following dependencies:

github.com/gorilla/mux: A request router and dispatcher.






