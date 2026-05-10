@seeded-demo
Feature: Kiosk Shopping

  Scenario: Seeded demo bootstrap is available for kiosk browsing
    Given I start a new basket for the seeded demo customer
    Then the basket was started

  Scenario: Seeded kiosk checkout creates a live order projection
    Given I start a new basket for the seeded demo customer
    And the basket was started
    And I add the seeded demo items
    And the items were added
    And I fetch the basket snapshot
    When I authorize payment for the basket total
    Then the payment is authorized
    When I check out the basket
    Then the basket is checked out
    And the checked out basket eventually becomes an order
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
