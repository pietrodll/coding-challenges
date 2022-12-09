#[derive(PartialEq)]
enum Shape {
    Rock,
    Paper,
    Scissors,
}

fn shape_score(shape: &Shape) -> i32 {
    match shape {
        Shape::Rock => 1,
        Shape::Paper => 2,
        Shape::Scissors => 3,
    }
}

fn parse_opponent_shape(str: &str) -> Shape {
    match str {
        "A" => Shape::Rock,
        "B" => Shape::Paper,
        "C" => Shape::Scissors,
        _ => panic!("Invalid input"),
    }
}

fn parse_my_shape(str: &str) -> Shape {
    match str {
        "X" => Shape::Rock,
        "Y" => Shape::Paper,
        "Z" => Shape::Scissors,
        _ => panic!("Invalid input"),
    }
}

fn parse_line(line: &str) -> (Shape, Shape) {
    let mut split_iter = line.split_whitespace();
    let opponent_shape = parse_opponent_shape(split_iter.next().unwrap());
    let my_shape = parse_my_shape(split_iter.next().unwrap());

    return (opponent_shape, my_shape);
}

fn parse_input(data: &String) -> Vec<(Shape, Shape)> {
    data.trim().lines().map(parse_line).collect()
}

fn get_stronger_shape(shape: &Shape) -> Shape {
    match shape {
        Shape::Rock => Shape::Paper,
        Shape::Paper => Shape::Scissors,
        Shape::Scissors => Shape::Rock,
    }
}

fn game_score(opponent: &Shape, me: &Shape) -> i32 {
    if opponent == me {
        return 3 + shape_score(me);
    }

    if get_stronger_shape(opponent) == *me {
        return 6 + shape_score(me);
    }

    return shape_score(me);
}

pub fn run_first_part(data: &String) -> String {
    parse_input(data)
        .iter()
        .map(|tup| game_score(&tup.0, &tup.1))
        .sum::<i32>()
        .to_string()
}

fn get_weaker_shape(shape: &Shape) -> Shape {
    match shape {
        Shape::Rock => Shape::Scissors,
        Shape::Paper => Shape::Rock,
        Shape::Scissors => Shape::Paper,
    }
}

enum RoundResult {
    Lose,
    Draw,
    Win,
}

fn parse_round_result(str: &str) -> RoundResult {
    match str {
        "X" => RoundResult::Lose,
        "Y" => RoundResult::Draw,
        "Z" => RoundResult::Win,
        _ => panic!("Invalid round result"),
    }
}

fn parse_input_second_part(data: &String) -> Vec<(Shape, RoundResult)> {
    data.trim()
        .lines()
        .map(|line| {
            let mut split_iter = line.split_whitespace();
            let opponent_shape = parse_opponent_shape(split_iter.next().unwrap());
            let round_result = parse_round_result(split_iter.next().unwrap());

            return (opponent_shape, round_result);
        })
        .collect()
}

fn game_score_second_part(opponent: &Shape, expected_result: &RoundResult) -> i32 {
    match expected_result {
        RoundResult::Lose => shape_score(&get_weaker_shape(opponent)),
        RoundResult::Draw => 3 + shape_score(opponent),
        RoundResult::Win => 6 + shape_score(&get_stronger_shape(opponent)),
    }
}

pub fn run_second_part(data: &String) -> String {
    parse_input_second_part(data)
        .iter()
        .map(|tup| game_score_second_part(&tup.0, &tup.1))
        .sum::<i32>()
        .to_string()
}
