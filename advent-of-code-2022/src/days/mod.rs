mod day1;
mod day2;
mod day3;

pub const DAYS: &'static [&'static (fn(&String) -> String, fn(&String) -> String)] = &[
    &(day1::run_first_part, day1::run_second_part),
    &(day2::run_first_part, day2::run_second_part),
    &(day3::run_first_part, day3::run_second_part),
];
