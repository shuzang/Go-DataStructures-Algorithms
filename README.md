Data Structure

- LinkedList

  - singlyLinkedList

    ```go
    func (list *singlyLinkedList) headInsert(val interface{})
    func (list *singlyLinkedList) tailInsert(val interface{})
    func (list *singlyLinkedList) insert(val interface{}, index int)
    func (list *singlyLinkedList) delete(index int)
    func (list *singlyLinkedList) get(index int) interface{}
    func (list *singlyLinkedList) printList()
    ```

  - doublyLinkedList

    ```go
    func (list *doublyLinkedList) headInsert(val interface{})
    func (list *doublyLinkedList) tailInsert(val interface{})
    func (list *doublyLinkedList) insert(val interface{}, index int)
    func (list *doublyLinkedList) delete(index int)
    func (list *doublyLinkedList) get(index int) interface{}
    func (list *doublyLinkedList) printList()
    ```

  - circularLinkedList

    ```go
    func convertCircular(singlylinkedlist *node)
    func traversal(circularlinkedlist *node)
    ```

- Stack

  - arrayStack

    ```go
    func (s *ArrayStack) IsEmpty() bool
    func (s *ArrayStack) Push(t Item)
    func (s *ArrayStack) Pop() Item
    func (s *ArrayStack) Peek() Item
    ```

  - linkedStack

    ```go
    func (s *LinkedStack) IsEmpty() bool
    func (s *LinkedStack) Push(t interface{})
    func (s *LinkedStack) Pop() interface{}
    func (s *LinkedStack) Peek() interface{}
    ```

- Queue

  ```go
  func (s *Queue) IsEmpty() bool
  func (s *Queue) EnQueue(t interface{})
  func (s *Queue) DnQueue() interface{}
  func (s *Queue) Peek() interface{}
  ```

  