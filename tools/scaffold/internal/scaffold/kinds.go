package scaffold

import "path"

// kind 描述一种内容种类：落在哪棵树、落成什么形态、写什么内容。
// 这张表是「内容区结构」的唯一权威；加一种种类只需要加一行。
type kind struct {
	// name 是命令行上的种类名。
	name string
	// root 是内容区里的那棵树，用斜杠分隔，便于跨平台输出。
	root string
	// dirEntry 非空时条目是一个目录，目录里放这个入口文件。
	dirEntry string
	// fileExt 非空时条目是主题目录下的单个文件，扩展名取自这里。
	fileExt string
	// content 生成条目内容，入参是已经定好的标题（缺省即条目名）。
	content func(title string) string
	// titled 表示这种种类接受 --title。
	titled bool
	// numbered 表示条目名带两位题号，编号由工具按主题目录的现状分配。
	numbered bool
	// argLabel 是第二个位置参数在用法提示里的名字，缺省叫 name。
	argLabel string
}

// relPath 给出条目相对于仓库根的路径。
func (k kind) relPath(topic, name string) string {
	if k.dirEntry != "" {
		return path.Join(k.root, topic, name, k.dirEntry)
	}

	return path.Join(k.root, topic, name+k.fileExt)
}

// argName 是第二个位置参数在提示里叫什么。
func (k kind) argName() string {
	if k.argLabel == "" {
		return "name"
	}

	return k.argLabel
}

// usageLine 是这种种类在出错时给出的用法提示。
func (k kind) usageLine() string {
	if k.titled {
		return "scaffold " + k.name + " [--dry-run] [--title <标题>] <topic> <" + k.argName() + ">"
	}

	return "scaffold " + k.name + " [--dry-run] <topic> <" + k.argName() + ">"
}

var kinds = []kind{
	{
		name:    "note",
		root:    "src/notes",
		fileExt: ".md",
		content: noteContent,
		titled:  true,
	},
	{
		name:     "exercise",
		root:     "src/homework",
		dirEntry: "main.go",
		content:  exerciseContent,
		numbered: true,
		argLabel: "slug",
	},
}

func lookupKind(name string) (kind, bool) {
	for _, k := range kinds {
		if k.name == name {
			return k, true
		}
	}

	return kind{}, false
}
