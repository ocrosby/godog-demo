Feature: Albums
  As a user
  I want to interact with albums and their related photos
  So that I can explore the available albums and their contents using various HTTP methods

  @albums
  Scenario: Create a new album
    Given a new album
    And the new album has a user id of 1
    And the new album has an id of 101
    And the new album has a title of "New Album"
    When I create the new album
    Then there should be no errors
    And the response should be successful
    And the album should have an id of 101


  @albums
  Scenario: Get all albums
    When I request all albums
    Then there should be no errors
    Then the response should be successful
    And there should be 100 albums in the response body
    
  @albums
  Scenario Outline: Get album
    When I request album <albumId>
    Then there should be no errors
    And the response should be successful
    And the album should have an id of <albumId>
    And the album should have a user id of <userId>
    And the album should have a title of "<title>"

    Examples:
      | albumId | userId | title                                    |
      | 1       | 1      | quidem molestiae enim                         |
      | 2       | 1      | sunt qui excepturi placeat culpa         |
      | 3       | 1      | omnis laborum odio                       |
      | 4       | 1      | non esse culpa molestiae omnis sed optio |


  @albums
  Scenario: Delete album 1
    When I delete album 1
    Then there should be no errors
    And the response should be successful

  @albums
  Scenario: Delete album that does not exist
    When I delete album 101
    Then there should be no errors
    And the response should be successful

  @albums
  Scenario: Get album 100
    When I request album 100
    Then there should be no errors
    And the response should be successful
    And the album should have an id of 100
    And the album should have a user id of 10
    And the album should have a title of "enim repellat iste"
