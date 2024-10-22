# gobc
Block chain compendium

### Blockchain

![blockchain decentralized](./assets/0_blockchain_decentralized.png)

In a **blockchain**, a block is a digital container that holds data, such as transactions, and links to previous blocks, forming a chain. Here’s a brief explanation of key components in a block:

1. **Previous Hash**: Each block contains the hash (a cryptographic code) of the previous block. This links blocks together, ensuring that if any block is tampered with, it breaks the chain’s integrity.

2. **Timestamp**: This records the exact time when the block was created, ensuring that the order of blocks is maintained over time.

3. **Nonce**: A random number used in the mining process to solve a cryptographic puzzle. Miners adjust the nonce to find a hash that meets the blockchain's difficulty requirements, securing the network.

4. **Transactions**: A list of transactions within the block, containing details such as sender, receiver, and the amount of cryptocurrency or data being transferred.

Together, these elements secure the blockchain's immutability and transparency.



---
### SAGA

The **Saga pattern** is a design pattern used to manage distributed transactions in microservices architectures. In traditional monolithic systems, a single transaction could be easily managed with ACID properties (Atomicity, Consistency, Isolation, Durability). However, in a microservices environment, where each service has its own database, it's difficult to maintain a single atomic transaction across multiple services.

The Saga pattern breaks a distributed transaction into a series of smaller transactions. Each of these smaller transactions updates a service and publishes an event to trigger the next step in the process. If a step fails, a compensating transaction is executed to undo the previous step(s), ensuring that the system remains consistent.

There are two main types of Saga implementations:

1. **Choreography-based Saga**:
   - In this approach, each service involved in the saga publishes an event when a transaction is completed. This event triggers the next service in the process to start its transaction.
   - There is no central orchestrator; instead, each service listens to events and reacts accordingly.
   - This approach is simple but can become complex as the number of services grows, leading to a lot of inter-service communication.

2. **Orchestration-based Saga**:
   - In this approach, a central orchestrator manages the flow of the saga. The orchestrator sends commands to each service, telling them what to do at each step.
   - If a step fails, the orchestrator issues compensating commands to rollback or handle the failure.
   - This approach provides more control but introduces a single point of failure (the orchestrator).

### Key Points:
- **Compensating transactions**: Unlike ACID transactions, Sagas rely on compensating transactions to handle failures, essentially undoing the effects of previous steps.
- **Failure management**: Sagas can handle partial failures by rolling back certain steps, while other steps that have already been completed successfully remain intact.
- **Event-driven**: Sagas often leverage event-driven architectures, making them a good fit for microservices that communicate asynchronously.

The Saga pattern is widely used in distributed systems where maintaining strong consistency is difficult or impractical, and eventual consistency is acceptable.



---
### Blockchain vs. SAGA 

Blockchain is **not directly based on the Saga pattern, though they share some common concepts**, particularly regarding **distributed systems** and **failure handling**. Here's a breakdown of how they differ and where the overlap might be:

### Blockchain:
- **Immutable Transactions**: In blockchain, once a transaction is confirmed and added to the block, it is immutable. There's no concept of rolling back or compensating transactions like in a Saga.
- **Consensus Mechanisms**: Blockchain relies on consensus algorithms (e.g., Proof of Work, Proof of Stake) to validate and agree on the order of transactions across a distributed network of nodes. This ensures that all nodes maintain a consistent state of the blockchain.
- **Decentralization**: Blockchain is inherently decentralized, with no single authority controlling the transaction process. This contrasts with the orchestration model of Saga, where a central orchestrator manages the workflow.

### Saga Pattern:
- **Compensating Transactions**: In a Saga, if a step fails, a compensating transaction is executed to undo or correct the previous steps, ensuring consistency in the system.
- **Central or Event-driven Control**: Sagas can be orchestrated centrally (Orchestration-based Saga) or managed in a decentralized, event-driven way (Choreography-based Saga), but they still maintain a predefined flow of operations.
- **Transactional Integrity**: The Saga pattern is focused on managing long-running transactions across distributed services while ensuring eventual consistency through retries or rollbacks.

### Similarities:
- **Distributed Nature**: Both blockchain and the Saga pattern deal with distributed systems where a single atomic transaction across all nodes/services is difficult to achieve.
- **Failure Handling**: Blockchain and Saga patterns both handle failures but in different ways. Sagas handle failures by rolling back or compensating transactions, while blockchain handles failures by rejecting invalid transactions and ensuring consistency via consensus.

### Key Differences:
- **Immutability vs. Rollback**: Blockchain transactions are permanent, while the Saga pattern focuses on compensating for failures by rolling back or compensating actions.
- **Consensus vs. Central Orchestration**: Blockchain uses decentralized consensus, while Sagas (in the orchestration model) may use a central coordinator to manage the transaction flow.

So, while blockchain and the Saga pattern are both tools for managing distributed processes, their underlying principles are quite different. **Blockchain ensures immutability and decentralized consensus, whereas the Saga pattern focuses on compensating for failures in distributed transactions to maintain consistency.**



---
The number of members required for consensus in a blockchain network depends on the specific consensus algorithm being used. Here are **some examples of popular algorithms and their consensus requirements**:

### 1. **Proof of Work (PoW)**
   - **Consensus Members: All participating nodes (miners) can attempt to solve the cryptographic puzzle, but only one miner (the one that solves the puzzle first) adds the block.**
   - **Verification: Once the block is proposed, it is broadcasted to the network, and the "majority" of nodes (51% or more) must agree** that the block is valid.
   - **Example: "Bitcoin" uses PoW**, and although any miner can propose a block, a majority of the network needs to validate and accept the block.

### 2. **Proof of Stake (PoS)**
   - **Consensus Members: In PoS, "validators" are chosen to propose and validate blocks based on the "amount" of cryptocurrency they hold (their stake).**
   - **Verification**: Often, **a majority of validators need to agree on the proposed block** for it to be added to the blockchain.
   - **Example: "Ethereum's" PoS consensus requires at least 2/3 of the validators (or 66%) to agree on a proposed block.**

### 3. **Byzantine Fault Tolerance (BFT)**
   - **Consensus Members: "All" participating nodes must work together to reach consensus**. Typically, BFT systems can tolerate up to 1/3 of nodes acting maliciously or failing.
   - **Verification: A common threshold is that at least 2/3 of the nodes must agree for the transaction** to be considered valid.
   - **Example**: Tendermint, used by the Cosmos blockchain, employs BFT and **requires more than 66% of the validators** to reach consensus.

### 4. **Delegated Proof of Stake (DPoS)**
   - **Consensus Members: A "small number" of elected validators or witnesses are chosen by the network to validate transactions and propose blocks.**
   - **Verification**: Typically, consensus requires a majority of the selected validators.
   - **Example**: EOS uses DPoS, where **21 elected block producers form consensus**.

### 5. **Raft / Paxos (Used in Private Blockchains)**
   - **Consensus Members: In permissioned blockchains**, consensus algorithms like Raft or Paxos are often used.
   - **Verification: These algorithms usually require a majority (51% or more) of nodes** to agree on a transaction before it is committed.
   - **Example**: Hyperledger Fabric can use Raft for ordering and consensus among its nodes.

### Summary:
- The number of members required for consensus depends on the **algorithm**.
- In **PoW**, it's about getting one miner to solve the puzzle, with a majority validating the block.
- In **PoS** and **BFT**, at least 2/3 of the validators must agree.
- In **permissioned blockchains**, a **simple majority (51%)** is often sufficient.

