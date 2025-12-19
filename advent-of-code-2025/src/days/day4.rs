use super::DayResult;
use crate::utils::matrix::Matrix;

fn parse_input(data: &String) -> Matrix<bool> {
    data.trim()
        .lines()
        .map(|line| {
            line.chars()
                .map(|chr| match chr {
                    '@' => true,
                    '.' => false,
                    _ => panic!("Unexpected character: {}", chr),
                })
                .collect()
        })
        .collect()
}

fn can_be_removed(data: &Matrix<bool>, row: usize, col: usize) -> bool {
    data.neighbors(row, col)
        .filter(|neighbor_val| **neighbor_val)
        .count()
        < 4
}

/// Counts the number of paper rolls (true in the data matrix) that are surrounded by
/// less than 4 paper rolls
fn first_part(data: &Matrix<bool>) -> u32 {
    data.entries_iter()
        .filter(|((row, col), val)| **val && can_be_removed(data, *row, *col))
        .count() as u32
}

fn second_part(data: &Matrix<bool>) -> u32 {
    let mut total_removed: u32 = 0;
    let mut copy = Matrix::new(data.rows(), data.cols(), false);
    let mut removed_count: isize = -1;

    // Initialize copy with the values from data
    for ((row, col), value) in data.entries_iter() {
        copy.set(row, col, *value);
    }

    while removed_count != 0 {
        let to_remove: Vec<(usize, usize)> = copy
            .entries_iter()
            .filter(|(pos, val)| **val && can_be_removed(&copy, pos.0, pos.1))
            .map(|entry| entry.0)
            .collect();

        for pos in &to_remove {
            copy.set(pos.0, pos.1, false);
        }

        removed_count = to_remove.len() as isize;
        total_removed += removed_count as u32;
    }

    total_removed
}

pub fn run(data: &String) -> DayResult {
    let matrix = parse_input(data);

    DayResult {
        first_part: first_part(&matrix).to_string(),
        second_part: second_part(&matrix).to_string(),
    }
}

#[cfg(test)]
mod test {
    use super::*;

    const INPUT_DATA: &str = "..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.";

    #[test]
    fn test_first_part() {
        let data = parse_input(&String::from(INPUT_DATA));
        assert_eq!(first_part(&data), 13);
    }

    #[test]
    fn test_second_part() {
        let data = parse_input(&String::from(INPUT_DATA));
        assert_eq!(second_part(&data), 43);
    }
}
