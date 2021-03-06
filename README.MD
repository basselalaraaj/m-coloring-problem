# m coloring problem

Given an undirected graph and a number m, determine if the graph can be coloured with at most m colours such that no two adjacent vertices of the graph are colored with the same color. Here coloring of a graph means the assignment of colors to all vertices. 

![vertex coloring](https://user-images.githubusercontent.com/5745279/131906463-aa103c3d-5add-408f-aeeb-797d0c0e7e4b.png)


#### Solution
Method: BFS  
Time Complexity: O(n + v)  
Language: Golang


#### Example 1
##### Input
```
graph = {0, 1, 1, 1},
        {1, 0, 1, 0},
        {1, 1, 0, 1},
        {1, 0, 1, 0}
```
##### Output
Following are the assigned colors
```
 1  2  3  2
```
Explanation: By coloring the vertices 
with following colors, adjacent 
vertices does not have same colors

#### Example 2
##### Input
```
graph = {1, 1, 1, 1},
        {1, 1, 1, 1},
        {1, 1, 1, 1},
        {1, 1, 1, 1}
```
##### Output
Solution does not exist.  