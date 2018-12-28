package main


type rPlugin int

func (rp rPlugin) RunPlugin(jArgs string) string {
	return "hello" + jArgs
}

// exported as symbol named "Greeter"
var RPlugin rPlugin
