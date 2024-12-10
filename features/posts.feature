Feature: Post
    In order to manage posts
    As a user
    I want to be able to create, read, update and delete posts

    @posts
    Scenario: Delete a post
        When I delete a post with id 1
        Then there should be no errors
        And the response should be successful
