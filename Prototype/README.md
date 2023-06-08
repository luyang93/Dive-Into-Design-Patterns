```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface inode {
    + clone() inode
    + print(string)
} 

class file {
    + name string
    + clone() inode
    + print(string)
}

class folder {
    + name string
    + children []inode
    + clone() inode
    + print(string)
}

class client {
    + inode.clone() inode
}

file ..|> inode
folder ..|> inode

client --> inode
```