mod day1;
mod day2;
mod day3;
mod day4;

pub struct DayResult {
    pub first_part: String,
    pub second_part: String,
}

pub const DAYS: &'static [fn(&String) -> DayResult] = &[day1::run, day2::run, day3::run, day4::run];
