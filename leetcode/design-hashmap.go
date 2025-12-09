package leetcode

type MyHashMap struct {
	Memory [1000][]MyKeyValue
}

type MyKeyValue struct {
	Key   int
	Value int
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (this *MyHashMap) Put(key int, value int) {
	hashKey := getHashKey(key)
	keyValue := findKeyValueByKey(this.Memory[hashKey], key)
	if keyValue == nil {
		this.Memory[hashKey] = append(this.Memory[hashKey], MyKeyValue{Key: key, Value: value})
	} else {
		keyValue.Value = value
	}
}

func (this *MyHashMap) Get(key int) int {
	hashKey := getHashKey(key)
	keyValue := findKeyValueByKey(this.Memory[hashKey], key)
	if keyValue == nil {
		return -1
	}
	return keyValue.Value
}

func (this *MyHashMap) Remove(key int) {
	hashKey := getHashKey(key)
	i := findKeyValueIndexByKey(this.Memory[hashKey], key)
	this.Memory[hashKey] = removeKeyInIndex(this.Memory[hashKey], i)
}

func getHashKey(key int) int {
	return key % 1000
}

func findKeyValueByKey(keyValues []MyKeyValue, key int) *MyKeyValue {
	i := findKeyValueIndexByKey(keyValues, key)
	if 0 <= i {
		return &(keyValues[i])
	}
	return nil
}

func findKeyValueIndexByKey(keyValues []MyKeyValue, key int) int {
	for i, keyValue := range keyValues {
		if keyValue.Key == key {
			return i
		}
	}
	return -1
}

func removeKeyInIndex(keyValues []MyKeyValue, i int) []MyKeyValue {
	if 0 <= i {
		lastIndexInBucket := len(keyValues) - 1
		keyValues[i] = keyValues[lastIndexInBucket]
		keyValues = keyValues[:lastIndexInBucket]
	}
	return keyValues
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
