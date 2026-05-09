Feature: Processing orders

  As a kiosk customer I can see checkout create an order for processing

  Background:
    Given I am a registered customer
    And I start a new basket
    And a store has the following items
      | Name              | Price |
      | Wizard w/ crystal | 9.99  |
    And I add the items
      | Name              | Quantity |
      | Wizard w/ crystal | 2        |
    And I authorize payment for the basket total
    And the payment is authorized
    When I check out the basket

  Scenario: Checkout creates an approved order
    Then the basket is checked out
    And the checked out basket eventually becomes an order
    And an order exists for the checked out basket
    And the order is visible through the ordering API
    And the order belongs to the current customer
    And the order references the authorized payment
    And the order contains the basket items
    And the order status is "approved"
