use std::{cell::RefCell, rc::Rc, vec::Vec};

#[derive(PartialEq, Clone, Copy)]
enum NodeType {
    Dir,
    File,
}

struct Node {
    name: String,
    node_type: NodeType,
    parent: Option<Rc<RefCell<Node>>>,
    size: i64,
    children: Vec<Rc<RefCell<Node>>>,
}

impl Node {
    fn new(
        name: String,
        size: i64,
        node_type: NodeType,
        parent: Option<Rc<RefCell<Node>>>,
    ) -> Rc<RefCell<Node>> {
        Rc::new(RefCell::new(Node {
            name,
            node_type,
            parent,
            size,
            children: Vec::new(),
        }))
    }

    fn new_file(name: String, size: i64, parent: Option<Rc<RefCell<Node>>>) -> Rc<RefCell<Node>> {
        Node::new(name, size, NodeType::File, parent)
    }

    fn new_dir(name: String, parent: Option<Rc<RefCell<Node>>>) -> Rc<RefCell<Node>> {
        Node::new(name, 0, NodeType::Dir, parent)
    }

    fn add_child(&mut self, child: Rc<RefCell<Node>>) {
        self.children.push(child)
    }

    fn find_child(&self, name: &String) -> Option<Rc<RefCell<Node>>> {
        for child in &self.children {
            if child.try_borrow().unwrap().name.eq(name) {
                return Option::Some(child.clone());
            }
        }

        return Option::None;
    }

    fn is_file(&self) -> bool {
        self.node_type == NodeType::File
    }

    fn is_dir(&self) -> bool {
        self.node_type == NodeType::Dir
    }
}

enum FsDescriptor {
    File(String, i64),
    Dir(String),
}

enum Command {
    Cd(String),
    Ls(Vec<FsDescriptor>),
}

fn parse_input(data: &str) -> Vec<Command> {
    let mut commands: Vec<Command> = Vec::new();

    let mut lines = data.lines().peekable();

    while lines.peek().is_some() {
        let cmd = lines.next().unwrap();

        if cmd.starts_with("$ cd ") {
            commands.push(Command::Cd(String::from(&cmd[5..])));
            continue;
        }

        let mut files: Vec<FsDescriptor> = Vec::new();

        while lines.peek().is_some() && !lines.peek().unwrap().starts_with('$') {
            let descriptor_str = lines.next().unwrap();

            if descriptor_str.starts_with("dir ") {
                files.push(FsDescriptor::Dir(descriptor_str[4..].to_string()));
            } else {
                let mut split = descriptor_str.split_ascii_whitespace();
                let size: i64 = split.next().unwrap().parse().unwrap();
                let name: String = split.next().unwrap().to_string();
                files.push(FsDescriptor::File(name, size));
            }
        }

        commands.push(Command::Ls(files))
    }

    return commands;
}

fn compute_filesystem(commands: &Vec<Command>) -> Rc<RefCell<Node>> {
    let root = Node::new("".to_string(), 0, NodeType::Dir, Option::None);

    let mut current: Rc<RefCell<Node>> = root.clone();

    for command in commands {
        match command {
            Command::Cd(dirname) => {
                let next: Rc<RefCell<Node>>;

                if dirname == ".." {
                    let curr = current.try_borrow().unwrap();
                    next = curr.parent.as_ref().unwrap().clone();
                } else if dirname == "/" {
                    next = root.clone();
                } else {
                    let maybe_child = current.borrow().find_child(&dirname);

                    if maybe_child.is_some() {
                        next = maybe_child.unwrap();
                    } else {
                        // The directory is not tracked yet, we need to add it to the children
                        let child = Node::new_dir(dirname.clone(), Option::Some(current.clone()));
                        let mut mut_curr = current.borrow_mut();
                        mut_curr.add_child(child.clone());
                        next = child.clone();
                    }
                }

                current = next;
            }
            Command::Ls(descriptors) => {
                let mut mut_current = current.borrow_mut();

                for descriptor in descriptors {
                    match descriptor {
                        FsDescriptor::File(name, size) => {
                            mut_current.add_child(Node::new_file(
                                name.clone(),
                                *size,
                                Option::Some(current.clone()),
                            ));
                        }
                        FsDescriptor::Dir(name) => {
                            mut_current.add_child(Node::new_dir(
                                name.clone(),
                                Option::Some(current.clone()),
                            ));
                        }
                    }
                }
            }
        }
    }

    return root;
}

fn populate_dir_size(node: Rc<RefCell<Node>>) {
    let mut n = node.borrow_mut();

    if n.is_file() {
        return;
    }

    let mut size = 0;

    for child in &n.children {
        if child.borrow().is_dir() {
            populate_dir_size(child.clone());
        }

        size += child.borrow().size;
    }

    n.size = size;
}

fn compute_dir_total_at_most_100000(node: &Rc<RefCell<Node>>) -> i64 {
    let mut res: i64 = 0;

    let n = node.borrow();

    if n.is_dir() && n.size <= 100_000 {
        res += n.size;
    }

    for child in &n.children {
        res += compute_dir_total_at_most_100000(&child);
    }

    return res;
}

fn list_dirs_rec(node: &Rc<RefCell<Node>>, list: &mut Vec<Rc<RefCell<Node>>>) {
    if node.borrow().is_dir() {
        list.push(node.clone());
    }

    for child in &node.borrow().children {
        list_dirs_rec(child, list)
    }
}

fn find_size_of_smallest_to_delete(node: &Rc<RefCell<Node>>, space_to_recover: i64) -> i64 {
    let mut directories_list: Vec<Rc<RefCell<Node>>> = Vec::new();
    list_dirs_rec(node, &mut directories_list);

    if directories_list.is_empty() {
        return 0;
    }

    let mut deleted_size = i64::MAX;

    for dir in directories_list {
        let size = dir.borrow().size;

        if size >= space_to_recover && size < deleted_size {
            deleted_size = size;
        }
    }

    return deleted_size;
}

pub fn run_first_part(data: &String) -> String {
    let commands = parse_input(data);
    let filesystem = compute_filesystem(&commands);

    populate_dir_size(filesystem.clone());

    return compute_dir_total_at_most_100000(&filesystem).to_string();
}

pub fn run_second_part(data: &String) -> String {
    let commands = parse_input(data);
    let filesystem = compute_filesystem(&commands);

    populate_dir_size(filesystem.clone());

    let total_size = filesystem.as_ref().borrow().size;
    let to_recover = 30000000 - (70000000 - total_size);

    return find_size_of_smallest_to_delete(&filesystem, to_recover).to_string();
}
