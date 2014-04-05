// Copyright 2014 Garrett D'Amore
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sp

// ListNode represents a node in a doubly linked list.  It is suitable for
// embedding directly into structures.  The intention is that this can be
// used to create linked lists that require no allocation activity.
type ListNode struct {
	next *ListNode
	prev *ListNode
	list *List
	// Value should contain either a pointer back to the enclosing
	// structure, or the value itself.
	Value interface{}
}

// List represents a doubly linked list.  Unlike container/list, this
// version relies on the consumer to preallocate nodes.  This has the
// advantage of avoiding extra allocation activity for lists that are
// used with a lot of insert/remove activity.
type List struct {
	ListNode
}

func (l *List) Init() {
	if l.list == nil {
		l.next = &l.ListNode
		l.prev = &l.ListNode
		l.list = l
	}
}

func (l *List) InsertHead(n *ListNode) {
	n.next = l.next
	n.prev = l.next.prev
	n.next.prev = n
	n.prev.next = n
	n.list = l
}

func (l *List) InsertTail(n *ListNode) {
	n.prev = l.prev
	n.next = l.prev.next
	n.next.prev = n
	n.prev.next = n
	n.list = l
}

func (l *List) HeadNode() *ListNode {
	if l.next == &l.ListNode {
		return nil
	}
	return l.next
}

func (l *List) TailNode() *ListNode {
	if l.prev == &l.ListNode {
		return nil
	}
	return l.prev
}

func (l *List) RemoveHead() *ListNode {
	n := l.HeadNode()
	if n != nil {
		l.Remove(n)
	}
	return n
}

func (l *List) RemoveTail() *ListNode {
	n := l.TailNode()
	if n != nil {
		l.Remove(n)
	}
	return n
}

func (l *List) Remove(n *ListNode) {
	if n.list != l {
		if n.list != nil {
			panic("Attempt to remove from wrong list!")
		}
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = nil
	n.next = nil
	n.list = nil
}
