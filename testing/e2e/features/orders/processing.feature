@seeded-demo
Feature: Processing orders

  Scenario: Seeded demo monolith happy path completes across services
    Given I start a new basket for the seeded demo customer
    And the basket was started
    And I add the seeded demo items
    And the items were added
    And I fetch the basket snapshot
    And I authorize payment for the basket total
    And the payment is authorized
    When I check out the basket
    Then the basket is checked out
    And the checked out basket eventually becomes an order
    And an order exists for the checked out basket
    And the order belongs to the current customer
    And the order references the authorized payment
    And the order contains the basket items
    And the order status is "approved"
    And the depot shopping list exists for the order
    And the shopping list status is "available"
    When I assign the shopping list to bot "demo-bot-001"
    Then the shopping list status is "assigned"
    When I complete the shopping list
    Then the shopping list status is "completed"
    And the order status is "ready"
    And an invoice exists for the checked out basket
    And a "ready" notification exists for the order
    And the search projection status is "Ready For Pickup"
    When I pay the invoice for the checked out basket
    Then the invoice status is "paid"
    And the order status is "completed"
    And the order has an invoice id
    And a "created" notification exists for the order
    And the search projection status is "Completed"
