Feature: Albums
  As a user
  I want to interact with albums and their related photos
  So that I can explore the available albums and their contents using various HTTP methods

    Background:
        Given I have albums with the following data:
        | id | title | userId |
        | 1  | Album 1 | 1 |
        | 2  | Album 2 | 1 |
        | 3  | Album 3 | 2 |
        And I have photos with the following data:
        | id | title | albumId |
        | 1  | Photo 1 | 1 |
        | 2  | Photo 2 | 1 |
        | 3  | Photo 3 | 2 |

    Scenario: Get all albums
        When I send a GET request to "/albums"
        Then the response status code should be 200
