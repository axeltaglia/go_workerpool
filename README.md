# Worker Pool API Server

This project is a simple Go-based API server with a task manager that manages a worker pool for handling concurrent tasks. It allows tasks to be created and processed asynchronously, reporting their progress.

## Project Structure

- **`main.go`**: The entry point of the application where the API server and task manager are initialized and started.
- **`apiServer`**: The package that handles HTTP routes and endpoints.
- **`taskManager`**: The package responsible for managing tasks and processing them using a worker pool.

## How to Run the Project

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/workerpool.git
   cd workerpool
   ```

2. **Build and run the project**:
   ```bash
   go run main.go
   ```

3. **Server Information**:
   The server runs on `http://localhost:8082`.

## API Services

### 1. Create a Task

- **Endpoint**: `POST /api/tasks/createTask`
- **Description**: Adds a new task to the task queue with a given task ID and workload duration.
- **Request Parameters**:
    - `taskID` (query parameter): The ID of the task (e.g., `task123`).
    - `workload` (query parameter): The workload duration in seconds (e.g., `5`).

- **Example Request**:
  ```http
  POST http://localhost:8082/api/tasks/createTask?taskID=task123&workload=5
  ```

- **Example Response**:
  ```text
  Task task123 created with workload 5 seconds.
  ```

## Package Descriptions

### Task Manager (`taskManager` package)

- **`Task` Interface**:
  ```go
  type Task interface {
  Process()      // The function that does the actual task work
  GetID() string // A method to get the task ID
  }
  ```

- **`TaskManager` Struct**:
  Manages a channel of tasks and provides methods to add and process them.

  ```go
  type TaskManager struct {
  taskQueue chan Task // Buffered channel for tasks
  }
  ```

- **Methods**:
    - `NewTaskManager(bufferSize int) *TaskManager`: Initializes a `TaskManager` with the specified buffer size.
    - `AddTask(task Task)`: Adds a task to the queue.
    - `Start()`: Runs a worker that continuously processes tasks from the queue.

### Concrete Task (`taskManager` package)

- **`ConcreteTask` Struct**:
  A concrete implementation of the `Task` interface. It simulates work by sleeping for a specified duration and reports progress.

  ```go
  type ConcreteTask struct {
  ID       string
  Workload int // Simulates the amount of work (sleep duration)
  }
  ```

- **Methods**:
    - `Process()`: Simulates work by sleeping for `Workload` seconds and prints the progress to the console.
    - `GetID()`: Returns the ID of the task.

## Example Output

When tasks are processed, the console shows real-time progress:

```text
ApiServer started in address :8082...
Task with ID: task123 added to the queue
Task task123. Progress: 0%
Task task123. Progress: 20%
Task task123. Progress: 40%
Task task123. Progress: 60%
Task task123. Progress: 80%
Task task123. Progress: 100%
```

## Notes

- The `TaskManager` runs as a worker pool with a buffered channel, allowing multiple tasks to be processed concurrently.
- Each `ConcreteTask` reports its progress every second until completion, updating the console in real-time.

## Future Improvements

- Add error handling and logging.
- Implement more sophisticated task types with different processing logic.
- Integrate a database or persistent storage to track task statuses.
- Add a web UI to visualize task progress.

`;
