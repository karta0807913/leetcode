package main

type Snapshot struct {
	Key int
	Val int
}

type SnapshotArray struct {
	snapshot      [][]Snapshot
	snapshotIndex int
}

func BindarySearch(snapshot []Snapshot, target int) (val int) {
	left, right := 0, len(snapshot)
	for left < right {
		mid := (left + right) / 2
		if snapshot[mid].Key < target {
			left = mid + 1
		} else if snapshot[mid].Key > target {
			right = mid
		} else {
			return snapshot[mid].Val
		}
	}
	return snapshot[left-1].Val
}

func Constructor(length int) SnapshotArray {
	snapshot := make([][]Snapshot, length)
	for i := range snapshot {
		snapshot[i] = []Snapshot{{}}
	}
	return SnapshotArray{
		snapshot:      snapshot,
		snapshotIndex: 0,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	snapshot := this.snapshot[index]
	if snapshot[len(snapshot)-1].Key == this.snapshotIndex {
		snapshot[len(snapshot)-1].Val = val
	} else {
		snapshot = append(snapshot, Snapshot{
			Key: this.snapshotIndex,
			Val: val,
		})
		this.snapshot[index] = snapshot
	}
}

func (this *SnapshotArray) Snap() int {
	n := this.snapshotIndex
	this.snapshotIndex += 1
	return n
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	return BindarySearch(this.snapshot[index], snap_id)
}
