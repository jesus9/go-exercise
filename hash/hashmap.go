package hash

//Key ciao
type Key string

// Value the content of the dictionary
type Value int

// ValueHashtable the set of Items
type ValueHashtable struct {
	Items map[Key]Value
	//lock  sync.RWMutex
}

//Hash the hash() private function uses the famous Horner's method
// to generate a hash of a string with O(n) complexity
/*func Hash(k Key) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}*/

// Put item with value v and key k into the hashtable
func (ht *ValueHashtable) Put(k Key, v Value) {
	//ht.lock.Lock()
	//defer ht.lock.Unlock()
	//i := Hash(k)
	if ht.Items == nil {
		ht.Items = make(map[Key]Value)
	}
	ht.Items[k] = v
}

// Remove item with key k from hashtable
func (ht *ValueHashtable) Remove(k Key) {
	//ht.lock.Lock()
	//defer ht.lock.Unlock()
	//i := Hash(k)
	delete(ht.Items, k)
}

//IfWordExist found a key in the hashtable
func (ht *ValueHashtable) IfWordExist(k Key) int {
	//i := Hash(k)
	if ht.Items[k] != 0 {
		return int(ht.Items[k])
	}
	return 0
}

// Get item with key k from the hashtable
func (ht *ValueHashtable) Get(k Key) Value {
	//ht.lock.RLock()
	//defer ht.lock.RUnlock()
	//i := Hash(k)
	return ht.Items[k]
}

// Size returns the number of the hashtable elements
func (ht *ValueHashtable) Size() int {
	//ht.lock.RLock()
	//defer ht.lock.RUnlock()
	return len(ht.Items)
}

// Increment the value of a key
func (ht *ValueHashtable) Increment(k Key) {
	//i := Hash(k)
	ht.Items[k]++
	/*j := ht.Items[i].Count
	v := &Value{string(k), j}
	v.Count++
	ht.Items[i] = *v*/
}
