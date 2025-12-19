mod day1;
mod day2;

pub struct DayResult {
    pub first_part: String,
    pub second_part: String,
}

pub const DAYS: &'static [fn(&String) -> DayResult] = &[day1::run, day2::run];
