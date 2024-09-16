package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID   int
	Name string
	Age  int
}
type Node struct {
	index int
	data  *User
	left  *Node
	right *Node
}
type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) insert(user *User) *BinaryTree {
	if t.root == nil {
		t.root = &Node{index: user.ID, data: user}
		return t
	}

	t.root.insert(user)
	return t
}

func (n *Node) insert(user *User) {
	newNode := &Node{index: user.ID, data: user}
	currentNode := n

	for {
		if newNode.index < currentNode.index {
			if currentNode.left == nil {
				currentNode.left = newNode
				return
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				currentNode.right = newNode
				return
			}
			currentNode = currentNode.right
		}
	}
}

func (t *BinaryTree) search(key int) *User {
	if t.root == nil {
		return nil
	}

	return t.root.search(key)
}

func (n *Node) search(key int) *User {
	currentNode := n

	for {
		if currentNode == nil {
			return nil
		}

		if currentNode.index == key {
			return currentNode.data
		} else if key < currentNode.index {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}
}

func generateData(n int) *BinaryTree {
	rand.Seed(time.Now().UnixNano())
	bt := &BinaryTree{}
	for i := 0; i < n; i++ {
		val := rand.Intn(100)
		bt.insert(&User{
			ID:   val,
			Name: fmt.Sprintf("User%d", val),
			Age:  rand.Intn(50) + 20,
		})
	}
	return bt
}

func main() {
	bt := generateData(50)
	user := bt.search(30)
	currentNode := bt.root
	for currentNode != nil {
		fmt.Println(currentNode.index)
		currentNode = currentNode.right
	}
	if user != nil {
		fmt.Printf("Найден пользователь: %+v\n", user)
	} else {
		fmt.Println("Пользователь не найден")
	}
}
