UPDATED DEC 30 2021

https://play.battlesnake.com/g/7a20be20-a259-4455-97cb-8c6ca49ded65/
    - turn 49
    - check lookDistance macro:
        - should have moved right instead of left
        PROBLEM SYNOPSIS:
            - moved left because more spaces were available to the left ("DUH")
            - BUT, left is a trap.
            - didn't consider moving right because the tail of kaa was obstructing path length
            - should've known that the tail would move from that square next turn, so long as kaa doesn't eat (kaa couldn't eat from that position)

https://play.battlesnake.com/g/a2f4144b-172f-459d-8110-2bfb06bc599e/
    - turn 17
    - check food priority
    PROBLEM SYNOPSIS:
        - moved left into food, risking head on collision with other snake of same length potentially also moving into same food
        - would have been the smart move if AngrySnek was just one unit smaller in length... but alas this was not the case.
        - should have turned right to avoid the food (because other snake our same length was immediately about to eat the food as well)


MEGA PROBLEM???
https://play.battlesnake.com/g/4ecff0d8-b2dc-4a39-a4f8-8c0d905d5238/
    - turn 49
    - why on earth did we turn left ?????
    - definitely should have turned right...


https://play.battlesnake.com/g/6bcf5ba4-b5c8-448b-af3f-aab78d0a7c73/
    - turn 97
    - mistaken to turn toward the food in this situation when you're already so long and you're completely boxed in 
    - food spawned at a bad time -- likely would have moved north if food hadn't just spawned where it did.

https://play.battlesnake.com/g/3902b7b2-7b90-413d-a763-72f8b34d04fe/
    - turn 95
    - mistaken to turn north and knot yourself up
    - definitely should have gone left or down in this position -- lookDistance should have known that ????
    - check lookDistance macro

https://play.battlesnake.com/g/dd68a88a-7493-496e-a2e2-439d47275823/
    - turn 72
    - moved left towards food, right in front of larger snake
    - died to larger snake
    - should have moved downward, having farther lookDistance and no risk of death
    - SOLUTION: lower food priority, especially when not so hungry