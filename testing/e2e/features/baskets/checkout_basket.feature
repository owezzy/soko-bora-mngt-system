Feature: Checking out baskets

  As a customer with items in a basket I can authorize payment and check out

  Background:
    Given I am a registered customer
    And I start a new basket
    And a store has the following items
      | Name              | Price |
      | Wizard w/ crystal | 9.99  |
    And I add the items
      | Name              | Quantity |
      | Wizard w/ crystal | 2        |

  Scenario: Checking out a basket with an authorized payment
    When I authorize payment for the basket total
    Then the payment is authorized
    And a payment record exists for the basket total
    And the authorized payment belongs to the current customer
    And the authorized payment amount is "19.98"
    When I check out the basket
    Then the basket is checked out

  @seeded-demo
  Scenario: Checking out a seeded demo basket with an authorized payment
    Given I start a new basket for the seeded demo customer
    And the basket was started
    And I add the seeded demo items
    And the items were added
    And I fetch the basket snapshot
    When I authorize payment for the basket total
    Then the payment is authorized
    And a payment record exists for the basket total
    And the authorized payment belongs to the current customer
    And the authorized payment amount is "30"
    When I check out the basket
    Then the basket is checked out
