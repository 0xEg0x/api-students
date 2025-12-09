# api-students
API to manage course students Elohim

Routes:
- GET /students - List all students
- POST /students - create student
- GET /students/:id - consult unic student
- PUT /students/:id - Update student
- DELETE /students/:id - Delete student
- GET /students?active=<true/false> - List all active/non-active students

Struct Student:
- Name
- CPF
- E-mail
- Age
- Active