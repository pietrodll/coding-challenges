use super::DayResult;

fn parse_input(data: &String) -> Vec<Vec<u32>> {
    data.trim()
        .lines()
        .map(|line| {
            line.chars()
                .map(|char| char.to_digit(10).unwrap())
                .collect()
        })
        .collect()
}

/// Returns the maximum of the slice and its index, taking the smallest
/// index if multiple values are maximal. Returns None if the slice is empty.
fn earliest_max_and_index(slice: &[u32]) -> Option<(usize, u32)> {
    slice
        .iter()
        .enumerate()
        .reduce(|(max_idx, max), (idx, digit)| {
            if digit > max {
                (idx, digit)
            } else {
                (max_idx, max)
            }
        })
        .map(|(idx, max)| (idx, *max))
}

fn highest_possible_joltage(battery_bank: &Vec<u32>, size: usize) -> u64 {
    let mut joltage: u64 = 0;
    let mut curr_size: usize = size;
    let mut curr_slice = &battery_bank[..];

    while curr_size > 0 && curr_slice.len() >= curr_size {
        // Exclude curr_size elements at the end of the slice when computing the
        // maximum: this is necessary to have enough elements left at each step
        let (max_index, max) =
            earliest_max_and_index(&curr_slice[..=curr_slice.len() - curr_size]).unwrap();

        joltage = 10 * joltage + max as u64;
        curr_slice = &curr_slice[max_index + 1..];
        curr_size -= 1;
    }

    joltage
}

fn first_part(batteries: &Vec<Vec<u32>>) -> u64 {
    batteries
        .iter()
        .map(|battery_bank| highest_possible_joltage(battery_bank, 2))
        .sum()
}

fn second_part(batteries: &Vec<Vec<u32>>) -> u64 {
    batteries
        .iter()
        .map(|battery_bank| highest_possible_joltage(battery_bank, 12))
        .sum()
}

pub fn run(data: &String) -> DayResult {
    let input = parse_input(data);

    let first_part = first_part(&input).to_string();
    let second_part = second_part(&input).to_string();

    DayResult {
        first_part,
        second_part,
    }
}

#[cfg(test)]
mod test {
    use super::*;

    const INPUT_DATA: &str = "987654321111111
811111111111119
234234234234278
818181911112111";

    #[test]
    fn test_highest_possible_joltage() {
        let batteries = parse_input(&String::from(INPUT_DATA));
        assert_eq!(highest_possible_joltage(&batteries[0], 12), 987654321111);
        assert_eq!(highest_possible_joltage(&batteries[0], 2), 98);
        assert_eq!(highest_possible_joltage(&batteries[1], 2), 89);
    }

    #[test]
    fn test_first_part() {
        let batteries = parse_input(&String::from(INPUT_DATA));
        assert_eq!(first_part(&batteries), 357);
    }

    #[test]
    fn test_second_part() {
        let batteries = parse_input(&String::from(INPUT_DATA));
        assert_eq!(second_part(&batteries), 3121910778619);
    }
}
