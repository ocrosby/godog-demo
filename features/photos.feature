Feature: Photos
    As a user
    I want to see photos
    So that I can see what the place looks like

    @photos
    Scenario: Get all photos
      When I send a "GET" request to "/photos"
      Then the response should be successful
      And there should be 5000 photos in the response body
