# **Clean Architecture Project**

This project was developed as part of the **Goexpert** course to explore and enhance Clean Architecture principles. The primary goal was to implement a `GetOrders` service that:

- Provides multi-protocol support: **gRPC**, **GraphQL**, and **REST**.
- Integrates with **RabbitMQ** for message brokering.

---

## **Table of Contents**

1. [Introduction](#introduction)
2. [Prerequisites](#prerequisites)
3. [Folder Structure](#folder-structure)
4. [Setup Instructions](#setup-instructions)
5. [RabbitMQ Configuration](#rabbitmq-configuration)
6. [API Communication](#api-communication)
   - [REST API](#rest-api)
   - [gRPC API](#grpc-api)
   - [GraphQL API](#graphql-api)
7. [Database Interaction](#database-interaction)

---

## **Introduction**

The project demonstrates Clean Architecture concepts by decoupling core business logic from implementation details. It supports multiple communication protocols and employs RabbitMQ for event-driven messaging. This modular approach ensures scalability and maintainability, making it suitable for real-world applications.

---

## **Prerequisites**

Ensure the following tools and dependencies are installed:

- **Docker & Docker Compose**: For containerized application setup.
- **Golang**: Version 1.19 or later.
- **MySQL Database**: As the primary data store.

---

## **Folder Structure**

```plaintext
CleanArch/
├── App/
│   ├── api/                     # Sample requests for REST, gRPC, and GraphQL
│   ├── configs/                 # Configuration files
│   ├── cmd/ordersystem/         # Main application entry point
│   ├── internal/                # Core application logic
│   │   ├── entity/              # Domain entities
│   │   ├── event/               # Event handling
│   │   ├── infra/               # Infrastructure and external services
│   │   │   ├── database/        # Database layer
│   │   │   ├── graph/           # GraphQL implementation
│   │   │   ├── grpc/            # gRPC services
│   │   │   ├── web/             # REST API handlers
│   ├── pkg/                     # Shared utilities
│   ├── docker-compose.yaml      # Docker setup
│   ├── go.mod                   # Go module dependencies
│   ├── go.sum                   # Dependency checksums
├── Readme/                      # Supporting documentation assets
├── README.md                    # Project documentation
```

---

## **Setup Instructions**

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/LuisGaravaso/desafios-fullcycle.git
    cd desafios-fullcycle/GoexpertCleanArch
    ```

2. **Start Services**:
    ```bash
    cd App
    docker-compose up -d
    ```

3. **Run the Application**:
    ```bash
    cd cmd/ordersystem
    go run main.go
    ```

4. **Check Application Logs**:
    The logs will confirm that the servers are running:
    ```plaintext
    2025/01/19 10:36:52 Loading Configs
    2025/01/19 10:36:52 Connecting to Database
    2025/01/19 10:36:52 Starting Servers
    2025/01/19 10:36:52 Starting web server on port :8000
    2025/01/19 10:36:52 Starting gRPC server on port 50051
    2025/01/19 10:36:52 Starting GraphQL server on port 8080
    2025/01/19 10:36:52 Servers ready to use!
    ```

---

## **RabbitMQ Configuration**

1. Access the RabbitMQ dashboard: `http://localhost:15672`.
2. Default credentials:
   - Username: `guest`
   - Password: `guest`
3. Create queues and exchanges needed for the `GetOrders` service.
4. Bind the queues to exchanges.

For detailed instructions, see [RabbitMQ Setup Guide](Readme/RabbitMQ.md).

---

## **API Communication**

### **REST API**

- **Base URL**: `http://localhost:8000`
- Test with provided requests in `App/api/http_requests.http`.

#### Example Requests

**Create Order**:
```http
POST http://localhost:8000/order HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
    "id": "created-with-http-request",
    "price": 1000,
    "tax": 150
}
```

**Get All Orders**:
```http
GET http://localhost:8000/orders HTTP/1.1
Host: localhost:8000
Content-Type: application/json
```

**Get Order by ID**:
```http
GET http://localhost:8000/order/created-with-http-request HTTP/1.1
Host: localhost:8000
Content-Type: application/json
```

**Get Non-Existent Order**:
```http
GET http://localhost:8000/order/this-id-does-not-exist HTTP/1.1
Host: localhost:8000
Content-Type: application/json
```

---

### **gRPC API**

- **Server Port**: `50051`
- Use the `evans` client for testing.
- Refer to `App/api/grpc_requests.txt` for Sample Requests 

#### Example Commands

1. **Start the gRPC Client**:
    ```bash
    evans -r repl
    ```

2. **Set Package**:
    ```bash
    package pb
    ```

3. **Create Order Service**:
    **Service**
    ```bash
    service CreateOrderService
    ```

    **Call CreateOrder**
    Call
    ```bash
    call CreateOrder
    ```

    Input
    ```plaintext
    id: "created-with-grpc"
    price: 2000
    tax: 300.5
    ```

   Expected Response:
   ```json
   {
     "finalPrice": 2300.5,
     "id": "created-with-grpc",
     "price": 2000,
     "tax": 300.5
   }
   ```

4. **Get Orders Service**:
    **Service**
    ```bash
    service GetOrderService
    ```

    **Call GetAllOrders**
    Call:
    ```bash
    call GetAllOrders
    ```     
    Expected Response: 
    ```
    List with JSONs for all orders in the Database
    ```
    
    **Call GetOrderById**
    Call:
    ```bash
    call GetOrderById
    ```  
    Input:
    ```bash
    id: "created-with-grpc"
    ```
    Expected Response:
   ```json
   {
     "finalPrice": 2300.5,
     "id": "created-with-grpc",
     "price": 2000,
     "tax": 300.5,
     "exists": true
   }
   ```



---

### **GraphQL API**

- **Server URL**: `http://localhost:8080`
- Interact with the GraphQL Playground
- Refer to `App/api/graphql_requests.txt` for Sample Requests 

#### Example Queries

**Create Order**:
```graphql
mutation {
  createOrder(input: { id: "created-with-graphql", price: 4000, tax: 400.5 }) {
    id
    price
    tax
    final_price
  }
}
```

**Get All Orders**:
```graphql
query {
  orders {
    id
    price
    tax
    final_price
  }
}
```

**Get Order by ID**:
```graphql
query {
  order(id: "created-with-graphql") {
    id
    price
    tax
    final_price
    exists
  }
}
```

---

## **Database Interaction**

1. Access the MySQL container:
    ```bash
    docker exec -it mysql bash
    ```

2. Open the MySQL shell:
    ```bash
    mysql -uroot -p
    ```

3. Query the `orders` table:
    ```sql
    use orders;
    select * from orders;
    ```

   **Example Output**:
   ```plaintext
   mysql> select * from orders;
   +---------------------------+-------+-------+-------------+
   | id                        | price | tax   | final_price |
   +---------------------------+-------+-------+-------------+
   | created-with-grpc         |  2000 | 300.5 |      2300.5 |
   | created-with-http-request |  1000 |   150 |        1150 |
   | created-with-graphql      |  4000 | 400.5 |      4400.5 |
   +---------------------------+-------+-------+-------------+
   3 rows in set (0.00 sec)
   ```

---