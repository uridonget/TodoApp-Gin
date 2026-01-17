# TodoApp-Gin

A simple Todo application built with Gin and GORM.

## Setup

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/TodoApp-Gin.git
    cd TodoApp-Gin
    ```

2.  **Create a `.env` file:**
    Create a file named `.env` in the root directory of the project and add the following content:

    ```
    DB_HOST=your_db_host
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
    DB_PORT=your_db_port
    PORT=8080
    ```
    Replace `your_db_host`, `your_db_user`, `your_db_password`, `your_db_name`, and `your_db_port` with your PostgreSQL database credentials.

3.  **Run the application:**
    ```bash
    go run main.go
    ```

## Endpoints

-   **GET /ping**
    Health check.

-   **GET /todos**
    Get all todo items.

-   **POST /todos**
    Create a new todo item.
    Request body:
    ```json
    {
        "title": "My new todo",
        "is_completed": false
    }
    ```

-   **PATCH /todos/:id**
    Update an existing todo item.
    Request body (fields are optional):
    ```json
    {
        "title": "Updated title",
        "is_completed": true
    }
    ```

-   **DELETE /todos/:id**
    Delete a todo item.