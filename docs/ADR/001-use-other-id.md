---
status: proposed
date: 2026-07-13
decision-makers: me
consulted: me
informed: me 
---
# Use an other id system

## Context and problem statement

The current id system is based on autoincrement al integers which are not suitable for security concerns : it exposes publicly primary key and records length

## Considered options

- UUID v4
- UUID v7
- ULID
- Surrogate key based on one the previous technologies

## Decision drivers

- uuidv4 si totally random with very poor risk of collision, but as can impact database performance if used as a primary key when there is a big number of records
- uuidv7 are ULID are partially random and contains a timestamp that makes them easily sortable and keep database operations (insertion, reading..) performant
- But the timestamp part can give information about the creation date of the record, which is not tolerated in some contexts.


## Decision outcome

UUIDv4 as a surrogate key allows to keep autoincremental integer as primary keys for backend concerns and will never be exposed.
Only the surrogate key will be exposed , giving no information to a potential attacker

### Consequences

A publicId field is added to every model that needs to be exposed.

### Confirmation
