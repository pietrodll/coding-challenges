mod day1;
mod day2;
mod day3;
mod day4;
mod day5;
mod day6;
mod day7;
mod day8;

pub const DAYS: &'static [&'static (fn(&String) -> String, fn(&String) -> String)] = &[
    &(day1::run_first_part, day1::run_second_part),
    &(day2::run_first_part, day2::run_second_part),
    &(day3::run_first_part, day3::run_second_part),
    &(day4::run_first_part, day4::run_second_part),
    &(day5::run_first_part, day5::run_second_part),
    &(day6::run_first_part, day6::run_second_part),
    &(day7::run_first_part, day7::run_second_part),
    &(day8::run_first_part, day8::run_second_part),
];
