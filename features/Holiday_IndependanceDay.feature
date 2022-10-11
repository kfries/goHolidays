Feature: Define New Years Day Holiday
  In order to properly identify when a day is the New Years Day Holiday
  As a developer
  I want to have a way to detect, and return this condition

  Background:
    Given I have a IndependenceDay Holiday Object

  Scenario: verify the name
    Then the name will be "Independence Day"

  Scenario: verify the description
    Then the description will be "July 4th"

  Scenario: verify a valid date
    When I check "2022-07-04"
    Then it returns true

  Scenario: verify a non-valid date
    When I check "2022-12-25"
    Then it will return false
