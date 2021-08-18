// code based on Junmin Lee tutorial
package main

import "fmt"


const ArraySize int  = 7

// backet Node
type bucketNode struct {
	key string 	// value represents the key
	next *bucketNode // points to next node
}

// newBucketNode function to create new bucketNode
func newBucketNode(value string) *bucketNode {
	bucketNode := new(bucketNode)
	bucketNode.key = value
	return bucketNode
}

// bucket
type bucket struct {
	head *bucketNode // Head contains linked list head
	lenght int // lenght of the bucket
}

// insert
func (b *bucket) insert(key string) error {
	if !b.search(key) {
		newNode := newBucketNode(key)
		newNode.next = b.head
		b.head = newNode
		b.lenght++
		return nil
	} 
	return fmt.Errorf("%s already exist", key)
}

// search
func (b *bucket) search(key string) bool {
	for currentNode := b.head ;currentNode != nil ; currentNode = currentNode.next {
		if currentNode.key == key {
			return true
		}
	}
	return  false
}

// delete
func (b *bucket) delete(key string) error {
	if b.head == nil {
		return fmt.Errorf("bucket empty")
	}
	if b.head.key == key {
		b.head = nil
		return nil
	}
	node := b.head
	for; node.next != nil && node.next.key != key ; node = node.next {
	}

	if node.next == nil {
		return fmt.Errorf("key do not exist")
	}
	node.next = node.next.next

	return nil
}


// Print print all the elements in the list
func (l bucket) Print() {
	for node := l.head ; node != nil ; node = node.next{
		fmt.Printf("%s ", node.key)
	}
	fmt.Println()
}

// HashTable represents hash table data struture
type HashTable struct {
	array [ArraySize]*bucket

}

// Insert element into hash table
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Delete an element from hash table
func (h *HashTable) Delete(key string) error {
	index := hash(key)
	if index != -1 {
		return h.array[index].delete(key)
	}
	return fmt.Errorf("key do not exist")
}

// Search if element exist in hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Search if element exist in hash table
func (h *HashTable) Print() {
	for _, bucket := range h.array {
		if bucket != nil {
			bucket.Print()
		}
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init
func Init() *HashTable {
	result := &HashTable{ }
	for index :=  range result.array {
		result.array[index]	= new(bucket)
	}
	return result
}

func main() {
	hash := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}
	for _, v := range list {
		hash.Insert(v)
	}
	hash.Print()
	hash.Delete("KENNY")
	hash.Print()
	hash.Delete("BUTTERS")
	hash.Print()
	hash.Delete("TOKEN")
	hash.Print()
	err := hash.Delete("TEST")
	if err != nil {
		fmt.Println(err.Error())
	}
	hash.Print()
	/*testBucket := &bucket{ }
	testBucket.insert("RANDY")
	testBucket.insert("VIVEK")
	fmt.Println(testBucket.search("RANDY"))
	*/
}