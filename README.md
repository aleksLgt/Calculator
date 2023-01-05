# Calculator
This project is a calculator, which can used for adding, subtracting, dividing and multiplying numbers using API. The backend is developed on Go.

List of APIs:
* GET /api/calculator/add - add two numbers
* GET /api/calculator/subtract - subtract two numbers
* GET /api/calculator/divide - divide two numbers
* GET /api/calculator/multiply - multiply two numbers

Params:
* firstValue - the first value
* secondValue - the second value

How to use:
1. Download and install git https://git-scm.com/ , golang https://go.dev/doc/install
2. Clone repo on your computer git clone https://github.com/aleksLgt/Calculator.git
3. Go to the project folder
4. Run "go run main.go"
5. If the program launch was successful, the application is available at "http://localhost:8080/"

Applied knowledge:
* Use Gin Web Framework for developing APIs
* Splitting the code into packages
* Using structures
* Using encapsulation at the expense of getters and setters
* Implementation of calculator logic
* Using external packages
* Error Handling (e.g. Cannot be divided by zero)
* Autotests
