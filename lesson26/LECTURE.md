# Lesson 26: Event Driven Models

Event brokers are responsible for facilitating communication between event producers and consumers in event-driven
architectures. They help decouple services, ensure reliable message delivery, and enable scalability

## Types

### Pub/Sub (Publish-Subscribe) Model

Pub/Sub (Publish-Subscribe) is an asynchronous messaging pattern where **producers (publishers)** send events to a
central topic, and **consumers (subscribers)** receive messages without direct coupling to producers.

#### Key Features

- **Many-to-many communication**: Multiple publishers can send events to a topic, and multiple subscribers can receive
  them.
- **Event broadcasting**: All subscribers receive published events if they are subscribed to the topic.
- **Loose coupling**: Publishers and subscribers don’t directly communicate; they only interact with the broker.
- **Scalability**: Designed for high-throughput, distributed systems.

#### Tools

- Google Cloud Pub/Sub
- AWS SNS (Simple Notification Service)
- Apache Kafka (in some configurations)
- NATS Streaming
- Redis Pub/Sub

#### Example

Imagine a banking system where multiple services need to react when a transaction is completed:

1. **Payment Service** publishes an `"TransactionCompleted"` event.
2. **Notification Service** and **Analytics Service** subscribe to the topic and receive the event independently.

```aiignore
Publisher (Payment Service) → "TransactionCompleted" → Pub/Sub Broker → Subscribers (Email Service, Fraud Detection)
```

#### Pros

- Broadcasting messages to multiple consumers
- Real-time notifications & event-driven analytics
- Microservices architectures

#### Cons

- Messages are not stored indefinitely (unless using Kafka, which supports durable storage)
- Ensuring order and exactly-once delivery can be complex

### Message Broker (Queue-Based) Model

A message broker is a middleware component that manages message distribution between producers and consumers. Unlike
Pub/Sub, **message queues typically ensure that each message is consumed by only one consumer**.

#### Key Features

- **Point-to-point messaging**: Each message is delivered to a single consumer, ensuring at-least-once processing.
- **Reliable delivery**: Messages are stored in a queue until processed.
- **Load balancing**: Multiple consumers can process messages in parallel, distributing workloads.

#### Tools

- RabbitMQ (traditional message queueing)
- Apache ActiveMQ
- Amazon SQS (Simple Queue Service)
- NATS JetStream
- Kafka (with consumer groups acting as queues)

#### Example

In a task processing system, where a background service needs to process uploaded files:

1. The **Uploader Service** sends `"FileUploaded"` messages to a queue.
2. A pool of **Workers** (consumers) process the messages **one at a time** (ensuring each file is processed only once).

```aiignore
Producer (Uploader Service) → Message Queue → Consumer (File Processor)
```

#### Pros

- Task queues (background jobs, file processing)
- Ensuring message processing once (work queues)
- Load balancing across multiple workers

#### Cons

- Not ideal for broadcasting events to multiple consumers
- Requires consumers to actively pull messages

### Log-Based Event Streaming

Unlike traditional Pub/Sub or message queues, **log-based brokers** persist messages as an ordered, immutable log.
Consumers **replay events** from the log at any time.

#### Key Features

- **Event Retention**: Messages are stored for a set period (or indefinitely).
- **Replayability**: Consumers can reprocess historical events.
- **Scalability**: Distributed, partitioned architecture allows high throughput.
- **Event Sourcing Support**: Can be used as a persistent event store.

#### Tools

- Apache Kafka
- Apache Pulsar
- Redpanda (Kafka-compatible)

#### Example

A **banking application** that needs to track transactions for analytics and fraud detection:

1. **Payment Service** writes events to Kafka (`transactions` topic).
2. **Analytics Service** and **Fraud Detection** service process the events at different times.

```aiignore
Producer → Kafka Topic (Ordered Log) → Consumers (Fraud Detection, Analytics)
```

#### Pros

- High-throughput, real-time data pipelines
- Event sourcing (storing all historical events)
- Streaming analytics (real-time dashboards)

#### Cons

- Complex setup & maintenance
- Managing partitions & scaling consumers

### Event Bus (Enterprise Integration Pattern)

An **event bus** is a higher-level abstraction that routes events between multiple applications, often used in
**enterprise service buses (ESB)** for integrating diverse services.

#### Key Features

- **Centralized event distribution**: Used in large-scale distributed systems.
- **Multiple event routing mechanisms**: Can support Pub/Sub, Queues, and Request-Response patterns.
- **Used in enterprise applications for service integration.**

#### Tools

- AWS EventBridge
- Azure Event Grid
- NATS (core messaging)

#### Pros

- Large-scale enterprise applications
- Complex event-driven architectures with multiple services
- Multi-cloud event routing

#### Cons

- Can introduce latency and single points of failure if not properly designed

## Summary

| Feature               | **Pub/Sub (Google Pub/Sub, SNS, Redis)** | **Message Queue (RabbitMQ, SQS)** | **Event Streaming (Kafka, Pulsar)** | **Event Bus (EventBridge, NATS)** |
|-----------------------|------------------------------------------|-----------------------------------|-------------------------------------|-----------------------------------|
| **Delivery Model**    | Many-to-many                             | One-to-one                        | Ordered event log                   | Enterprise-wide routing           |
| **Message Retention** | Short-lived or transient                 | Persistent until consumed         | Stored in a log (replayable)        | Varies                            |
| **Scalability**       | High                                     | Medium                            | Very High                           | High                              |
| **Use Case**          | Notifications, broadcasts                | Task queues, background jobs      | Data pipelines, event sourcing      | Multi-service orchestration       |
| **Replay Support**    | No (unless stored)                       | No                                | Yes                                 | Depends on implementation         |
| **Common Tools**      | Google Pub/Sub, AWS SNS, Redis           | RabbitMQ, SQS                     | Kafka, Pulsar                       | AWS EventBridge, NATS             |

✅ **Use Pub/Sub if:**

✔ You need real-time broadcasting to multiple subscribers.

✔ You want loose coupling between services.

✅ **Use Message Brokers if:**

✔ You need reliable, exactly-once task processing.

✔ You want load-balanced background processing (e.g., job queues).

✅ **Use Log-Based Streaming if:**

✔ You need event persistence & replayability.

✔ You're building high-throughput, distributed systems.

✅ **Use an Event Bus if:**

✔ You're integrating multiple cloud services & microservices.

✔ You need centralized event routing across an organization.
