# Problem

There are N players standing in a row, one player on a field. They are numbered from 0 to N−1 fror left to right. 

Players perform moves one by one from left to right, that is, in ascending order of numbers. Each player presses an arrow key in one of the four cardinal directions: left (' < '), right ('' > '), up (' '^') or down (' v '). A key press in the given direction means that the player attempts to move onto the closest field in the direction specified. A move can be performed only if there is no other player already standing on the target field.

Moves are represented as a string S of length N, where S[K] (for K within the range 0..N−1 ) is the direction of the K-th player's move. How many players will actually perform a move successfully? 

Write a function: 

```class Solution { public int solution(String S); }```

 which, given a string S of length N representing arrow keys pressed by each of the players, returns the number of players that will perform a move successfully. 
 
 Examples:
 
 1. Given S= `><^v` , your function should return 2. Player 0 cannot move rightwards, because play 1 is standing on the target field. Player 1 cannot move leftwards, because player 0 is standing on the target field. Players 2 and 3 can both perform their moves because there are no other players standing on their target fields. In the pictures below, players that will perform their moves successfully are marked green and players whose moves will fail are marked red. 
 
 2. Given S= `<<^<v>>`, your function should return 6. Players 0,1,2,3,4 can all perform their moves because there are no other players standing on their target fields. Player 5 pressed the right-arrow key, but the move cannot be performed as player 6 is already standing on the target field. The move of player 6 can be performed, though, as there is no other player standing on the target field

