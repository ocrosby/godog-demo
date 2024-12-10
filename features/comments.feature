Feature: Comments
    As a user
    I want to interact with comments
    So that I can explore the available comments and their contents using various HTTP methods

    @comments
    Scenario: Create a new comment
        Given a new comment
        And the new comment has a post id of 1
        And the new comment has an id of 501
        And the new comment has a name of "New Comment"
        And the new comment has an email of "fred.flintstone@bedrock.com"
        And the new comment has a body of "Yabba Dabba Doo!"
        When I create the new comment
        Then there should be no errors
        And the response should be successful
        And the comment should have an id of 501

    @comments
    Scenario: Get all comments
        When I send a "GET" request to "/comments"
        Then the response should be successful
        And there should be 500 comments in the response body


    @comments
    Scenario: Get comment 13
        When I request comment 13
        Then there should be no errors
        And the response should be successful
        And the comment should have an id of 13
        And the comment should have a post id of 1
        And the comment should have a name of "dolorum ut in voluptas"
        And the comment should have an email of "something"
        And the comment should have a body of "quidem molestiae enim"

    @comments
    Scenario: Delete a comment
        When I delete a comment with id 1
        Then there should be no errors
        And the response should be successful

    @comments
    Scenario: Delete a comment that does not exist
        When I delete a comment with id 501
        Then there should be no errors
        And the response should be successful
