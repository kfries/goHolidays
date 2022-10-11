Feature: Define New Years Day Holiday
  In order to properly identify when a day is the New Years Day Holiday
  As a developer
  I want to have a way to detect, and return this condition

  Background:
    Given I have a USFederalHoliday HolidayList Object

  Scenario: verify the name
    Then the title will be "US Federal Holidays"

  Scenario Outline: verify a the list of dates
    When I execute check with "<dateString>"
    Then it returns "<result>"

    Examples:
    | dateString | result | name             |
    | 2022-01-01 | true   | New Years        |
    | 2022-01-01 | true   | MLK              |
    | 2022-01-01 | true   | Presidents Day   |
    | 2022-01-01 | true   | Memorial day     |
    | 2022-01-01 | true   | Junteenth        |
    | 2022-01-01 | true   | Independence Day |
    | 2022-01-01 | true   | Labor Day        |
    | 2022-01-01 | true   | Columbus Day     |
    | 2022-01-01 | false  | Holloween        |
    | 2022-01-01 | true   | Veterens Day     |
    | 2022-01-01 | true   | Thanksgiving     |
    | 2022-01-01 | true   | Christmas        |
