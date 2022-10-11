Feature: Define New Years Day Holiday
  In order to properly identify when a day is the New Years Day Holiday
  As a developer
  I want to have a way to detect, and return this condition

  Background:
    Given I have a Juneteenth Holiday Object

  Scenario: verify the name
    Then the name will be "Juneteenth"

  Scenario: verify the description
    Then the description will be "June 19th"

  Scenario: verify a valid date
    When I check "2022-06-19"
    Then it returns true

  Scenario: verify a non-valid date
    When I check "2022-12-25"
    Then it will return false
