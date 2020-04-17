package talker

type Talker interface {
	SayHello(word string) (response string)
}
