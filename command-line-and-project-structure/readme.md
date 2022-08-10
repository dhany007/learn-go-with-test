1. Previews :
You have been asked to create a web server where users can track how many games players have won.
- GET /players/{name} should return a number indicating the total number of wins
- POST /players/{name} should record a win for that name, incrementing for every subsequent POST

2. Our product owner has a new requirement; to have a new endpoint called /league which returns a list of all players stored. She would like this to be returned as JSON.

3. Our product owner is somewhat perturbed by the software losing the scores when the server was restarted. This is because our implementation of our store is in-memory. She is also not pleased that we didn't interpret the /league endpoint should return the players ordered by the number of wins!

4. Our product owner now wants to pivot by introducing a second application - a command line application.
For now, it will just need to be able to record a player's win when the user types Ruth wins. The intention is to eventually be a tool for helping users play poker.
The product owner wants the database to be shared amongst the two applications so that the league updates according to wins recorded in the new application.