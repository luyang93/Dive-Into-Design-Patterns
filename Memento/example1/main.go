package main

type Editor struct {
	Text           string
	CurX, CurY     int
	SelectionWidth int
}

func (e *Editor) SetText(text string) {
	e.Text = text
}

func (e *Editor) SetCursor(x, y int) {
	e.CurX = x
	e.CurY = y
}

func (e *Editor) SetSelectionWidth(width int) {
	e.SelectionWidth = width
}

func (e *Editor) CreateSnapshot() *Snapshot {
	return &Snapshot{e, e.Text, e.CurX, e.CurY, e.SelectionWidth}
}

type Snapshot struct {
	editor         *Editor
	text           string
	curX, curY     int
	selectionWidth int
}

func NewSnapshot(editor *Editor, text string, curX, curY, selectionWidth int) *Snapshot {
	return &Snapshot{editor, text, curX, curY, selectionWidth}
}

func (s *Snapshot) Restore() {
	s.editor.SetText(s.text)
	s.editor.SetCursor(s.curX, s.curY)
	s.editor.SetSelectionWidth(s.selectionWidth)
}

type Command struct {
	backup []*Snapshot
	editor *Editor
}

func (c *Command) MakeBackup() {
	c.backup = append(c.backup, c.editor.CreateSnapshot())
}

func (c *Command) Undo() {
	if len(c.backup) > 0 {
		c.backup[len(c.backup)-1].Restore()
		c.backup = c.backup[:len(c.backup)-1]
	}
}
