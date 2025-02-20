# Lesson 25: Event Driven Development

## Core Concepts

### Events

Events are signals or messages that indicate that something has happened. Events can originate from various sources such
as user interactions (e.g., button clicks), sensors, external systems, or internal system changes.

Example:

- A user submits a form → `UserRegistered` event is fired.
- A payment is completed → `PaymentProcessed` event is emitted.

### Producers

Producers generate and emit events when something significant occurs. They do not need to know which components will
handle these events.

Example:

- A checkout system emits a `PaymentSuccess` event when a transaction is completed.

### Consumers (Subscribers)

Consumers listen for and respond to events by executing appropriate logic. They can be services, functions, or processes
that react asynchronously to events.

Example:

- A notification service listens for `PaymentSuccess` and sends a confirmation email to the user.

### Event Bus or Message Broker

The event bus is a communication channel that routes events between producers and consumers. Popular event brokers
include:

- Kafka (distributed streaming platform)
- RabbitMQ (message queue)
- NATS (lightweight messaging system)
- AWS EventBridge (serverless event bus)
- Redis Pub/Sub (lightweight event propagation)

## Pros

1. **Loose Coupling & High Modularity**

   _Benefit_: Components interact via events rather than direct dependencies, making it easier to replace, update, or
   extend functionality without affecting the entire system.

   _Example_: A payment service can emit an "OrderPaid" event, and multiple independent services (e.g., notifications,
   invoicing, analytics) can react without modifying the payment service.


2. **Scalability & Performance**

   _Benefit_: Since event-driven systems handle events asynchronously, they scale well under high loads by distributing
   processing across multiple consumers.

   _Example_: In a high-traffic e-commerce platform, order processing can be distributed among multiple consumers to
   prevent bottlenecks.


3. **Asynchronous Processing**

   _Benefit_: Events allow background processing without blocking the main application flow, leading to faster response
   times.

   _Example_: A user uploads a file, and while it's being processed (resized, converted), they can continue using the
   application.


4. **Resilience & Fault Tolerance**

   _Benefit_: If an event consumer crashes, the event is still in the queue and can be reprocessed once the service
   recovers.

   _Example_: If an email notification service goes down, it can still process "UserRegistered" events once it restarts.


5. **Real-Time Event Processing**

   _Benefit_: Ideal for applications that require real-time updates, such as stock trading, gaming, or IoT monitoring.

   _Example_: A real-time analytics dashboard updates immediately when new user actions occur.


6. **Extensibility & Flexibility**

   _Benefit_: New features can be added by introducing new event consumers without modifying existing producers.

   _Example_: Adding a fraud detection service that listens for "PaymentProcessed" events without changing the payment
   system.


7. **Better Handling of Distributed Systems & Microservices**

   _Benefit_: EDA naturally supports microservices by enabling services to communicate asynchronously.

   _Example_: An order service in an e-commerce system publishes "OrderPlaced", and separate services (shipping,
   billing, notifications) react independently.

## Cons

1. **Increased Complexity**

   _Issue_: EDA introduces complexity in managing event flows, debugging, and tracking dependencies between producers
   and consumers.

   _Example_: A system with multiple microservices emitting and consuming events can be hard to debug when failures
   occur.

   _Mitigation_: Use event logging, tracing tools (OpenTelemetry, Jaeger), and structured event schema registries (
   Apache Avro, JSON Schema).


2. **Event Ordering Challenges**

   _Issue_: In distributed environments, events may not always be processed in order.

   _Example_: "OrderShipped" event might be processed before "PaymentReceived", causing inconsistencies.

   _Mitigation_: Use event versioning, timestamps, and partitioning in event brokers like Kafka.


3. **Duplicate Event Processing & Idempotency**

   _Issue_: Events might be delivered multiple times due to retries, leading to duplicate processing.

   _Example_: A "PaymentProcessed" event could be processed twice, leading to double charging.

   _Mitigation_: Implement idempotency by tracking processed event IDs or using transactional outboxes.


4. **Debugging & Observability Challenges**

   _Issue_: Debugging is harder compared to synchronous request-response architectures.

   _Example_: If a user does not receive a confirmation email, tracing the event from "UserRegistered" to "SendEmail"
   can be difficult.

   _Mitigation_: Use distributed tracing (OpenTelemetry, Jaeger) and structured logging.


5. **Latency & Performance Overhead**

   _Issue_: Asynchronous processing introduces additional latency due to message queuing and event handling delays.

   _Example_: A user submits an order, but the confirmation email is delayed due to queue processing.

   _Mitigation_: Optimize event broker configurations, implement priority queues, and scale consumers dynamically.


6. **Schema Evolution & Compatibility Issues**

   _Issue_: Changes in event structure (e.g., adding/removing fields) may break consumers.

   _Example_: Adding a new customer_id field in "OrderPlaced" could cause old consumers to fail if they expect a
   different schema.

   _Mitigation_: Use schema versioning, JSON Schema, or Apache Avro with a schema registry.


7. **Higher Infrastructure & Tooling Requirements**

   _Issue_: Managing event brokers (Kafka, RabbitMQ, NATS) requires additional infrastructure and operational effort.

   _Example_: Running Kafka at scale requires configuring brokers, partitions, consumer groups, and monitoring tools.

   _Mitigation_: Consider managed services (AWS EventBridge, Azure Event Grid, Google Pub/Sub) for reducing
   infrastructure complexity.

## Patterns

**Simple Event Notification**

Producer emits an event, and one or more consumers react. There is no expectation of a response.

Example:

1. A user uploads a file → "FileUploaded" event is emitted.
2. A background job resizes images without blocking the upload process.

**Event-Carried State Transfer**

Events contain the data required for processing, reducing the need for additional queries.

Example:

1. `OrderPlaced` event carries order details, so the consumer does not need to query the database.

**Event Sourcing**

Instead of storing the latest state, the system persists a sequence of events to reconstruct state later.

Example:

1. A bank application records `MoneyDeposited` and `MoneyWithdrawn` events to calculate an account balance.

**CQRS (Command Query Responsibility Segregation)**

Commands (write operations) and queries (read operations) are separated, improving scalability.

Example:

1. A user places an order (Command), and another service queries orders (Query).