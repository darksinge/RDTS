# Orcinator (Resilient Distributed Transactions)

## Introduction

The goal of Orcinator is to provide a reliable pattern for guaranteeing ACD
transactions (*atomic, consistent, durable*) in a distributed system.

In a typical monolithic application, it's possible to guarantee ACID
transactions. Note the "I" in ACID, which stands for "isolation", which by
definition is not possible in a distributed system. Hence, Orcinator guarantees ACD
transactions, but not ACID transactions.

## Terminology

- __DTRX (*Distributed Transaction*)__: A transaction that is executed across
  multiple nodes in a distributed system. A DTRX is comprised of multiple,
  individual transactions that are executed in a state machine.
- __Compensating Transaction__: A transaction that is executed to roll back changes
  in response to a failed transaction.
- __Saga__: A pattern used in microservices architecture to execute a series of
  transactions in a distributed system.

## Features of Orcinator

1. Will provide at least once delivery guarantee of all messages in a DTRX
   within a given amount of time.
2. If any transaction in a DTRX fail, Orcinator will attempt to rollback the
   transactions in a DTRX by publishing new messages to trigger compensating
   transactions.
3. If any compensating transactions fail or Orcinator detects a bad state, Orcinator will
   pause any new transactions, and alert administrators of the failure. Orcinator will
   remain in this state until human intervention acknowledges the failure.

## Example
Suppose we have a distributed system with three nodes, A, B, and C.  Node A
publishes an event to which nodes B and C are subscribed. The local transaction
succeeds in node A and B, but fails on node C.  Node A, B, and C have all
defined compensating transactions that will be executed in the event of a
failure.

When node C's local transaction fails, one of the following can happen:
- Node C fails gracefully and notifies Orcinator of the failure.
- Node C fails silently and *does not* notify Orcinator of the failure.

In the first case, Orcinator will immediately publish compensating transactions to
node A, B, and C.

In the second case, Orcinator will wait for a timeout period before publishing
compensating transactions to node A, B, and C.
