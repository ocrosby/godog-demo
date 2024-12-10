Feature: Todos
  As a user
  I want to interact with todos
  So that I can explore the available todos and their contents using various HTTP methods

  @todos
  Scenario: Get all todos
      When I send a "GET" request to "/todos"
      Then the response should be successful
      And there should be 200 todos in the response body
