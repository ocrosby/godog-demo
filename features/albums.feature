Feature: Albums
  As a user
  I want to interact with albums and their related photos
  So that I can explore the available albums and their contents using various HTTP methods

  @albums
  Scenario: Get all albums
    When I send a "GET" request to "/albums"
    Then the response should be successful
    And there should be 100 albums in the response body
    
