use crate::days::DAYS;

mod days;
mod utils;

fn read_input(day: i32) -> String {
    std::fs::read_to_string(format!("data/day{day}.txt")).unwrap()
}

fn run_day(day: i32, runner: fn(&String) -> days::DayResult) {
    println!("Running day {}", day);
    let data = read_input(day);
    let result = runner(&data);
    println!("Result of part 1: {}", result.first_part);
    println!("Result of part 2: {}", result.second_part);
}

fn main() {
    let args: Vec<String> = std::env::args().collect();

    match args.get(1) {
        Some(day_arg) => match day_arg.parse::<usize>() {
            Ok(day) => match DAYS.get(day - 1) {
                Some(runner) => run_day(day.try_into().unwrap(), *runner),
                None => panic!("Day {} not yet implemented", day),
            },
            Err(_) => panic!("Cannot parse day arg"),
        },
        None => panic!("Missing day argument"),
    }
}
