use std::vec::Vec;
use regex::Regex;

struct MovingDirection {
    quantity: usize,
    from: usize,
    to: usize,
}

fn parse_moving_direction(line: &str) -> MovingDirection {
    lazy_static::lazy_static! {
        static ref RE: Regex = Regex::new(r"^move\s(\d+)\sfrom\s(\d+)\sto\s(\d+)$").unwrap();
    }

    let captures = RE.captures(line).unwrap();
    return MovingDirection {
        quantity: captures.get(1).unwrap().as_str().parse().unwrap(),
        from: captures.get(2).unwrap().as_str().parse().unwrap(),
        to: captures.get(3).unwrap().as_str().parse().unwrap(),
    };
}

fn parse_size(data: &String) -> (usize, usize) {
    for (idx, line) in data.lines().enumerate() {
        if line.starts_with(" 1") {
            let size_str = line.trim().split(' ').last().unwrap();

            return (idx, size_str.parse().unwrap());
        }
    }

    panic!("could not find size")
}

fn parse_input(data: &String) -> (Vec<Vec<char>>, Vec<MovingDirection>) {
    let (size_line_idx, size) = parse_size(data);
    let lines: Vec<&str> = data.lines().collect();
    let mut stacks: Vec<Vec<char>> = Vec::new();

    for idx in 0..size {
        let mut vec: Vec<char> = Vec::new();

        for line_idx in (0..size_line_idx).rev() {
            let chr = lines
                .get(line_idx)
                .unwrap()
                .chars()
                .nth(4 * idx + 1)
                .unwrap();

            if chr != ' ' {
                vec.push(chr);
            }
        }

        stacks.push(vec);
    }

    let mut moving_directions: Vec<MovingDirection> = Vec::new();

    for line_idx in (size_line_idx + 2)..lines.len() {
        moving_directions.push(parse_moving_direction(lines.get(line_idx).unwrap()));
    }

    return (stacks, moving_directions);
}

fn top_crates(stacks: &Vec<Vec<char>>) -> String {
    stacks.iter().map(|stack| stack.last().unwrap()).collect()
}

pub fn run_first_part(data: &String) -> String {
    let (mut stacks, directions) = parse_input(data);

    for direction in directions {
        for _ in 0..direction.quantity {
            let from = &mut stacks[direction.from - 1];
            let chr = from.pop().unwrap();

            let to = &mut stacks[direction.to - 1];
            to.push(chr);
        }
    }

    top_crates(&stacks)
}

pub fn run_second_part(data: &String) -> String {
    let (mut stacks, directions) = parse_input(data);

    for direction in directions {
        let from = &mut stacks[direction.from - 1];

        let mut to_move = from.split_off(from.len() - direction.quantity);

        let to = &mut stacks[direction.to - 1];
        to.append(&mut to_move);
    }

    top_crates(&stacks)
}
