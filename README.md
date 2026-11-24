# Max-Heap / Min-Heap Implementation

> Striver SDE Sheet 069 · Day 12 Heaps — one array-backed heap with sift-up and sift-down; max and min are the same code with the comparison flipped. Go, stdlib only, tested.

## Problem

Implement a binary heap from scratch — Push, Pop, Peek — in both max and min flavors. The sheet marks this "only for interviews": production Go uses `container/heap`, but you're expected to write the mechanics cold.

## Approach

A binary heap is a complete binary tree flattened into a slice: node i's children live at 2i+1 and 2i+2, its parent at (i−1)/2 — no pointers, perfect locality. Push appends and *sifts up*, swapping with the parent while it outranks it; Pop swaps the last element into the root, shrinks, and *sifts down*, always swapping with the higher-priority child. Max vs min differ in exactly one comparison, so both flavors share the implementation through a `before` function — the same trick `sort.Slice` uses. The only invariant anywhere: every parent outranks its children.

**Complexity:** Push/Pop O(log n), Peek O(1), stored in exactly n ints.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
