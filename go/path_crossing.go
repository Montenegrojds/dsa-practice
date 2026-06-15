/*
Path Crossing
Easy Topics Company Tags

You are given a string path, where path[i] = 'N', 'S', 'E' or 'W', each representing moving one unit north, south, east, or west, respectively. You start at the origin (0, 0) on a 2D plane and walk on the path specified by path.

Return true if the path crosses itself at any point, that is, if at any time you are on a location you have previously visited. Return false otherwise.

Example 1:

Input: path = "NES"

Output: false

Explanation: Notice that the path doesn't cross any point more than once.

Example 2:

Input: path = "NESWW"

Output: true

Explanation: Notice that the path visits the origin twice.

Constraints:

    1 <= path.length <= 10,000
    path[i] is either 'N', 'S', 'E', or 'W'.


*/


func isPathCrossing(path string) bool {
    c:=[]int{0,0}
    hmap:= make(map[string]bool)
    hmap["0,0"]=true
    for i:=0;i<len(path);i++{
        if path[i]=='N'{
            c[0]++
        }else if path[i]=='S'{
            c[0]--
        }else if path[i]=='E'{
            c[1]++
        }else if path[i]=='W'{
            c[1]--
        }
        key := fmt.Sprintf("%d,%d",c[0],c[1])
        if hmap[key]{
            return true
        }else{
            hmap[key]=true
        }
    

    }
    return false


}