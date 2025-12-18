mod day1;

pub struct DayResult {
    pub first_part: String,
    pub second_part: String,
}

pub const DAYS: &'static [fn(&String) -> DayResult] = &[
    day1::run,
];
