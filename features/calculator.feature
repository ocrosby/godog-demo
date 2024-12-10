Feature: Calculator
    In order to avoid silly mistakes
    As a math enthusiast
    I want to be able to calculate stuff

    @calculator
    Scenario: Add
        Given I have a new calculator
        When I add 2 and 3
        Then the result should be 5

    @calculator
    Scenario: Subtract
        Given I have a new calculator
        When I subtract 3 from 5
        Then the result should be 2

    @calculator
    Scenario: Multiply
        Given I have a new calculator
        When I multiply 2 by 3
        Then the result should be 6

    @calculator
    Scenario: Divide
        Given I have a new calculator
        When I divide 6 by 2
        Then the result should be 3

    @calculator
    Scenario: Divide by zero
        Given I have a new calculator
        When I divide 6 by 0
        Then the result should be an error

