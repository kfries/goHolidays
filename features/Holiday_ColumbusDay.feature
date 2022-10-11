Feature: Define New Years Day Holiday
  In order to properly identify when a day is the New Years Day Holiday
  As a developer
  I want to have a way to detect, and return this condition

  Background:
    Given I have a ColumbusDay Holiday Object

  Scenario: verify the name
    Then the name will be "Columbus Day / Indigenous People Day"

  Scenario: verify the description
    Then the description will be "Second Monday in October"

  Scenario: verify a valid date in 2020
    When I check "2020-10-12"
    Then it returns true

  Scenario: verify a valid date in 2021
    When I check "2021-10-11"
    Then it returns true

  Scenario: verify a valid date in 2022
    When I check "2022-10-10"
    Then it returns true

  Scenario: verify a non-valid date
    When I check "2022-01-01"
    Then it will return false
