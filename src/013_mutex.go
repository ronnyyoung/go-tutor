package main

import (
    "fmt"
    "sync"
)


// 支持多线程的map
type MTMap struct {
    inter_map map[string]string
    // 这里使用指针是为了防止互斥号被浅拷贝
    *sync.RWMutex // 匿名成员，可以通过MTMap直接访问RWMutex的方法
}

func (m *MTMap) Insert(key string, value string) (string, bool){
    m.Lock()
    defer m.Unlock()
    old_value, ok := m.inter_map[key]
    m.inter_map[key] = value
    return old_value, ok
}

func (m *MTMap) Get(key string) (string, bool){
    m.RLock()
    defer m.RUnlock()
    old_value, ok := m.inter_map[key]
    return old_value, ok
}

func (m *MTMap) Delete(key string) (string, bool){
    m.Lock()
    defer m.Unlock()
    value, ok  := m.inter_map[key]
    if ok {
        delete(m.inter_map, key)
    }
    return value, ok
}

func main() {
   var m = MTMap {
       map[string]string{},
       new(sync.RWMutex),
   }
   m.Insert("name", "yansheng")
   fmt.Println(m.Get("name"))
   value, ok := m.Delete("name")
   fmt.Println(value, ok)
}
