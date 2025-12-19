use super::DayResult;

fn parse_input(data: &String) -> Vec<(u64, u64)> {
    data.trim()
        .split(',')
        .map(|id_pair| {
            let ids: Vec<u64> = id_pair
                .splitn(2, '-')
                .filter_map(|id_str| id_str.parse::<u64>().ok())
                .collect();

            (ids[0], ids[1])
        })
        .collect()
}

/// Checks if the number is a repeated sequence of digits of the size passed as parameter.
/// Note: `size` is expected to be a divisor of the number of digits of `number`.
fn is_repeated_digit_sequence(number: &u64, size: u32) -> bool {
    let size_as_usize: usize = size.try_into().unwrap();
    let mut digits: Vec<u64> = Vec::with_capacity(size_as_usize);

    let mut current = *number;

    for _ in 0..size {
        digits.push(current % 10);
        current = current / 10;
    }

    let mut digits_cycle = digits.iter().cycle();

    while current != 0 {
        let expected_digit = *digits_cycle.next().unwrap();
        if current % 10 != expected_digit {
            return false;
        }

        current = current / 10;
    }

    return true;
}

/// An ID is considered invalid if it is made of a sequence of digits repeated twice
fn is_invalid(id: &u64) -> bool {
    let num_digits = id.ilog10() + 1;

    if num_digits % 2 != 0 {
        return false;
    }

    is_repeated_digit_sequence(id, num_digits / 2)
}

/// Calculate the sum of all invalid IDs in the ranges defined by the input data
fn first_part(ranges: &Vec<(u64, u64)>) -> u64 {
    ranges
        .iter()
        .flat_map(|pair| pair.0..=pair.1)
        .filter(is_invalid)
        .sum()
}

/// Returns true if the id is composed of repeated sequences of digits of any size,
/// excluding numbers composed of only one digit
fn is_invalid_extended(id: &u64) -> bool {
    let num_digits = id.ilog10() + 1;

    num_digits > 1
        && (1..=num_digits.isqrt())
            .filter(|size| num_digits % size == 0)
            .flat_map(|size| {
                if size == 1 {
                    vec![size]
                } else {
                    vec![size, num_digits / size]
                }
            })
            .any(|size| is_repeated_digit_sequence(id, size))
}

/// Calculate the sum of all extended invalid IDs in the ranges defined by the input data
fn second_part(ranges: &Vec<(u64, u64)>) -> u64 {
    ranges
        .iter()
        .flat_map(|pair| pair.0..=pair.1)
        .filter(is_invalid_extended)
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
    use std::sync::LazyLock;

    const INPUT_DATA: &str = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";
    static PARSED_INPUT: LazyLock<Vec<(u64, u64)>> =
        LazyLock::new(|| parse_input(&String::from(INPUT_DATA)));

    #[test]
    fn test_is_invalid() {
        assert!(is_invalid(&11));
        assert!(is_invalid(&22));
        assert!(is_invalid(&1010));
        assert!(is_invalid(&1188511885));
        assert!(is_invalid(&222222));
        assert!(!is_invalid(&1698522));
        assert!(!is_invalid(&1110));
    }

    #[test]
    fn test_is_repeated_digit_sequence() {
        assert!(is_repeated_digit_sequence(&565656, 2));
        assert!(is_repeated_digit_sequence(&22, 1));
    }

    #[test]
    fn test_first_part() {
        assert_eq!(first_part(&PARSED_INPUT), 1227775554);
    }

    #[test]
    fn test_second_part() {
        assert_eq!(second_part(&PARSED_INPUT), 4174379265);
    }
}
