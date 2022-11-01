package cborg

type ResultCreateDatabase struct {
	Success  bool
	Message  string
	Database Database
}

type ResultDropDatabase struct {
	Success bool
	Message string
}

type ResultListDatabases struct {
	Success   bool
	Message   string
	Databases []string
}

type ResultCreateCollection struct {
	Success    bool
	Message    string
	Collection Collection
}

type ResultDropCollection struct {
	Success bool
	Message string
}

type ResultListCollections struct {
	Success     bool
	Message     string
	Collections []string
}

type ResultFindOne struct {
	Success bool
	Message string
	Item    interface{}
}

type ResultInsertOne struct {
	Success bool
	Message string
}

type ResultUpdateOne struct {
	Success bool
	Message string
}

type ResultUpdateAll struct {
	Success bool
	Message string
}

type ResultDeleteOne struct {
	Success bool
	Message string
}

type ResultDeleteAll struct {
	Success bool
	Message string
}
