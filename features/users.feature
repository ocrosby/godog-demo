Feature: User
  In order to manage users
  As an admin
  I want to be able to create, update and delete users

  @users
  Scenario: Get all users
      When I send a "GET" request to "/users"
      Then the response should be successful
      And there should be 10 users in the response body
