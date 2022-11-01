package ops

const (
	opCreateDB         uint64 = 1
	opDropDB           uint64 = 2
	opCreateCollection uint64 = 3
	opDropCollection   uint64 = 4
	opInsertOne        uint64 = 5
	opFindOne          uint64 = 6
	opUpdateOne        uint64 = 7
	opUpdateAll        uint64 = 8
	opDeleteOne        uint64 = 9
	opDeleteAll        uint64 = 10
	opListDBs          uint64 = 20
	opListCollections  uint64 = 21
)
