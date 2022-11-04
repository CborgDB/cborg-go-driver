package cborg

type OpCode uint64

const (
	OpCreateDatabase   OpCode = 1
	OpDropDatabase     OpCode = 2
	OpCreateCollection OpCode = 3
	OpDropCollection   OpCode = 4
	OpInsertOne        OpCode = 5
	OpFindOne          OpCode = 6
	OpUpdateOne        OpCode = 7
	OpUpdateAll        OpCode = 8
	OpDeleteOne        OpCode = 9
	OpDeleteAll        OpCode = 10
	OpListDatabases    OpCode = 20
	OpListCollections  OpCode = 21
	OpReply            OpCode = 666
)
