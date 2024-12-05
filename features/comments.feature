Feature: Comments
    As a user
    I want to interact with comments
    So that I can explore the available comments and their contents using various HTTP methods

    @comments
    Scenario: Get all comments
        When I send a "GET" request to "/comments"
        Then the response should be successful
        And there should be 500 comments in the response body
