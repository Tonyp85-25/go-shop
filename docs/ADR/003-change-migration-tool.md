---
status: proposed
date: 2026-07-20
decision-makers: me
consulted: me
informed: me 
---
# Change database migration tool

## Context and problem statement

Golang migrate CLI has very few commands and slows migration process by adding a new version in database even if migration failed. Forcing to fix script and manually set previous version field to previous one, otherwise no migration is applied to database.

## Considered options

- Goose
- sqlx

## Decision drivers

- presence of a CLI
- tool present in versionned files

## Decision outcome

Goose provides a CLI with numerous options and is easier to use 

### Consequences

New env varables are added for Goose

### Confirmation
Goose advantages works as expected