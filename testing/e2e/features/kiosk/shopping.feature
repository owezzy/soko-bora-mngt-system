@seeded-demo
Feature: Kiosk Shopping

  Scenario: Seeded demo bootstrap is available for kiosk browsing
    Given I start a new basket for the seeded demo customer
    Then the basket was started
