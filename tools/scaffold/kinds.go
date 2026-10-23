package main

// kind 描述一种内容种类：它落在哪棵树、落成什么形态、写什么内容。
// 这张表是「内容区结构」的唯一权威；加一种内容种类只需要加一行。
type kind struct {
	// name 是命令行上的种类名。
	name string
	// root 是内容区里的那棵树，用斜杠分隔，便于跨平台输出。
	root string
	// file 是落在条目目录里的入口文件名；为空表示该种类只落一个文件（由 entry 指定）。
	entry string
}

var kinds = []kind{
	{name: "lecture", root: "src/lectures", entry: "main.go"},
}

func lookupKind(name string) (kind, bool) {
	for _, k := range kinds {
		if k.name == name {
			return k, true
		}
	}

	return kind{}, false
}
