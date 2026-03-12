# Pattern Recognition Checklist

Before solving, force yourself to ask:

## 1) Input Shape
- array
- string
- linked list
- tree
- graph

## 2) Target Type
- exact value
- pair
- count
- longest
- shortest
- top K
- minimum feasible

## 3) Constraints
- Does order matter?
- Is it contiguous?
- Is input sorted (or can sorting help)?
- Are you tracking "best so far" repeatedly?
- Are you traversing relationships?

## Quick Decision Rules (If -> Then)

- If input is `array/string` and target is `exact value (single lookup)` -> `Hash Map / Set` (or `Binary Search` if sorted).
- If input is `array` and target is `pair/triplet sum` and order does not matter -> `Hash Map / Set`.
- If input is `sorted array` and target is `pair/triplet` with order/position relevance -> `Two Pointers`.
- If input is `array/string`, target is `longest/shortest`, and segment must be contiguous -> `Sliding Window`.
- If input is `array/string`, contiguous range sums/counts are needed -> `Prefix Sum` (+ `Hash Map` for sum-to-k variants).
- If target is `first/last position`, `kth by order in sorted data`, or exact in sorted input -> `Binary Search`.
- If target is `minimum feasible` or `maximum possible` with monotonic condition -> `Binary Search on answer`.
- If input has `matching`, `nested`, or `previous/next greater` behavior -> `Stack` (often monotonic stack).
- If input is `linked list` and asks reverse/cycle/middle/merge -> `Linked List pointers` (fast-slow, rewiring).
- If input is `tree` and asks levels/path/depth/nearest -> `DFS/BFS on Tree`.
- If input is `graph` and asks connectivity/shortest hops/components -> `BFS/DFS on Graph`.
- If target is `top K`, `kth largest/smallest`, or streaming best elements -> `Heap / Priority Queue`.
- If target is `best so far` with local optimal choices -> `Greedy` (validate correctness).
- If choices branch combinatorially and need all combinations/valid configs -> `Backtracking`.
- If repeated overlapping subproblems appear (same state recomputed) -> `Dynamic Programming`.

## Ultra-Short Flow

1. `Tree/Graph` input -> think `DFS/BFS` first.
2. `Linked List` input -> think pointer techniques first.
3. `Contiguous array/string` objective -> think `Sliding Window` or `Prefix Sum`.
4. `Sorted` + search/feasibility -> think `Binary Search` (possibly on answer).
5. `Pair/triplet` in sorted array -> `Two Pointers`; unsorted -> `Hash Map`.
6. `Top K` / stream extrema -> `Heap`.
