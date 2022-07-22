package datastructure

type Queue []string

func (q *Queue) GetLength() int {
	return len(*q)
}

func (q *Queue) Enqueue(value string) {
	*q = append(*q, value)
}

func (q *Queue) Dequeue() (string, bool) {
	if q.GetLength() == 0 {
		return "", false
	}
	element := (*q)[0]
	*q = (*q)[:1]
	return element, true
}

func CreateNewQueue() *Queue {
	return &Queue{}
}
