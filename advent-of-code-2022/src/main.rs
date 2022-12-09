use std::env;

use crate::days::DAYS;

mod days;

fn read_input(day: i32) -> String {
    std::fs::read_to_string(format!("data/day{day}.txt")).unwrap()
}

fn run_day(day: i32, runner: &'static (fn(&String) -> String, fn(&String) -> String)) {
    println!("Running day {}", day);
    let data = read_input(day);
    println!("Result of part 1: {}", runner.0(&data));
    println!("Result of part 2: {}", runner.1(&data));
}

fn main() {
    let args: Vec<String> = env::args().collect();

    match args.get(1) {
        Some(day_arg) => match day_arg.parse::<usize>() {
            Ok(day) => match DAYS.get(day - 1) {
                Some(runner) => run_day(day.try_into().unwrap(), runner),
                None => panic!("Day {} not yet implemented", day),
            },
            Err(_) => panic!("Cannot parse day arg"),
        },
        None => panic!("Missing day argument"),
    }
}
