# Project Description

This project was developed as part of a technical assessment for the Backend Developer position. The goal is to build an order-matching engine for trading, using Golang as the main programming language and Tarantool as the database.

## Project Structure
- main.go: The main file initializes the connection to Tarantool and simulates order creation and processing.

- engine.go: Contains the matching engine logic, responsible for processing orders, calculating margins, and managing positions.
- order.go: Defines the Order and Position structures, as well as the PNL calculation logic.
- engine_test.go: Unit tests for the matching engine, simulating various order scenarios.

### Development Process

1. Choice of Language and Tools.

Golang was chosen for its efficiency and simplicity in handling concurrency, which is crucial for real-time order processing. Tarantool was selected as the database for its high performance and suitability for heavy read/write workloads.

2. Environment Setup.

We used Docker to create an isolated environment for Golang and Tarantool, ensuring compatibility across different operating systems. Docker Compose was configured to spin up both services and ensure they communicate within the same network.

3. Implementation of the Matching Engine.

The matching engine was designed to process orders and update user positions. The logic includes:

Unrealized PNL Calculation: Based on the simulated current price and the entry price.
Margin Check: Ensures the user has enough margin before processing the order.
Concurrency: sync.Mutex was used to ensure thread-safe operations on Positions.

4. Integration with Tarantool.

We used the Replace function of the Tarantool Go driver to store positions. A savePosition function was created to encapsulate the storage logic and facilitate mocking during testing.

5. Unit Testing.

We developed unit tests to validate the ProcessOrder logic. Mocks were used to simulate the storage function, enabling tests to run without external dependencies.

## How to Run
Requirements
- Docker and Docker Compose
- Golang

Steps to Run the Project
1. Clone the Repository

```
git clone <repo-link> <folder>
cd <folder>
```
2. Start Docker Compose

```
docker-compose up --build
```
or (for newer versions)
```
docker compose up --build
```

3. Executing the test
```
docker-compose exec -it backend go test
```
or (for newer versions)
```
docker compose exec -it backend go test
```

4. Stoping the container
```
docker-compose down
```
### Implementation Decisions and Challenges

- Data Persistence: Tarantool was chosen for its low latency, and the Replace function was used for position management.
- Concurrency and Data Safety: sync.Mutex was employed to ensure thread-safe access to in-memory data structures.
- Order Simulation: The math/rand package was used to generate random orders, simulating a real trading environment.

#### Possible Improvements

- Integration Testing: Add tests to verify interactions between Golang and Tarantool.
- Error Handling: Enhance error capture and handling for specific cases.
- Monitoring and Logging: Implement a more robust logging solution and add monitoring to track performance metrics.
