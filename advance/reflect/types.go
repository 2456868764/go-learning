package reflect

type Blog struct {
	Id      int64
	Title   string
	ThumbUp int64
	author  string
}

func (b *Blog) ChangeTitle(newTitle string) {
	b.Title = newTitle
}

func (b *Blog) IncreaseThumbUp() {
	b.ThumbUp += 1
}

func (b *Blog) changeId(newId int64) {
	b.Id = newId
}

func (b Blog) GetId() int64 {
	return b.Id
}
