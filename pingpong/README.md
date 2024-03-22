# Problem
A **Ping-Pong Simulator** is a conceptual system designed to simulate the interaction between two processes, akin to players in a ping-pong game, using the principles of concurrency. The environment starts with two players ready to exchange messages through a series of operations that simulate the game. The operations involved are:

* "Ping": Indicates a message sent from the first player to the second.
* "Pong": Indicates a response from the second player back to the first.
The simulation is configured to run for a predefined number of rounds, with each round consisting of a "Ping" sent and a "Pong" received, demonstrating the synchronization and communication between concurrent processes.

The goal is to ensure the messages are exchanged in the correct order, showcasing the ability to manage concurrent operations and maintain synchronization between them. The simulation must start with a "Ping" and alternate with "Pong", accurately reflecting the exchange between the two players for the specified number of rounds.

Write a function: func simulatePingPong(rounds int) that orchestrates this ping-pong message exchange between two concurrent processes for a given number of rounds rounds. This function is responsible for setting up the game, managing its execution, and ensuring a graceful shutdown, all while maintaining the correct order and count of message exchanges.

Challenges include implementing effective concurrency control, managing resource synchronization, and providing clear, understandable feedback about the simulation's progress and outcome.

**Assumptions**:

* The number of rounds is a positive integer, set within a practical range for the simulation.
* The simulation accurately adheres to the specified sequence and number of message exchanges.
* Adequate error handling is in place to address potential synchronization challenges or unforeseen termination scenarios.
  
Focus on the correctness, efficiency of your concurrency mechanisms, and the clarity of simulation output. The robustness of your implementation will be a crucial aspect of the assessment.