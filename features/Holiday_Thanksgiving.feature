Feature: Define New Years Day Holiday
  In order to properly identify when a day is the New Years Day Holiday
  As a developer
  I want to have a way to detect, and return this condition

  // 	4,
  // 	time.Thursday,
  // 	time.November,
  // }

  Background:
    Given I have a Thanksgiving Holiday Object

  Scenario: verify the name
    Then the name will be "Thanksgiving"

  Scenario: verify the description
    Then the description will be "4th thursday in November"

  Scenario: verify a valid date in 2020
    When I check "2020-11-26"
    Then it returns true

  Scenario: verify a valid date in 2021
    When I check "2021-11-25"
    Then it returns true

  Scenario: verify a valid date in 2022
    When I check "2022-11-24"
    Then it returns true

  Scenario: verify a non-valid date
    When I check "2022-01-01"
    Then it will return false
