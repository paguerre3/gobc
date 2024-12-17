# Event Sourcing and CQRS

---
### Event Sourcing  

**Explanation**:

Event Sourcing is a design pattern where the application's state is derived from a sequence of immutable events stored in an event store. Instead of persisting the current state, every change is captured as an event, allowing the state to be reconstructed by replaying these events.  

***Example***:

A **bank account** application uses Event Sourcing.  
- Events:  
  - `AccountOpened(id=123, initialBalance=1000)`  
  - `MoneyDeposited(id=123, amount=500)`  
  - `MoneyWithdrawn(id=123, amount=200)`  
- To calculate the account balance, the system replays these events:  
  `1000 + 500 - 200 = 1300`  

By replaying events, you not only get the current state (`balance=1300`), but also the full history of how it was achieved.  

---

### CQRS (Command Query Responsibility Segregation)  

**Explanation**:

CQRS is a design pattern that separates the responsibilities of **command** operations (writes) and **query** operations (reads) into distinct models. This allows for optimizing read and write operations independently, improving scalability and performance.  

***Example***:

An **e-commerce application**:  
- **Command Model** (Write side):  
  When a user places an order:  
  - Command: `PlaceOrder(customerId=456, items=[item1, item2])`  
  - Event: `OrderPlaced(orderId=789, items=[item1, item2])`  
- **Query Model** (Read side):  
  When the user views their order history:  
  - Query: `GetOrders(customerId=456)`  
  - Result: `[Order {id=789, items=[item1, item2], status="Shipped"}]`  

The write side captures events, while the read side optimizes the data for fast lookups.  
Combined with Event Sourcing, the read model can be rebuilt by replaying the stored events.  

---

### **How Event Sourcing and CQRS Work Together**  
- **Event Sourcing** ensures all changes are recorded as events.  
- **CQRS** uses these events for efficient data writes (command) and builds separate read models optimized for queries.  

For instance, in the e-commerce example:  
- Event: `OrderPlaced(orderId=789, items=[item1, item2])` is recorded (Event Sourcing).  
- Read Model: Precomputes customer order histories for fast display (CQRS).

---

# Original Patterns Definition

[Event Sourcing](https://microservices.io/patterns/data/event-sourcing.html)
[CQRS](https://microservices.io/patterns/data/cqrs.html)