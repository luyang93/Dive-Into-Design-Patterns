package main

import "fmt"

type ComponentWithContextualHelp interface {
	ShowHelp() string
	SetContainer(Container)
}

type Container interface {
	ComponentWithContextualHelp
	Add(child ComponentWithContextualHelp)
}

type Component struct {
	tooltipText string
	container   Container
}

func (c *Component) ShowHelp() string {
	if c.tooltipText != "" {
		return c.tooltipText
	}
	if c.container != nil {
		return c.container.ShowHelp()
	}
	return ""
}

func (c *Component) SetContainer(container Container) {
	c.container = container
}

type BaseContainer struct {
	Component
	children []ComponentWithContextualHelp
}

func (bc *BaseContainer) Add(child ComponentWithContextualHelp) {
	bc.children = append(bc.children, child)
	child.SetContainer(bc)
}

type Button struct {
	Component
}

type Panel struct {
	BaseContainer
	modalHelpText string
}

func (p *Panel) ShowHelp() string {
	if p.modalHelpText != "" {
		return p.modalHelpText
	}
	return p.Component.ShowHelp()
}

type Dialog struct {
	BaseContainer
	wikiPageURL string
}

func (d *Dialog) ShowHelp() string {
	if d.wikiPageURL != "" {
		return d.wikiPageURL
	}
	return d.Component.ShowHelp()
}

func main() {
	dialog := Dialog{wikiPageURL: "http://..."}
	panel := Panel{modalHelpText: "This panel is used for…"}
	ok := Button{Component: Component{tooltipText: "This is an OK button…"}}

	panel.Add(&ok)
	dialog.Add(&panel2)

	fmt.Println(dialog.ShowHelp())
	fmt.Println(panel.ShowHelp())
	fmt.Println(ok.ShowHelp())
}
